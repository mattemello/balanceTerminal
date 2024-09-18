package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattemello/balanceTerminal/errorHand"
	"github.com/mattemello/balanceTerminal/sqlScript"
	"github.com/mattemello/balanceTerminal/ui"
	_ "github.com/mattn/go-sqlite3"
)

var Db = sqlScript.CreationTable()

func main() {

	sqlScript.TakeValue()

	var pages = ui.PageCreation()
	app := ui.AppCreation()
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		} else if event.Rune() == 110 {
			pages.SwitchToPage("Insert")
		} else if event.Key() == 259 {
			ui.SwitchFocus(1)
		} else if event.Key() == 260 {
			ui.SwitchFocus(-1)
		} else if event.Rune() == 97 {
			pages.SwitchToPage("Add")
		}

		return event
	})

	err := app.SetRoot(pages, true).Run()
	errorhand.HandlerError(err, errorhand.TakeFileLine()+" error with the run of the app")

}
