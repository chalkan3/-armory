package kong

import (
	"context"
	"encoding/json"
	"scheduler/pkg/etcd"
	v1beta1 "scheduler/pkg/manifest/apigateway/v1beta1"
)

type Repository interface {
	Create(manifest *v1beta1.Kong) error
	Patch(manifest *v1beta1.Kong)
	Remove(manifest *v1beta1.Kong)
	Get(manifest *v1beta1.Kong) (*v1beta1.Kong, error)
	List() ([]v1beta1.Kong, error)
}

type repository struct {
	etcd *etcd.ETCD
}

func NewRepository(etcd *etcd.ETCD) Repository {
	return &repository{
		etcd: etcd,
	}
}
func (r *repository) Create(manifest *v1beta1.Kong) error {

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
func (r *repository) Patch(manifest *v1beta1.Kong)  {}
func (r *repository) Remove(manifest *v1beta1.Kong) {}
func (r *repository) Get(manifest *v1beta1.Kong) (*v1beta1.Kong, error) {
	var fetched *v1beta1.Kong
	response, err := r.etcd.Get(context.Background(), manifest.Metadata.ID)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(response.Kvs[0].Value, &fetched)
	fetched.Metadata.Revision = response.Header.GetRevision()
	return fetched, nil
}

func (r *repository) List() ([]v1beta1.Kong, error) {
	var nodes []v1beta1.Kong
	response, err := r.etcd.List(context.Background(), "postgres-")
	if err != nil {
		return nil, err
	}

	for _, item := range response.Kvs {
		node := v1beta1.Kong{}
		if err := json.Unmarshal(item.Value, &node); err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}
