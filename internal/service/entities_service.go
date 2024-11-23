package service

import (
	"errors"
	"keygenpass/internal/domain"
)

type EntitiesService interface {
	CreateEntity(id int, name string, active bool, url domain.Url, key domain.Keygen) (*domain.Entities, error)
}

type entitiesService struct{}

func NewEntitiesService() EntitiesService {
	return &entitiesService{}
}

func (s *entitiesService) CreateEntity(id int, name string, active bool, url domain.Url, key domain.Keygen) (*domain.Entities, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	return &domain.Entities{
		ID:     id,
		Name:   name,
		Active: active,
		Url:    url,
		Key:    key,
	}, nil
}
