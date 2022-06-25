package loadbalancer

import (
	"context"
	"encoding/json"
	"scheduler/pkg/etcd"
	nodesv1beta1 "scheduler/pkg/manifest/kubernetes/v1beta1"
)

type Repository interface {
	Create(manifest *nodesv1beta1.KubeNodes) error
	Patch(manifest *nodesv1beta1.KubeNodes)
	Remove(manifest *nodesv1beta1.KubeNodes)
	Get(manifest *nodesv1beta1.KubeNodes) (*nodesv1beta1.KubeNodes, error)
}

type repository struct {
	etcd *etcd.ETCD
}

func NewRepository(etcd *etcd.ETCD) Repository {
	return &repository{
		etcd: etcd,
	}
}
func (r *repository) Create(manifest *nodesv1beta1.KubeNodes) error {

	manifestJson, err := json.Marshal(manifest)
	if err != nil {
		return err
	}

	_, err = r.etcd.Put(context.TODO(), manifest.Metadata.ID, string(manifestJson))
	if err != nil {
		return err
	}

	return nil

}
func (r *repository) Patch(manifest *nodesv1beta1.KubeNodes)  {}
func (r *repository) Remove(manifest *nodesv1beta1.KubeNodes) {}
func (r *repository) Get(manifest *nodesv1beta1.KubeNodes) (*nodesv1beta1.KubeNodes, error) {
	var fetched *nodesv1beta1.KubeNodes
	response, err := r.etcd.Get(context.Background(), manifest.Metadata.ID)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response.Kvs[0].Value, &manifest)
	return fetched, nil
}
