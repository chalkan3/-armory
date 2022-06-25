package v1beta1

import (
	"time"

	"github.com/google/uuid"
)

type Metadata struct {
	ID        string    `json:"id,omitempty" yaml:"id"`
	Name      string    `json:"name,omitempty" yaml:"name"`
	Revision  int64     `json:"revision,omitempty" yaml:"revision"`
	CreatedAT time.Time `json:"created_at,omitempty" yaml:"createdAt"`
}

func (m *Metadata) GenerateID() *Metadata {
	m.ID = "nodes-" + uuid.New().String()
	return m
}
