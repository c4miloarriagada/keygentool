package router

import (
	"keygenpass/internal/handler"
	"keygenpass/internal/service"
	"keygenpass/internal/ui"
	"keygenpass/pkg"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func AppRouter() *tview.Application {

	navigationChannel := make(chan string)

	app := pkg.App()

	entitiesService := service.NewEntitiesService()

	entitiesHandler := handler.NewEntitiesHandler(entitiesService)

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

	return app
}
