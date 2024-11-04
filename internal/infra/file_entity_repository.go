package infra

import (
	"encoding/json"
	"io/ioutil"
	"keygenpass/internal/domain"
)

type FileEntityRepository struct {
	filePath string
}

func NewFileEntityRepository(filePath string) *FileEntityRepository {
	return &FileEntityRepository{filePath: filePath}
}

func (repo *FileEntityRepository) Save(entity domain.Entities) error {
	data, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(repo.filePath, data, 0644)
}
