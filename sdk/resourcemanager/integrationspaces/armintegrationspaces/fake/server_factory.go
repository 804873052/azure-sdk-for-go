//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armintegrationspaces.ClientFactory type.
type ServerFactory struct {
	ApplicationResourcesServer    ApplicationResourcesServer
	ApplicationsServer            ApplicationsServer
	BusinessProcessVersionsServer BusinessProcessVersionsServer
	BusinessProcessesServer       BusinessProcessesServer
	InfrastructureResourcesServer InfrastructureResourcesServer
	OperationsServer              OperationsServer
	SpacesServer                  SpacesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armintegrationspaces.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armintegrationspaces.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                             *ServerFactory
	trMu                            sync.Mutex
	trApplicationResourcesServer    *ApplicationResourcesServerTransport
	trApplicationsServer            *ApplicationsServerTransport
	trBusinessProcessVersionsServer *BusinessProcessVersionsServerTransport
	trBusinessProcessesServer       *BusinessProcessesServerTransport
	trInfrastructureResourcesServer *InfrastructureResourcesServerTransport
	trOperationsServer              *OperationsServerTransport
	trSpacesServer                  *SpacesServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "ApplicationResourcesClient":
		initServer(s, &s.trApplicationResourcesServer, func() *ApplicationResourcesServerTransport {
			return NewApplicationResourcesServerTransport(&s.srv.ApplicationResourcesServer)
		})
		resp, err = s.trApplicationResourcesServer.Do(req)
	case "ApplicationsClient":
		initServer(s, &s.trApplicationsServer, func() *ApplicationsServerTransport { return NewApplicationsServerTransport(&s.srv.ApplicationsServer) })
		resp, err = s.trApplicationsServer.Do(req)
	case "BusinessProcessVersionsClient":
		initServer(s, &s.trBusinessProcessVersionsServer, func() *BusinessProcessVersionsServerTransport {
			return NewBusinessProcessVersionsServerTransport(&s.srv.BusinessProcessVersionsServer)
		})
		resp, err = s.trBusinessProcessVersionsServer.Do(req)
	case "BusinessProcessesClient":
		initServer(s, &s.trBusinessProcessesServer, func() *BusinessProcessesServerTransport {
			return NewBusinessProcessesServerTransport(&s.srv.BusinessProcessesServer)
		})
		resp, err = s.trBusinessProcessesServer.Do(req)
	case "InfrastructureResourcesClient":
		initServer(s, &s.trInfrastructureResourcesServer, func() *InfrastructureResourcesServerTransport {
			return NewInfrastructureResourcesServerTransport(&s.srv.InfrastructureResourcesServer)
		})
		resp, err = s.trInfrastructureResourcesServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "SpacesClient":
		initServer(s, &s.trSpacesServer, func() *SpacesServerTransport { return NewSpacesServerTransport(&s.srv.SpacesServer) })
		resp, err = s.trSpacesServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}