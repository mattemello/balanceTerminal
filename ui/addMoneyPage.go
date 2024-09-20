package ui

import (
	"strconv"
	"unicode"

	"github.com/gdamore/tcell/v2"
	errorhand "github.com/mattemello/balanceTerminal/errorHand"
	"github.com/mattemello/balanceTerminal/sqlScript"

	//"github.com/mattemello/balanceTerminal/errorHand"
	"github.com/rivo/tview"
)

func addMoney() *tview.Form {

	form := tview.NewForm()

	form.SetFieldBackgroundColor(tcell.Color(tcell.ColorValues[12]))
	form.SetFieldTextColor(tcell.ColorSnow)
	form.SetLabelColor(tcell.ColorWhiteSmoke)
	form.SetButtonBackgroundColor(tcell.Color(tcell.ColorValues[12]))

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Key() == 259 {
			SwitchFocus(form, 1)
		} else if event.Key() == 260 {
			SwitchFocus(form, -1)
		}

		return event
	})

	var moneyToAdd float64

	form.AddInputField("Insert money to add: ", "", 20, func(textToCheck string, lastChar rune) bool {
		if unicode.IsDigit(lastChar) || lastChar == '.' {
			return true
		}

		return false
	}, func(text string) {
		moneyToAdd, _ = strconv.ParseFloat(text, 32)

	})

	form.AddButton("Save money", func() {
		err := sqlScript.SaveMoneyDB(float32(moneyToAdd))
		if err != nil {
			errorhand.BadSaving(err)
		} else {
			pages.RemovePage("menu")
			pages.AddAndSwitchToPage("Main", menuCreation(), true)
		}
	})

	return form
}
