package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func AppCreation() *tview.Application {

	app := tview.NewApplication()
	return app
}

func PageCreation() *tview.Pages {

	pages := tview.NewPages()
	textA := tview.NewTextView().SetText("(q) for quit").SetTextColor(tcell.ColorSnow)
	textA.SetBackgroundColor(tcell.ColorBlack)

	pages.AddPage("Menu", flexCreation().AddItem(tview.NewBox().SetBorder(true), 0, 1, false).AddItem(textA, 0, 4, false), true, true)
	pages.AddPage("Insert", flexCreation().AddItem(tview.NewBox().SetBorder(true), 0, 1, false).AddItem(tview.NewBox().SetBorder(true), 0, 1, false).AddItem(tview.NewBox().SetBorder(true), 0, 1, false), true, false)

	return pages
}

/*
for _, dimension := range dimensions {
	flex.AddItem(textA, 0, dimension, false)
}*/

func flexCreation() *tview.Flex {

	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	return flex

}
