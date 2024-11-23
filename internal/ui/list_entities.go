package ui

import (
	"fmt"
	"keygenpass/internal/handler"
	"keygenpass/internal/infra"

	"github.com/rivo/tview"
)

func NewListEntitiesUI(app *tview.Application, handler *handler.EntitiesHandler) *tview.Table {
	entitiesRepo := infra.InMemoryEntityRespository()
	allEntities, err := entitiesRepo.FindAll()

	table := tview.NewTable()
	table.SetCell(0, 0, tview.NewTableCell("ID").SetSelectable(false))
	table.SetCell(0, 1, tview.NewTableCell("Nombre").SetSelectable(false))
	table.SetCell(0, 2, tview.NewTableCell("Activo").SetSelectable(false))
	table.SetCell(0, 3, tview.NewTableCell("URL").SetSelectable(false))
	if err != nil {
		table.SetCell(1, 0, tview.NewTableCell("Error al cargar entidades"))
		return table
	}

	for i, entity := range allEntities {
		table.SetCell(i+1, 0, tview.NewTableCell(fmt.Sprintf("%d", entity.ID)))
		table.SetCell(i+1, 1, tview.NewTableCell(entity.Name))
		table.SetCell(i+1, 2, tview.NewTableCell(fmt.Sprintf("%t", entity.Active)))
		table.SetCell(i+1, 3, tview.NewTableCell(entity.Url.Name))
	}

	return table
}
