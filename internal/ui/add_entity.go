package ui

import (
	"fmt"
	"keygenpass/internal/handler"
	"keygenpass/internal/infra"
	"strconv"

	"github.com/rivo/tview"
)

func NewAddEntityUI(app *tview.Application, handler *handler.EntitiesHandler, navigationChannel chan string) *tview.Form {
	var id int
	var name, urlName, hashedPwd string
	var active, keyActive bool
	repo := infra.InMemoryEntityRespository()
	form := tview.NewForm().
		AddInputField("ID", "", 20, nil, func(text string) {
			id, _ = strconv.Atoi(text)
		}).
		AddInputField("Name", "", 20, nil, func(text string) {
			name = text
		}).
		AddCheckbox("Active", false, func(checked bool) {
			active = checked
		}).
		AddInputField("URL", "", 40, nil, func(text string) {
			urlName = text
		}).
		AddInputField("Hashed Password", "", 40, nil, func(text string) {
			hashedPwd = text
		}).
		AddCheckbox("Key Active", false, func(checked bool) {
			keyActive = checked
		}).
		AddButton("Save", func() {
			entity, err := handler.CreateEntity(id, name, active, urlName, hashedPwd, keyActive)

			if err != nil {

				modal := tview.NewModal().
					SetText("Error al crear la entidad: " + err.Error()).
					AddButtons([]string{"OK"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						navigationChannel <- "addEntity"
					})
				app.SetRoot(modal, true)
			} else {
				todo, _ := repo.FindAll()
				repo.Save(*entity)
				modal := tview.NewModal().
					SetText("Entidad creada exitosamente:\n" + fmt.Sprintf("%+v", todo)).
					AddButtons([]string{"OK"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						navigationChannel <- "listEntities"
					})
				app.SetRoot(modal, true)
			}
		}).
		AddButton("Cancel", func() {
			app.Stop()
		})

	form.SetBorder(true).SetTitle("Add New Entity").SetTitleAlign(tview.AlignLeft)
	return form
}
