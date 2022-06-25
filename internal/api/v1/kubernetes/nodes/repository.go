package nodes

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
	List() ([]nodesv1beta1.KubeNodes, error)
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

	json.Unmarshal(response.Kvs[0].Value, &fetched)
	fetched.Metadata.Revision = response.Header.GetRevision()
	return fetched, nil
}

func (r *repository) List() ([]nodesv1beta1.KubeNodes, error) {
	var nodes []nodesv1beta1.KubeNodes
	response, err := r.etcd.List(context.Background(), "nodes-")
	if err != nil {
		return nil, err
	}

	for _, item := range response.Kvs {
		node := nodesv1beta1.KubeNodes{}
		if err := json.Unmarshal(item.Value, &node); err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}
