package service

import (
	"errors"
	"keygenpass/internal/domain"
)

type EntitesService struct{}

func (s *EntitesService) CreateEntity(id int, name string, active bool, url domain.Url, key domain.Keygen) (*domain.Entities, error) {

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	return &domain.Entities{ID: id,
		Name:   name,
		Active: active,
		Url:    url,
		Key:    key,
	}, nil
}
