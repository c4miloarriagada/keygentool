package main

import (
	"fmt"
	"keygenpass/internal/router"

	"keygenpass/internal/utils"
)

func main() {
	hookFunctions()

	app := router.AppRouter()

	if err := app.Run(); err != nil {
		fmt.Printf("Error al ejecutar la aplicación: %v\n", err)
	}
}

func hookFunctions() {
	utils.Init()
	defer utils.Close()

}
