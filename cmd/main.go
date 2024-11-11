package main

import (
	"fmt"
	"keygenpass/internal/handler"
	"keygenpass/internal/service"
	"keygenpass/internal/ui"
	"keygenpass/internal/utils"
	"keygenpass/pkg"

	"github.com/gdamore/tcell/v2"
)

func main() {
	utils.Init()
	defer utils.Close()

	app := pkg.App()

	entitiesService := service.NewEntitiesService()
	entitiesHandler := handler.NewEntitiesHandler(entitiesService)

	navigation := ui.NewNavigation(app)
	navigation.RegisterScreen("addEntity", ui.NewAddEntityUI(app, entitiesHandler))
	navigation.RegisterScreen("listEntities", ui.NewListEntitiesUI(app, entitiesHandler))

	navigation.NavigateTo("addEntity")

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'a':
			navigation.NavigateTo("addEntity")
		case 'l':
			navigation.NavigateTo("listEntities")
		}
		return event
	})

	if err := app.Run(); err != nil {
		fmt.Printf("Error al ejecutar la aplicación: %v\n", err)
	}
}
