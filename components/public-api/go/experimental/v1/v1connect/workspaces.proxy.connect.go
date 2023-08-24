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

var _ WorkspacesServiceHandler = (*ProxyWorkspacesServiceHandler)(nil)

type ProxyWorkspacesServiceHandler struct {
	Client v1.WorkspacesServiceClient
	UnimplementedWorkspacesServiceHandler
}

func (s *ProxyWorkspacesServiceHandler) ListWorkspaces(ctx context.Context, req *connect_go.Request[v1.ListWorkspacesRequest]) (*connect_go.Response[v1.ListWorkspacesResponse], error) {
	resp, err := s.Client.ListWorkspaces(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyWorkspacesServiceHandler) GetWorkspace(ctx context.Context, req *connect_go.Request[v1.GetWorkspaceRequest]) (*connect_go.Response[v1.GetWorkspaceResponse], error) {
	resp, err := s.Client.GetWorkspace(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyWorkspacesServiceHandler) GetOwnerToken(ctx context.Context, req *connect_go.Request[v1.GetOwnerTokenRequest]) (*connect_go.Response[v1.GetOwnerTokenResponse], error) {
	resp, err := s.Client.GetOwnerToken(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyWorkspacesServiceHandler) CreateAndStartWorkspace(ctx context.Context, req *connect_go.Request[v1.CreateAndStartWorkspaceRequest]) (*connect_go.Response[v1.CreateAndStartWorkspaceResponse], error) {
	resp, err := s.Client.CreateAndStartWorkspace(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyWorkspacesServiceHandler) StartWorkspace(ctx context.Context, req *connect_go.Request[v1.StartWorkspaceRequest]) (*connect_go.Response[v1.StartWorkspaceResponse], error) {
	resp, err := s.Client.StartWorkspace(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyWorkspacesServiceHandler) StopWorkspace(ctx context.Context, req *connect_go.Request[v1.StopWorkspaceRequest]) (*connect_go.Response[v1.StopWorkspaceResponse], error) {
	resp, err := s.Client.StopWorkspace(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyWorkspacesServiceHandler) DeleteWorkspace(ctx context.Context, req *connect_go.Request[v1.DeleteWorkspaceRequest]) (*connect_go.Response[v1.DeleteWorkspaceResponse], error) {
	resp, err := s.Client.DeleteWorkspace(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyWorkspacesServiceHandler) UpdatePort(ctx context.Context, req *connect_go.Request[v1.UpdatePortRequest]) (*connect_go.Response[v1.UpdatePortResponse], error) {
	resp, err := s.Client.UpdatePort(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}

func (s *ProxyWorkspacesServiceHandler) ListWorkspaceClasses(ctx context.Context, req *connect_go.Request[v1.ListWorkspaceClassesRequest]) (*connect_go.Response[v1.ListWorkspaceClassesResponse], error) {
	resp, err := s.Client.ListWorkspaceClasses(ctx, req.Msg)
	if err != nil {
		// TODO(milan): Convert to correct status code
		return nil, err
	}

	return connect_go.NewResponse(resp), nil
}
