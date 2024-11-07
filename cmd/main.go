package main

import (
	"fmt"
	"keygenpass/internal/infra"
	"keygenpass/internal/utils"
)

func simulateClose() {
	fmt.Println("EXEC CLEAN, SAVE FXS...")
}

func main() {

	utils.Init()

	// utils.Close(simulateClose)
	// fileRepo := infra.NewFileEntityRepository("internal/resources/entity.json")

	// repo := infra.NewInMemoryEntityRepository()
	// entity := domain.Entities{
	// 	ID:     1,
	// 	Name:   "Bank Account",
	// 	Active: true,
	// 	Url:    domain.Url{Name: "https://www.bank.com", HasUrl: true},
	// 	Key:    domain.Keygen{HashedPwd: "test", Active: true, CreatedAt: time.Now()},
	// }
	// entity2 := domain.Entities{
	// 	ID:     2,
	// 	Name:   "Bank Account",
	// 	Active: true,
	// 	Url:    domain.Url{Name: "https://www.bank2.com", HasUrl: true},
	// 	Key:    domain.Keygen{HashedPwd: "test", Active: true, CreatedAt: time.Now()},
	// }

	// if err := repo.Save(entity); err != nil {
	// 	fmt.Println("Error al guardar la entidad:", err)
	// } else {
	// 	fmt.Println("Entidad guardada en memoria y archivo con éxito.")
	// }
	// if err := repo.Save(entity2); err != nil {
	// 	fmt.Println("Error al guardar la entidad:", err)
	// } else {
	// 	fmt.Println("Entidad guardada en memoria y archivo con éxito.")
	// }
	// allRepos, _ := repo.FindAll()
	// fileRepo.Save(allRepos)
	repo := infra.InMemoryEntityRespository()
	fmt.Println(repo.FindAll())
	select {}
}
