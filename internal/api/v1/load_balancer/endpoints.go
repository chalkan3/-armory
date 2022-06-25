package loadbalancer

import (
	"context"

	nodesv1beta1 "scheduler/pkg/manifest/kubernetes/v1beta1"

	"github.com/go-kit/kit/endpoint"
)

type CreateLoadBalancerHARequest struct {
	Manifest *nodesv1beta1.KubeNodes `json:"manifest,omitempty"`
}
type CreateLoadBalancerHAResponse struct {
	Manifest *nodesv1beta1.KubeNodes `json:"manifest,omitempty"`
}

func CreateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		req := rawRequest.(CreateLoadBalancerHARequest)
		svc.Create(req.Manifest)

		return CreateLoadBalancerHAResponse{
			req.Manifest,
		}, nil
	}
}
