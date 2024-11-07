package infra

import (
	"errors"
	"keygenpass/internal/domain"
	"sync"
)

type InMemoryEntityRepository struct {
	entities map[int]domain.Entities
}

var (
	instance  *InMemoryEntityRepository
	once      sync.Once
	waitGroup sync.WaitGroup
)

func InMemoryEntityRespository() *InMemoryEntityRepository {
	once.Do(func() {
		waitGroup.Add(1)
		instance = &InMemoryEntityRepository{entities: make(map[int]domain.Entities)}
		waitGroup.Done()
	})
	waitGroup.Wait()
	return instance
}

func (repo *InMemoryEntityRepository) Save(entity domain.Entities) error {
	repo.entities[entity.ID] = entity
	return nil
}

func (repo *InMemoryEntityRepository) FindByID(id int) (domain.Entities, error) {
	entity, exists := repo.entities[id]
	if !exists {
		return domain.Entities{}, errors.New("entity not found")
	}
	return entity, nil
}

func (repo *InMemoryEntityRepository) FindAll() ([]domain.Entities, error) {
	var allEntities []domain.Entities
	for _, entity := range repo.entities {
		allEntities = append(allEntities, entity)
	}
	return allEntities, nil
}

func (repo *InMemoryEntityRepository) Delete(id int) error {
	if _, exists := repo.entities[id]; !exists {
		return errors.New("entity not found")
	}
	delete(repo.entities, id)
	return nil
}
