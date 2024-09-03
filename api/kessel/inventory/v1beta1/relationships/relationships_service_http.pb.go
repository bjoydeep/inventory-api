// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             (unknown)
// source: kessel/inventory/v1beta1/relationships/relationships_service.proto

package relationships

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationKesselPolicyRelationshipServiceCreatePolicyRelationship = "/kessel.inventory.v1beta1.relationships.KesselPolicyRelationshipService/CreatePolicyRelationship"
const OperationKesselPolicyRelationshipServiceDeleteResourceRelationshipByUrn = "/kessel.inventory.v1beta1.relationships.KesselPolicyRelationshipService/DeleteResourceRelationshipByUrn"
const OperationKesselPolicyRelationshipServiceUpdateResourceRelationshipByUrnHs = "/kessel.inventory.v1beta1.relationships.KesselPolicyRelationshipService/UpdateResourceRelationshipByUrnHs"

type KesselPolicyRelationshipServiceHTTPServer interface {
	CreatePolicyRelationship(context.Context, *CreatePolicyRelationshipRequest) (*CreatePolicyRelationshipResponse, error)
	DeleteResourceRelationshipByUrn(context.Context, *DeleteResourceRelationshipByUrnRequest) (*DeleteResourceRelationshipByUrnResponse, error)
	UpdateResourceRelationshipByUrnHs(context.Context, *UpdateResourceRelationshipByUrnHsRequest) (*UpdateResourceRelationshipByUrnHsResponse, error)
}

func RegisterKesselPolicyRelationshipServiceHTTPServer(s *http.Server, srv KesselPolicyRelationshipServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/inventory/v1beta1/resource-relationships/policyRelationships", _KesselPolicyRelationshipService_CreatePolicyRelationship0_HTTP_Handler(srv))
	r.PUT("/api/inventory/v1beta1/resource-relationships/policyRelationships", _KesselPolicyRelationshipService_UpdateResourceRelationshipByUrnHs0_HTTP_Handler(srv))
	r.DELETE("/api/inventory/v1beta1/resource-relationships/policyRelationships", _KesselPolicyRelationshipService_DeleteResourceRelationshipByUrn0_HTTP_Handler(srv))
}

func _KesselPolicyRelationshipService_CreatePolicyRelationship0_HTTP_Handler(srv KesselPolicyRelationshipServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePolicyRelationshipRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKesselPolicyRelationshipServiceCreatePolicyRelationship)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePolicyRelationship(ctx, req.(*CreatePolicyRelationshipRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreatePolicyRelationshipResponse)
		return ctx.Result(200, reply)
	}
}

func _KesselPolicyRelationshipService_UpdateResourceRelationshipByUrnHs0_HTTP_Handler(srv KesselPolicyRelationshipServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateResourceRelationshipByUrnHsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKesselPolicyRelationshipServiceUpdateResourceRelationshipByUrnHs)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateResourceRelationshipByUrnHs(ctx, req.(*UpdateResourceRelationshipByUrnHsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateResourceRelationshipByUrnHsResponse)
		return ctx.Result(200, reply)
	}
}

func _KesselPolicyRelationshipService_DeleteResourceRelationshipByUrn0_HTTP_Handler(srv KesselPolicyRelationshipServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteResourceRelationshipByUrnRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKesselPolicyRelationshipServiceDeleteResourceRelationshipByUrn)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteResourceRelationshipByUrn(ctx, req.(*DeleteResourceRelationshipByUrnRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteResourceRelationshipByUrnResponse)
		return ctx.Result(200, reply)
	}
}

type KesselPolicyRelationshipServiceHTTPClient interface {
	CreatePolicyRelationship(ctx context.Context, req *CreatePolicyRelationshipRequest, opts ...http.CallOption) (rsp *CreatePolicyRelationshipResponse, err error)
	DeleteResourceRelationshipByUrn(ctx context.Context, req *DeleteResourceRelationshipByUrnRequest, opts ...http.CallOption) (rsp *DeleteResourceRelationshipByUrnResponse, err error)
	UpdateResourceRelationshipByUrnHs(ctx context.Context, req *UpdateResourceRelationshipByUrnHsRequest, opts ...http.CallOption) (rsp *UpdateResourceRelationshipByUrnHsResponse, err error)
}

type KesselPolicyRelationshipServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewKesselPolicyRelationshipServiceHTTPClient(client *http.Client) KesselPolicyRelationshipServiceHTTPClient {
	return &KesselPolicyRelationshipServiceHTTPClientImpl{client}
}

func (c *KesselPolicyRelationshipServiceHTTPClientImpl) CreatePolicyRelationship(ctx context.Context, in *CreatePolicyRelationshipRequest, opts ...http.CallOption) (*CreatePolicyRelationshipResponse, error) {
	var out CreatePolicyRelationshipResponse
	pattern := "/api/inventory/v1beta1/resource-relationships/policyRelationships"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKesselPolicyRelationshipServiceCreatePolicyRelationship))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *KesselPolicyRelationshipServiceHTTPClientImpl) DeleteResourceRelationshipByUrn(ctx context.Context, in *DeleteResourceRelationshipByUrnRequest, opts ...http.CallOption) (*DeleteResourceRelationshipByUrnResponse, error) {
	var out DeleteResourceRelationshipByUrnResponse
	pattern := "/api/inventory/v1beta1/resource-relationships/policyRelationships"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationKesselPolicyRelationshipServiceDeleteResourceRelationshipByUrn))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *KesselPolicyRelationshipServiceHTTPClientImpl) UpdateResourceRelationshipByUrnHs(ctx context.Context, in *UpdateResourceRelationshipByUrnHsRequest, opts ...http.CallOption) (*UpdateResourceRelationshipByUrnHsResponse, error) {
	var out UpdateResourceRelationshipByUrnHsResponse
	pattern := "/api/inventory/v1beta1/resource-relationships/policyRelationships"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKesselPolicyRelationshipServiceUpdateResourceRelationshipByUrnHs))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}