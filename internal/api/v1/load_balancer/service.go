package loadbalancer

import (
	nodesv1beta1 "scheduler/pkg/manifest/kubernetes/v1beta1"
	"scheduler/pkg/scheduler"
	"time"
)

type Service interface {
	Create(manifest *nodesv1beta1.KubeNodes) (*nodesv1beta1.KubeNodes, error)
	Patch(manifest *nodesv1beta1.KubeNodes) (*nodesv1beta1.KubeNodes, error)
	Remove(manifest *nodesv1beta1.KubeNodes) (*nodesv1beta1.KubeNodes, error)
	Get(manifest *nodesv1beta1.KubeNodes) (*nodesv1beta1.KubeNodes, error)
}

type service struct {
	repository Repository
	scheduler  *scheduler.Scheduler
}

func NewService(repository Repository, scheduler *scheduler.Scheduler) Service {
	return &service{
		repository: repository,
		scheduler:  scheduler,
	}
}
func (s *service) Create(manifest *nodesv1beta1.KubeNodes) (*nodesv1beta1.KubeNodes, error) {
	manifest.Metadata.CreatedAT = time.Now()
	manifest.Metadata.GenerateID()

	err := s.repository.Create(manifest)
	s.scheduler.Schedule("node", "{}", time.Now().Add(30*time.Second))

	return manifest, err
}
func (s *service) Patch(manifest *nodesv1beta1.KubeNodes) (*nodesv1beta1.KubeNodes, error) {
	return s.repository.Get(manifest)
}
func (s *service) Remove(manifest *nodesv1beta1.KubeNodes) (*nodesv1beta1.KubeNodes, error) {
	return s.repository.Get(manifest)
}
func (s *service) Get(manifest *nodesv1beta1.KubeNodes) (*nodesv1beta1.KubeNodes, error) {
	return s.repository.Get(manifest)
}
