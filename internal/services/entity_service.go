package services

import "keygenpass/internal/domain"

type EntityService struct {
	repo domain.EntityRepository
}

func NewEntityService(repo domain.EntityRepository) *EntityService {
	return &EntityService{repo: repo}
}

func (s *EntityService) CreateEntity(entity domain.Entities) error {
	return s.repo.Save(entity)
}

func (s *EntityService) GetAllEntities() ([]domain.Entities, error) {
	return s.repo.FindAll()
}
