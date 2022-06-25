package redis

import (
	"context"

	v1beta1 "scheduler/pkg/manifest/redis/v1beta1"

	"github.com/go-kit/kit/endpoint"
)

type CreateRedisRequest struct {
	Manifest *v1beta1.Redis `json:"manifest,omitempty"`
}
type CreateRedisResponse struct {
	Manifest *v1beta1.Redis `json:"manifest,omitempty"`
}

func CreateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		req := rawRequest.(CreateRedisRequest)
		svc.Create(req.Manifest)

		return CreateRedisResponse{
			req.Manifest,
		}, nil
	}
}

type PatchRedisRequest struct {
	Manifest *v1beta1.Redis `json:"manifest,omitempty"`
}
type PatchRedisResponse struct {
	Manifest *v1beta1.Redis `json:"manifest,omitempty"`
}

func PatchEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		req := rawRequest.(PatchRedisRequest)
		svc.Patch(req.Manifest)

		return PatchRedisResponse{
			req.Manifest,
		}, nil
	}
}

type ListRedisRequest struct {
}
type ListRedisResponse struct {
	Manifest []v1beta1.Redis `json:"manifest,omitempty"`
}

func ListEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		nodes, err := svc.List()
		if err != nil {
			return nil, err
		}

		return ListRedisResponse{
			nodes,
		}, nil
	}
}

type RemoveRedisRequest struct {
	Manifest v1beta1.Redis `json:"manifest,omitempty"`
}
type RemoveRedisResponse struct {
	Manifest []v1beta1.Redis `json:"manifest,omitempty"`
}

func RemoveEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		req := rawRequest.(RemoveRedisRequest)

		_, err := svc.Remove(&req.Manifest)
		if err != nil {
			return nil, err
		}

		return ListRedisResponse{}, nil
	}
}
