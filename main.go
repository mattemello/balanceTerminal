package main

import (
	"github.com/mattemello/balanceTerminal/errorHand"
	"github.com/mattemello/balanceTerminal/sqlScript"
	"github.com/mattemello/balanceTerminal/ui"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func CallErr() {
	ui.PageError()
}

var Db = sqlScript.CreationTable()

func main() {

	fileLog := errorhand.SetLogFile()
	log.SetOutput(fileLog)
	defer fileLog.Close()

	sqlScript.TakeValue()
	sqlScript.QuantityMoney()
	sqlScript.TakeTags()

	var pages = ui.PageCreation()
	app := ui.AppCreation()

	err := app.SetRoot(pages, true).Run()
	errorhand.HandlerError(err, errorhand.TakeFileLine()+" error with the run of the app")

}
