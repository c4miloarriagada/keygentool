package main

import (
	"fmt"
	"keygenpass/internal/router"

	"keygenpass/internal/utils"
)

func main() {
	utils.Init()
	defer utils.Close()

	app := router.AppRouter()

	if err := app.Run(); err != nil {
		fmt.Printf("Error al ejecutar la aplicación: %v\n", err)
	}
}
