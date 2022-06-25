package nodes

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
	List() ([]nodesv1beta1.KubeNodes, error)
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

	if manifest.Spec.Node.Primary && manifest.Spec.Node.Types == "master" {
		s.scheduler.Schedule("create-node-master-primary", manifest, time.Now().Add(30*time.Second))
		return manifest, err
	}

	if manifest.Spec.Node.Types == "master" {
		s.scheduler.Schedule("create-node-master", manifest, time.Now().Add(30*time.Second))
		return manifest, err

	}

	s.scheduler.Schedule("create-node-worker", manifest, time.Now().Add(30*time.Second))

	return manifest, err
}
func (s *service) Patch(manifest *nodesv1beta1.KubeNodes) (*nodesv1beta1.KubeNodes, error) {
	return s.repository.Get(manifest)
}
func (s *service) Remove(manifest *nodesv1beta1.KubeNodes) (*nodesv1beta1.KubeNodes, error) {

	fetchedManifest, err := s.repository.Get(manifest)
	if err != nil {
		return nil, err
	}
	s.scheduler.Schedule("destroy-node", fetchedManifest, time.Now().Add(30*time.Second))
	return fetchedManifest, nil
}
func (s *service) Get(manifest *nodesv1beta1.KubeNodes) (*nodesv1beta1.KubeNodes, error) {
	return s.repository.Get(manifest)
}

func (s *service) List() ([]nodesv1beta1.KubeNodes, error) {
	return s.repository.List()
}
