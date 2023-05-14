package container

import (
	"template/pkg/domain"
)

type server struct {
}

func NewRepository() *server {
	return &server{}
}

func (s *server) GetDomain() ([]*domain.Entity, error) {
	return nil, nil
}

func (s *server) CreateDomain() ([]*domain.Entity, error) {
	return nil, nil
}

func (s server) UpdateDomain() ([]*domain.Entity, error) {
	return nil, nil
}

func (s *server) DeleteDomain() ([]*domain.Entity, error) {
	return nil, nil
}
