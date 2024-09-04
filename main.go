package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattemello/balanceTerminal/ui"
)

func main() {

	app := ui.AppCreation()
	pages := ui.PageCreation()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		} else if event.Rune() == 97 {
			pages.SwitchToPage("Insert")
		}

		return event
	})
	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}

}
