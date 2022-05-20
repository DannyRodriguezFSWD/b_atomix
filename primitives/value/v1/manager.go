// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	"github.com/atomix/runtime/api/atomix/value/v1"
	"github.com/atomix/runtime/pkg/errors"
	"github.com/atomix/runtime/pkg/primitive"
)

func newValueV1ManagerServer(proxies *primitive.Service[Value]) v1.ValueManagerServer {
	return &valueV1ManagerServer{
		proxies: proxies,
	}
}

type valueV1ManagerServer struct {
	proxies *primitive.Service[Value]
}

func (s *valueV1ManagerServer) Create(ctx context.Context, request *v1.CreateRequest) (*v1.CreateResponse, error) {
	conn, err := s.proxies.Connect(ctx, request.Primitive)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = conn.Create(ctx, request.Primitive.PrimitiveID)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &v1.CreateResponse{}, nil
}

func (s *valueV1ManagerServer) Close(ctx context.Context, request *v1.CloseRequest) (*v1.CloseResponse, error) {
	conn, err := s.proxies.GetConn(request.PrimitiveID)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = conn.Close(ctx, request.PrimitiveID)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &v1.CloseResponse{}, nil
}

var _ v1.ValueManagerServer = (*valueV1ManagerServer)(nil)
