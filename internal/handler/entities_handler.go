package handler

import (
	"keygenpass/internal/domain"
	"keygenpass/internal/service"
	"time"
)

type EntitiesHandler struct {
	service service.EntitiesService
}

func NewEntitiesHandler(service service.EntitiesService) *EntitiesHandler {
	return &EntitiesHandler{service: service}
}

func (h *EntitiesHandler) CreateEntity(id int, name string, active bool, urlName string, hashedPwd string, keyActive bool) (*domain.Entities, error) {
	url := domain.Url{
		Name:   urlName,
		HasUrl: urlName != "",
	}
	key := domain.Keygen{
		HashedPwd: hashedPwd,
		Active:    keyActive,
		CreatedAt: time.Now(),
	}

	return h.service.CreateEntity(id, name, active, url, key)
}
