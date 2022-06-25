package kong

import (
	"context"

	v1beta1 "scheduler/pkg/manifest/apigateway/v1beta1"

	"github.com/go-kit/kit/endpoint"
)

type CreateKongRequest struct {
	Manifest *v1beta1.Kong `json:"manifest,omitempty"`
}
type CreateKongResponse struct {
	Manifest *v1beta1.Kong `json:"manifest,omitempty"`
}

func CreateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		req := rawRequest.(CreateKongRequest)
		svc.Create(req.Manifest)

		return CreateKongResponse{
			req.Manifest,
		}, nil
	}
}

type PatchKongRequest struct {
	Manifest *v1beta1.Kong `json:"manifest,omitempty"`
}
type PatchKongResponse struct {
	Manifest *v1beta1.Kong `json:"manifest,omitempty"`
}

func PatchEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		req := rawRequest.(PatchKongRequest)
		svc.Patch(req.Manifest)

		return PatchKongResponse{
			req.Manifest,
		}, nil
	}
}

type ListKongRequest struct {
}
type ListKongResponse struct {
	Manifest []v1beta1.Kong `json:"manifest,omitempty"`
}

func ListEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		nodes, err := svc.List()
		if err != nil {
			return nil, err
		}

		return ListKongResponse{
			nodes,
		}, nil
	}
}

type RemoveKongRequest struct {
	Manifest v1beta1.Kong `json:"manifest,omitempty"`
}
type RemoveKongResponse struct {
	Manifest []v1beta1.Kong `json:"manifest,omitempty"`
}

func RemoveEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		req := rawRequest.(RemoveKongRequest)

		_, err := svc.Remove(&req.Manifest)
		if err != nil {
			return nil, err
		}

		return ListKongResponse{}, nil
	}
}
