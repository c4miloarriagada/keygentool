package main

import (
	"fmt"
	"keygenpass/internal/handler"
	"keygenpass/internal/service"
	"keygenpass/internal/ui"
	"keygenpass/internal/utils"
	"keygenpass/pkg"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	utils.Init()
	defer utils.Close()

	app := pkg.App()

	entitiesService := service.NewEntitiesService()
	entitiesHandler := handler.NewEntitiesHandler(entitiesService)

	navigationChannel := make(chan string)
	navigation := ui.NewNavigation(app)

	navigation.RegisterScreen("addEntity", func() tview.Primitive {
		return ui.NewAddEntityUI(app, entitiesHandler, navigationChannel)
	})
	navigation.RegisterScreen("listEntities", func() tview.Primitive {
		return ui.NewListEntitiesUI(app, entitiesHandler)
	})

	go func() {
		for screen := range navigationChannel {
			navigation.NavigateTo(screen)
		}
	}()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case rune(tcell.KeyRight):
			navigationChannel <- "addEntity"
		case rune(tcell.KeyLeft):
			navigationChannel <- "listEntities"
		}
		return event
	})

	navigationChannel <- "addEntity"

	if err := app.Run(); err != nil {
		fmt.Printf("Error al ejecutar la aplicación: %v\n", err)
	}
}
