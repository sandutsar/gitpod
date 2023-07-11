// Copyright (c) 2023 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-proxy-gen. DO NOT EDIT.

package v1connect

import (
	context "context"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/gitpod-io/gitpod/components/public-api/go/experimental/v1"
)

var _ OIDCServiceHandler = (*ProxyOIDCServiceHandler)(nil)

type ProxyOIDCServiceHandler struct {
	Client v1.OIDCServiceClient
	UnimplementedOIDCServiceHandler
}

func (s *ProxyOIDCServiceHandler) CreateClientConfig(ctx context.Context, req *connect_go.Request[v1.CreateClientConfigRequest]) (*connect_go.Response[v1.CreateClientConfigResponse], error) {
	resp, err := s.Client.CreateClientConfig(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOIDCServiceHandler) GetClientConfig(ctx context.Context, req *connect_go.Request[v1.GetClientConfigRequest]) (*connect_go.Response[v1.GetClientConfigResponse], error) {
	resp, err := s.Client.GetClientConfig(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOIDCServiceHandler) ListClientConfigs(ctx context.Context, req *connect_go.Request[v1.ListClientConfigsRequest]) (*connect_go.Response[v1.ListClientConfigsResponse], error) {
	resp, err := s.Client.ListClientConfigs(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOIDCServiceHandler) UpdateClientConfig(ctx context.Context, req *connect_go.Request[v1.UpdateClientConfigRequest]) (*connect_go.Response[v1.UpdateClientConfigResponse], error) {
	resp, err := s.Client.UpdateClientConfig(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOIDCServiceHandler) DeleteClientConfig(ctx context.Context, req *connect_go.Request[v1.DeleteClientConfigRequest]) (*connect_go.Response[v1.DeleteClientConfigResponse], error) {
	resp, err := s.Client.DeleteClientConfig(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyOIDCServiceHandler) SetClientConfigActivation(ctx context.Context, req *connect_go.Request[v1.SetClientConfigActivationRequest]) (*connect_go.Response[v1.SetClientConfigActivationResponse], error) {
	resp, err := s.Client.SetClientConfigActivation(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}