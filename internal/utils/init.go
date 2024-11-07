package utils

import (
	"encoding/json"
	"fmt"
	"keygenpass/internal/domain"
	"keygenpass/internal/infra"
	"os"
)

var RESOURCESPATH = "/internal/resources"
var JSONFORMAT = ".json"

func Init() {
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error trying to get path @init:", err)
		return
	}
	path := currentPath + RESOURCESPATH
	loadEntityFromOs(path + "/entity" + JSONFORMAT)
}

func loadEntityFromOs(path string) {
	repo := infra.InMemoryEntityRespository()
	if !FileExists(path) {
		fmt.Println("Entities does not exist.")
		return
	}
	var entities []domain.Entities
	dat, err := os.ReadFile(path)
	CheckError(err)
	unmarshallError := json.Unmarshal(dat, &entities)
	CheckError(unmarshallError)

	for _, entity := range entities {
		repo.Save(entity)
	}

}
