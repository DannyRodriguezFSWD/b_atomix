// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	counterv1 "github.com/atomix/runtime/api/atomix/runtime/atomic/counter/v1"
	"github.com/atomix/runtime/sdk/pkg/runtime"
	"google.golang.org/grpc"
)

const (
	Name       = "AtomicCounter"
	APIVersion = "v1"
)

var Type = runtime.NewType[counterv1.AtomicCounterServer](Name, APIVersion, register, resolve)

func register(server *grpc.Server, delegate *runtime.Delegate[counterv1.AtomicCounterServer]) {
	counterv1.RegisterAtomicCounterServer(server, newAtomicCounterServer(delegate))
}

func resolve(client runtime.Client) (counterv1.AtomicCounterServer, bool) {
	if counter, ok := client.(AtomicCounterProvider); ok {
		return counter.AtomicCounter(), true
	}
	return nil, false
}

type AtomicCounterProvider interface {
	AtomicCounter() counterv1.AtomicCounterServer
}
