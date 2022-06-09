// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	setv1 "github.com/atomix/runtime/api/atomix/set/v1"
	"github.com/atomix/runtime/pkg/errors"
	"github.com/atomix/runtime/pkg/logging"
	"github.com/atomix/runtime/pkg/primitive"
	"io"
)

func newSetServer(proxies *primitive.Manager[setv1.SetClient]) setv1.SetServer {
	return &setServer{
		proxies: proxies,
	}
}

type setServer struct {
	proxies *primitive.Manager[setv1.SetClient]
}

func (s *setServer) Create(ctx context.Context, request *setv1.CreateRequest) (*setv1.CreateResponse, error) {
	log.Debugw("Create",
		logging.Stringer("CreateRequest", request))
	proxy, err := s.proxies.Create(ctx)
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

func (s *setServer) Close(ctx context.Context, request *setv1.CloseRequest) (*setv1.CloseResponse, error) {
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

func (s *setServer) Size(ctx context.Context, request *setv1.SizeRequest) (*setv1.SizeResponse, error) {
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

func (s *setServer) Contains(ctx context.Context, request *setv1.ContainsRequest) (*setv1.ContainsResponse, error) {
	log.Debugw("Contains",
		logging.Stringer("ContainsRequest", request))
	proxy, err := s.proxies.Get(ctx)
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Contains",
			logging.Stringer("ContainsRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response, err := proxy.Contains(ctx, request)
	if err != nil {
		log.Warnw("Contains",
			logging.Stringer("ContainsRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	log.Debugw("Contains",
		logging.Stringer("ContainsResponse", response))
	return response, nil
}

func (s *setServer) Add(ctx context.Context, request *setv1.AddRequest) (*setv1.AddResponse, error) {
	log.Debugw("Add",
		logging.Stringer("AddRequest", request))
	proxy, err := s.proxies.Get(ctx)
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Add",
			logging.Stringer("AddRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response, err := proxy.Add(ctx, request)
	if err != nil {
		log.Warnw("Add",
			logging.Stringer("AddRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	log.Debugw("Add",
		logging.Stringer("AddResponse", response))
	return response, nil
}

func (s *setServer) Remove(ctx context.Context, request *setv1.RemoveRequest) (*setv1.RemoveResponse, error) {
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

func (s *setServer) Clear(ctx context.Context, request *setv1.ClearRequest) (*setv1.ClearResponse, error) {
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

func (s *setServer) Events(request *setv1.EventsRequest, server setv1.Set_EventsServer) error {
	log.Debugw("Events",
		logging.Stringer("EventsRequest", request),
		logging.String("State", "started"))
	proxy, err := s.proxies.Get(server.Context())
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Events",
			logging.Stringer("EventsRequest", request),
			logging.Error("Error", err))
		return err
	}
	client, err := proxy.Events(server.Context(), request)
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Events",
			logging.Stringer("EventsRequest", request),
			logging.Error("Error", err))
		return err
	}
	for {
		response, err := client.Recv()
		if err == io.EOF {
			log.Debugw("Events",
				logging.Stringer("EventsRequest", request),
				logging.String("State", "complete"))
			return nil
		}
		if err != nil {
			err = errors.ToProto(err)
			log.Warnw("Events",
				logging.Stringer("EventsRequest", request),
				logging.Error("Error", err))
			return err
		}
		log.Warnw("Events",
			logging.Stringer("EventsResponse", response))
		err = server.Send(response)
		if err != nil {
			err = errors.ToProto(err)
			log.Warnw("Events",
				logging.Stringer("EventsRequest", request),
				logging.Error("Error", err))
			return err
		}
	}
}

func (s *setServer) Elements(request *setv1.ElementsRequest, server setv1.Set_ElementsServer) error {
	log.Debugw("Elements",
		logging.Stringer("ElementsRequest", request),
		logging.String("State", "started"))
	proxy, err := s.proxies.Get(server.Context())
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Elements",
			logging.Stringer("ElementsRequest", request),
			logging.Error("Error", err))
		return err
	}
	client, err := proxy.Elements(server.Context(), request)
	if err != nil {
		err = errors.ToProto(err)
		log.Warnw("Elements",
			logging.Stringer("ElementsRequest", request),
			logging.Error("Error", err))
		return err
	}
	for {
		response, err := client.Recv()
		if err == io.EOF {
			log.Debugw("Elements",
				logging.Stringer("ElementsRequest", request),
				logging.String("State", "complete"))
			return nil
		}
		if err != nil {
			err = errors.ToProto(err)
			log.Warnw("Elements",
				logging.Stringer("ElementsRequest", request),
				logging.Error("Error", err))
			return err
		}
		log.Warnw("Elements",
			logging.Stringer("ElementsResponse", response))
		err = server.Send(response)
		if err != nil {
			err = errors.ToProto(err)
			log.Warnw("Elements",
				logging.Stringer("ElementsRequest", request),
				logging.Error("Error", err))
			return err
		}
	}
}

var _ setv1.SetServer = (*setServer)(nil)
