package postgres

import (
	v1beta1 "scheduler/pkg/manifest/postgres/v1beta1"
	"scheduler/pkg/scheduler"
	"time"
)

type Service interface {
	Create(manifest *v1beta1.Postgres) (*v1beta1.Postgres, error)
	Patch(manifest *v1beta1.Postgres) (*v1beta1.Postgres, error)
	Remove(manifest *v1beta1.Postgres) (*v1beta1.Postgres, error)
	Get(manifest *v1beta1.Postgres) (*v1beta1.Postgres, error)
	List() ([]v1beta1.Postgres, error)
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
func (s *service) Create(manifest *v1beta1.Postgres) (*v1beta1.Postgres, error) {
	manifest.Metadata.CreatedAT = time.Now()
	manifest.Metadata.GenerateID()

	err := s.repository.Create(manifest)

	s.scheduler.Schedule("create-postgres", manifest, time.Now().Add(30*time.Second))

	return manifest, err
}
func (s *service) Patch(manifest *v1beta1.Postgres) (*v1beta1.Postgres, error) {
	return s.repository.Get(manifest)
}
func (s *service) Remove(manifest *v1beta1.Postgres) (*v1beta1.Postgres, error) {

	fetchedManifest, err := s.repository.Get(manifest)
	if err != nil {
		return nil, err
	}
	s.scheduler.Schedule("destroy-postgres", fetchedManifest, time.Now().Add(30*time.Second))
	return fetchedManifest, nil
}
func (s *service) Get(manifest *v1beta1.Postgres) (*v1beta1.Postgres, error) {
	return s.repository.Get(manifest)
}

func (s *service) List() ([]v1beta1.Postgres, error) {
	return s.repository.List()
}
