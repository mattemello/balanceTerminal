package main

import (
	"database/sql"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/mattemello/balanceTerminal/ui"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "money.db")
	if err != nil {
		log.Fatal("Error with the database: ", err)
		return
	}

	queryCre := "CREATE TABLE IF NOT EXISTS spendingMoney( id_transition INTEGER PRIMARY KEY, valur_tansaction REAL, tags TEXT, date DATE )"

	if _, err := db.Exec(queryCre); err != nil {
		log.Fatal("Error with the database: ", err)
		return
	}

	/*_, err = db.Exec("INSERT INTO spendingMoney VALUES(0, 0.5, 'real', 05/12/2021);")
	if err != nil {
		log.Fatal("29 Error with the database: ", err)
		return
	}

	_, err = db.Exec("DELETE FROM spendingMoney WHERE id_transition=0;")
	if err != nil {
		log.Fatal("Error with the database: ", err)
		return
	}*/

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
