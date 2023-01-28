// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package v3beta4

import (
	"context"
	atomixv3beta3 "github.com/atomix/atomix/controller/pkg/apis/atomix/v3beta3"
	atomixv3beta4 "github.com/atomix/atomix/controller/pkg/apis/atomix/v3beta4"
	"github.com/atomix/atomix/controller/pkg/controller/util/k8s/conversion"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/json"
	"net/http"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

const (
	dataStoreConvertPath = "/convert-data-store"
)

func addDataStoreController(mgr manager.Manager) error {
	mgr.GetWebhookServer().Register(storageProfileConvertPath, &conversion.Webhook{
		Handler: &DataStoreConverter{
			client: mgr.GetClient(),
			scheme: mgr.GetScheme(),
		},
	})
	return nil
}

// DataStoreConverter is a mutating webhook that converts old versions of DataStore resources
type DataStoreConverter struct {
	client  client.Client
	scheme  *runtime.Scheme
	decoder *conversion.Decoder
}

// InjectDecoder :
func (i *DataStoreConverter) InjectDecoder(decoder *conversion.Decoder) error {
	i.decoder = decoder
	return nil
}

// Handle :
func (i *DataStoreConverter) Handle(ctx context.Context, request conversion.Request) conversion.Response {
	log.Infof("Received conversion request for DataStore '%s'", request.UID)

	desiredGV, err := schema.ParseGroupVersion(request.DesiredAPIVersion)
	if err != nil {
		return conversion.Errored(http.StatusBadRequest, "failed parsing desired API version: %s", err.Error())
	}

	objectsOut := make([]runtime.RawExtension, 0, len(request.Objects))
	for _, rawObjIn := range request.Objects {
		objIn := &unstructured.Unstructured{}
		if err := objIn.UnmarshalJSON(rawObjIn.Raw); err != nil {
			return conversion.Errored(http.StatusBadRequest, "converter failed unmarshaling object: %s", err.Error())
		}

		var in runtime.Object
		var out runtime.Object
		switch objIn.GetAPIVersion() {
		case atomixv3beta3.SchemeGroupVersion.Identifier():
			switch desiredGV.Version {
			case atomixv3beta4.SchemeGroupVersion.Version:
				in = &atomixv3beta3.DataStore{}
				out = &atomixv3beta4.DataStore{}
			}
		case atomixv3beta4.SchemeGroupVersion.Identifier():
			switch desiredGV.Version {
			case atomixv3beta3.SchemeGroupVersion.Version:
				in = &atomixv3beta4.DataStore{}
				out = &atomixv3beta3.DataStore{}
			}
		}

		if err := i.decoder.Decode(rawObjIn, in); err != nil {
			return conversion.Errored(http.StatusBadRequest, "converter failed decoding DataStore: %s", err.Error())
		}
		if err := i.scheme.Convert(in, out, nil); err != nil {
			return conversion.Errored(http.StatusBadRequest, "converter failed converting DataStore: %s", err.Error())
		}
		bytesOut, err := json.Marshal(out)
		if err != nil {
			return conversion.Errored(http.StatusBadRequest, "converter failed encoding DataStore: %s", err.Error())
		}
		rawObjOut := runtime.RawExtension{
			Raw:    bytesOut,
			Object: out,
		}
		objectsOut = append(objectsOut, rawObjOut)
	}
	return conversion.Converted(objectsOut...)
}

var _ conversion.Handler = &DataStoreConverter{}
