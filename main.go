package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattemello/balanceTerminal/errorHand"
	"github.com/mattemello/balanceTerminal/sqlScript"
	"github.com/mattemello/balanceTerminal/ui"
	_ "github.com/mattn/go-sqlite3"
)

var pages = ui.PageCreation()

func SwitchPage(s string) {
	pages.SwitchToPage(s)
}

func main() {

	_ = sqlScript.CreationTable()

	app := ui.AppCreation()
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		}
		if event.Rune() == 110 {
			SwitchPage("Insert")
		}

		return event
	})

	err := app.SetRoot(pages, true).Run()
	errorhand.HandlerError(err)

}
