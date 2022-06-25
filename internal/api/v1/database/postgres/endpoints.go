package postgres

import (
	"context"

	v1beta1 "scheduler/pkg/manifest/postgres/v1beta1"

	"github.com/go-kit/kit/endpoint"
)

type CreatePostgresRequest struct {
	Manifest *v1beta1.Postgres `json:"manifest,omitempty"`
}
type CreatePostgresResponse struct {
	Manifest *v1beta1.Postgres `json:"manifest,omitempty"`
}

func CreateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		req := rawRequest.(CreatePostgresRequest)
		svc.Create(req.Manifest)

		return CreatePostgresResponse{
			req.Manifest,
		}, nil
	}
}

type PatchPostgresRequest struct {
	Manifest *v1beta1.Postgres `json:"manifest,omitempty"`
}
type PatchPostgresResponse struct {
	Manifest *v1beta1.Postgres `json:"manifest,omitempty"`
}

func PatchEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		req := rawRequest.(PatchPostgresRequest)
		svc.Patch(req.Manifest)

		return PatchPostgresResponse{
			req.Manifest,
		}, nil
	}
}

type ListPostgresRequest struct {
}
type ListPostgresResponse struct {
	Manifest []v1beta1.Postgres `json:"manifest,omitempty"`
}

func ListEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		nodes, err := svc.List()
		if err != nil {
			return nil, err
		}

		return ListPostgresResponse{
			nodes,
		}, nil
	}
}

type RemovePostgresRequest struct {
	Manifest v1beta1.Postgres `json:"manifest,omitempty"`
}
type RemovePostgresResponse struct {
	Manifest []v1beta1.Postgres `json:"manifest,omitempty"`
}

func RemoveEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		req := rawRequest.(RemovePostgresRequest)

		_, err := svc.Remove(&req.Manifest)
		if err != nil {
			return nil, err
		}

		return ListPostgresResponse{}, nil
	}
}
