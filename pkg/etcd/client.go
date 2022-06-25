package etcd

import (
	"context"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type ETCD struct {
	client *clientv3.Client
}

func NewETCD() *ETCD {
	return new(ETCD)
}

func (etcd *ETCD) Connection(timeout time.Duration, endpoints ...string) (*ETCD, error) {
	cli, err := clientv3.New(clientv3.Config{
		DialTimeout: timeout,
		Endpoints:   endpoints,
	})
	if err != nil {
		return nil, err
	}

	etcd.client = cli

	return etcd, nil
}

func (etcd *ETCD) Client() *clientv3.Client {
	return etcd.client
}

func (etcd *ETCD) Put(ctx context.Context, k string, v string) (*clientv3.PutResponse, error) {
	return etcd.client.Put(ctx, k, v)
}

func (etcd *ETCD) Get(ctx context.Context, k string) (*clientv3.GetResponse, error) {
	return etcd.client.Get(ctx, k)
}

func (etcd *ETCD) List(ctx context.Context, prefix string) (*clientv3.GetResponse, error) {
	opts := []clientv3.OpOption{
		clientv3.WithPrefix(),
	}
	return etcd.client.Get(ctx, prefix, opts...)
}
