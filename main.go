package main

import (
	"database/sql"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/mattemello/balanceTerminal/ui"
	"github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "money.db")
	if err != nil {
		log.Fatal("Error with the database: ", err)
		return
	}

	queryCre := "CREAT "

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
