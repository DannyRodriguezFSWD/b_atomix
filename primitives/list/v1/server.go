// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	listv1 "github.com/atomix/runtime/api/atomix/list/v1"
	"github.com/atomix/runtime/pkg/errors"
	"github.com/atomix/runtime/pkg/logging"
	"github.com/atomix/runtime/pkg/primitive"
)

func newListServer(proxies *primitive.Manager[listv1.ListServer]) listv1.ListServer {
	return &listServer{
		proxies: proxies,
	}
}

type listServer struct {
	proxies *primitive.Manager[listv1.ListServer]
}

func (s *listServer) Create(ctx context.Context, request *listv1.CreateRequest) (*listv1.CreateResponse, error) {
	log.Debugw("Create",
		logging.Stringer("CreateRequest", request))
	proxy, err := s.proxies.Connect(ctx)
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Create",
			logging.Stringer("CreateRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response, err := proxy.Create(ctx, request)
	if err != nil {
		log.Warnw("Create",
			logging.Stringer("CreateRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	log.Debugw("Create",
		logging.Stringer("CreateResponse", response))
	return response, nil
}

func (s *listServer) Close(ctx context.Context, request *listv1.CloseRequest) (*listv1.CloseResponse, error) {
	log.Debugw("Close",
		logging.Stringer("CloseRequest", request))
	proxy, err := s.proxies.Close(ctx)
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Close",
			logging.Stringer("CloseRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response, err := proxy.Close(ctx, request)
	if err != nil {
		log.Warnw("Close",
			logging.Stringer("CloseRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	log.Debugw("Close",
		logging.Stringer("CloseResponse", response))
	return response, nil
}

func (s *listServer) Size(ctx context.Context, request *listv1.SizeRequest) (*listv1.SizeResponse, error) {
	log.Debugw("Size",
		logging.Stringer("SizeRequest", request))
	proxy, err := s.proxies.Get(ctx)
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Size",
			logging.Stringer("SizeRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response, err := proxy.Size(ctx, request)
	if err != nil {
		log.Warnw("Size",
			logging.Stringer("SizeRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	log.Debugw("Size",
		logging.Stringer("SizeResponse", response))
	return response, nil
}

func (s *listServer) Append(ctx context.Context, request *listv1.AppendRequest) (*listv1.AppendResponse, error) {
	log.Debugw("Append",
		logging.Stringer("AppendRequest", request))
	proxy, err := s.proxies.Get(ctx)
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Append",
			logging.Stringer("AppendRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response, err := proxy.Append(ctx, request)
	if err != nil {
		log.Warnw("Append",
			logging.Stringer("AppendRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	log.Debugw("Append",
		logging.Stringer("AppendResponse", response))
	return response, nil
}

func (s *listServer) Insert(ctx context.Context, request *listv1.InsertRequest) (*listv1.InsertResponse, error) {
	log.Debugw("Insert",
		logging.Stringer("InsertRequest", request))
	proxy, err := s.proxies.Get(ctx)
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Insert",
			logging.Stringer("InsertRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response, err := proxy.Insert(ctx, request)
	if err != nil {
		log.Warnw("Insert",
			logging.Stringer("InsertRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	log.Debugw("Insert",
		logging.Stringer("InsertResponse", response))
	return response, nil
}

func (s *listServer) Get(ctx context.Context, request *listv1.GetRequest) (*listv1.GetResponse, error) {
	log.Debugw("Get",
		logging.Stringer("GetRequest", request))
	proxy, err := s.proxies.Get(ctx)
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Get",
			logging.Stringer("GetRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response, err := proxy.Get(ctx, request)
	if err != nil {
		log.Warnw("Get",
			logging.Stringer("GetRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	log.Debugw("Get",
		logging.Stringer("GetResponse", response))
	return response, nil
}

func (s *listServer) Set(ctx context.Context, request *listv1.SetRequest) (*listv1.SetResponse, error) {
	log.Debugw("Set",
		logging.Stringer("SetRequest", request))
	proxy, err := s.proxies.Get(ctx)
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Set",
			logging.Stringer("SetRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response, err := proxy.Set(ctx, request)
	if err != nil {
		log.Warnw("Set",
			logging.Stringer("SetRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	log.Debugw("Set",
		logging.Stringer("SetResponse", response))
	return response, nil
}

func (s *listServer) Remove(ctx context.Context, request *listv1.RemoveRequest) (*listv1.RemoveResponse, error) {
	log.Debugw("Remove",
		logging.Stringer("RemoveRequest", request))
	proxy, err := s.proxies.Get(ctx)
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Remove",
			logging.Stringer("RemoveRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response, err := proxy.Remove(ctx, request)
	if err != nil {
		log.Warnw("Remove",
			logging.Stringer("RemoveRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	log.Debugw("Remove",
		logging.Stringer("RemoveResponse", response))
	return response, nil
}

func (s *listServer) Clear(ctx context.Context, request *listv1.ClearRequest) (*listv1.ClearResponse, error) {
	log.Debugw("Clear",
		logging.Stringer("ClearRequest", request))
	proxy, err := s.proxies.Get(ctx)
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Clear",
			logging.Stringer("ClearRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response, err := proxy.Clear(ctx, request)
	if err != nil {
		log.Warnw("Clear",
			logging.Stringer("ClearRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	log.Debugw("Clear",
		logging.Stringer("ClearResponse", response))
	return response, nil
}

func (s *listServer) Events(request *listv1.EventsRequest, server listv1.List_EventsServer) error {
	log.Debugw("Events",
		logging.Stringer("EventsRequest", request))
	proxy, err := s.proxies.Get(server.Context())
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Events",
			logging.Stringer("EventsRequest", request),
			logging.Error("Error", err))
		return err
	}
	err = proxy.Events(request, server)
	if err != nil {
		log.Warnw("Events",
			logging.Stringer("EventsRequest", request),
			logging.Error("Error", err))
		return err
	}
	return nil
}

func (s *listServer) Elements(request *listv1.ElementsRequest, server listv1.List_ElementsServer) error {
	log.Debugw("Elements",
		logging.Stringer("ElementsRequest", request))
	proxy, err := s.proxies.Get(server.Context())
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Elements",
			logging.Stringer("ElementsRequest", request),
			logging.Error("Error", err))
		return err
	}
	err = proxy.Elements(request, server)
	if err != nil {
		log.Warnw("Elements",
			logging.Stringer("ElementsRequest", request),
			logging.Error("Error", err))
		return err
	}
	return nil
}

var _ listv1.ListServer = (*listServer)(nil)
