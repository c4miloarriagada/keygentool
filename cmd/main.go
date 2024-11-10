package main

import (
	"fmt"
	"keygenpass/internal/handler"
	"keygenpass/internal/service"
	"keygenpass/internal/ui"
	"keygenpass/internal/utils"
	"keygenpass/pkg"
)

func simulateClose() {
	fmt.Println("EXEC CLEAN, SAVE FXS...")
}

func main() {

	utils.Init()
	app := pkg.App()
	entitiesService := &service.EntitesService{}
	entitiesHandler := handler.NewEntitiesHandler(entitiesService)

	form := ui.NewAddEntityUI(app, entitiesHandler)
	utils.Close()
	if err := app.SetRoot(form, true).Run(); err != nil {
		panic(err)
	}

	select {}
}
