package infra

import (
	"encoding/json"
	"keygenpass/internal/domain"
	"os"
)

type FileEntityRepository struct {
	filePath string
}

func NewFileEntityRepository(filePath string) *FileEntityRepository {
	return &FileEntityRepository{filePath: filePath}
}

func (repo *FileEntityRepository) Save(entity []domain.Entities) error {
	data, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	return os.WriteFile(repo.filePath, data, 0644)
}
