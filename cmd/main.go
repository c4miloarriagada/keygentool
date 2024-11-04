package main

import (
	"fmt"
	"keygenpass/internal/domain"
	"keygenpass/internal/infra"
)

func main() {

	repo := infra.NewInMemoryEntityRepository()

	entity := domain.Entities{
		ID:     1,
		Name:   "Test Entity",
		Active: true,
	}

	err := repo.Save(entity)
	if err != nil {
		fmt.Println("Error al guardar la entidad:", err)
	}

	// Obtener la entidad por ID
	retrievedEntity, err := repo.FindByID(1)
	if err != nil {
		fmt.Println("Error al obtener la entidad:", err)
	} else {
		fmt.Println("Entidad obtenida:", retrievedEntity)
	}

	// Obtener todas las entidades
	allEntities, err := repo.FindAll()
	if err != nil {
		fmt.Println("Error al obtener todas las entidades:", err)
	} else {
		fmt.Println("Todas las entidades:", allEntities)
	}

	// Eliminar la entidad por ID
	err = repo.Delete(1)
	if err != nil {
		fmt.Println("Error al eliminar la entidad:", err)
	} else {
		fmt.Println("Entidad eliminada exitosamente.")
	}
}
