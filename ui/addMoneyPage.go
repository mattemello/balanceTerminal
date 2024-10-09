package ui

import (
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gdamore/tcell/v2"
	errorhand "github.com/mattemello/balanceTerminal/errorHand"
	"github.com/mattemello/balanceTerminal/sqlScript"

	"github.com/rivo/tview"
)

func addFMoney() *tview.Flex {
	f := flexCreation()

	f.AddItem(addMoney(), 0, 9, true)
	f.AddItem(footSetAdd(), 0, 1, false)

	return f
}

func footSetAdd() *tview.TextView {

	text1 := tview.NewTextView().SetText("(esc) back \n (" + string(tcell.RuneUArrow) + "/" + string(tcell.RuneDArrow) + ") move  ").SetTextColor(tcell.ColorSnow)
	text1.SetTextAlign(tview.AlignCenter).SetBorder(true)

	return text1
}

func addMoney() *tview.Form {

	form := tview.NewForm()

	form.SetFieldBackgroundColor(tcell.Color(tcell.ColorValues[12]))
	form.SetFieldTextColor(tcell.ColorSnow)
	form.SetLabelColor(tcell.ColorWhiteSmoke)
	form.SetButtonBackgroundColor(tcell.Color(tcell.ColorValues[12]))

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Key() == tcell.KeyEsc {
			pages.SwitchToPage("Main")
		}

		if event.Key() == 258 {
			SwitchFocus(form, 1)
		} else if event.Key() == 257 {
			SwitchFocus(form, -1)
		}

		return event
	})

	var moneyToAdd float64

	form.SetBorder(true).SetTitle("Page to add the money")

	form.AddInputField("Insert money to add: ", "", 20, func(textToCheck string, lastChar rune) bool {
		if unicode.IsDigit(lastChar) || lastChar == '.' {
			if len(strings.Split(textToCheck, ".")) > 2 {
				return false
			}
			return true
		}

		return false
	}, func(text string) {
		moneyToAdd, _ = strconv.ParseFloat(text, 32)

	})

	form.AddButton("Save money", func() {
		var mv sqlScript.Movement

		mv.Money = float32(moneyToAdd)
		mv.Date = time.Now()
		mv.Tags = ""
		mv.Add = true

		err := sqlScript.SaveTransaction(mv)
		if err != nil {
			errorhand.BadSaving(err)
			pages.AddAndSwitchToPage("err", PageError("Error whit the save, controll the log file"), true)
		} else {
			sqlScript.SaveMove(mv)
			pages.RemovePage("menu")
			pages.AddAndSwitchToPage("Main", menuCreation(), true)
		}
	})

	return form
}
