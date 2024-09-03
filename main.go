package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		}

		return event
	})

	pages := tview.NewPages()

	textA := tview.NewTextView().SetText("(q) for quit") //.SetBackgroundColor(tcell.ColorBlack)

	pages.AddPage("Menu", textA, true, true)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}

}
