// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	listv1 "github.com/atomix/runtime/api/atomix/list/v1"
	"github.com/atomix/runtime/pkg/runtime"
	"google.golang.org/grpc"
)

const (
	Name       = "List"
	APIVersion = "v1"
)

var Type = runtime.NewType[listv1.ListClient](Name, APIVersion, register, resolve)

func register(server *grpc.Server, delegate *runtime.Delegate[listv1.ListClient]) {
	listv1.RegisterListServer(server, newListServer(delegate))
}

func resolve(client runtime.Client) (listv1.ListClient, bool) {
	if provider, ok := client.(ListProvider); ok {
		return provider.List(), true
	}
	return nil, false
}

type ListProvider interface {
	List() listv1.ListClient
}
