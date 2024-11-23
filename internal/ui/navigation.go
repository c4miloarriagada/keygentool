package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Primitive interface {
	Draw(screen tcell.Screen)
	GetRect() (int, int, int, int)
	SetRect(x, y, width, height int)
	InputHandler() func(event *tcell.EventKey, setFocus func(p Primitive))
	Focus(delegate func(p Primitive))
	Blur()
}

type Navigation struct {
	app     *tview.Application
	screens map[string]func() tview.Primitive
}

func NewNavigation(app *tview.Application) *Navigation {
	return &Navigation{
		app:     app,
		screens: make(map[string]func() tview.Primitive),
	}
}

func (n *Navigation) RegisterScreen(name string, generator func() tview.Primitive) {
	n.screens[name] = generator
}

func (n *Navigation) NavigateTo(name string) {
	if generator, ok := n.screens[name]; ok {
		screen := generator()
		n.app.SetRoot(screen, true).SetFocus(screen)
	}
}
