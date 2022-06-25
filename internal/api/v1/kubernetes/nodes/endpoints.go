package nodes

import (
	"context"

	nodesv1beta1 "scheduler/pkg/manifest/kubernetes/v1beta1"

	"github.com/go-kit/kit/endpoint"
)

type CreateWorkerNodeRequest struct {
	Manifest *nodesv1beta1.KubeNodes `json:"manifest,omitempty"`
}
type CreateWorkerNodeResponse struct {
	Manifest *nodesv1beta1.KubeNodes `json:"manifest,omitempty"`
}

func CreateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		req := rawRequest.(CreateWorkerNodeRequest)
		svc.Create(req.Manifest)

		return CreateWorkerNodeResponse{
			req.Manifest,
		}, nil
	}
}

type PatchWorkerNodeRequest struct {
	Manifest *nodesv1beta1.KubeNodes `json:"manifest,omitempty"`
}
type PatchWorkerNodeResponse struct {
	Manifest *nodesv1beta1.KubeNodes `json:"manifest,omitempty"`
}

func PatchEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		req := rawRequest.(PatchWorkerNodeRequest)
		svc.Patch(req.Manifest)

		return PatchWorkerNodeResponse{
			req.Manifest,
		}, nil
	}
}

type ListNodeRequest struct {
}
type ListNodeResponse struct {
	Manifest []nodesv1beta1.KubeNodes `json:"manifest,omitempty"`
}

func ListEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		nodes, err := svc.List()
		if err != nil {
			return nil, err
		}

		return ListNodeResponse{
			nodes,
		}, nil
	}
}

type RemoveNodeRequest struct {
	Manifest nodesv1beta1.KubeNodes `json:"manifest,omitempty"`
}
type RemoveNodeResponse struct {
	Manifest []nodesv1beta1.KubeNodes `json:"manifest,omitempty"`
}

func RemoveEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		req := rawRequest.(RemoveNodeRequest)

		_, err := svc.Remove(&req.Manifest)
		if err != nil {
			return nil, err
		}

		return ListNodeResponse{}, nil
	}
}
