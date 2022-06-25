package redis

import (
	v1beta1 "scheduler/pkg/manifest/redis/v1beta1"
	"scheduler/pkg/scheduler"
	"time"
)

type Service interface {
	Create(manifest *v1beta1.Redis) (*v1beta1.Redis, error)
	Patch(manifest *v1beta1.Redis) (*v1beta1.Redis, error)
	Remove(manifest *v1beta1.Redis) (*v1beta1.Redis, error)
	Get(manifest *v1beta1.Redis) (*v1beta1.Redis, error)
	List() ([]v1beta1.Redis, error)
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
func (s *service) Create(manifest *v1beta1.Redis) (*v1beta1.Redis, error) {
	manifest.Metadata.CreatedAT = time.Now()
	manifest.Metadata.GenerateID()

	err := s.repository.Create(manifest)

	s.scheduler.Schedule("create-redis", manifest, time.Now().Add(30*time.Second))

	return manifest, err
}
func (s *service) Patch(manifest *v1beta1.Redis) (*v1beta1.Redis, error) {
	return s.repository.Get(manifest)
}
func (s *service) Remove(manifest *v1beta1.Redis) (*v1beta1.Redis, error) {

	fetchedManifest, err := s.repository.Get(manifest)
	if err != nil {
		return nil, err
	}
	s.scheduler.Schedule("destroy-redis", fetchedManifest, time.Now().Add(30*time.Second))
	return fetchedManifest, nil
}
func (s *service) Get(manifest *v1beta1.Redis) (*v1beta1.Redis, error) {
	return s.repository.Get(manifest)
}

func (s *service) List() ([]v1beta1.Redis, error) {
	return s.repository.List()
}
