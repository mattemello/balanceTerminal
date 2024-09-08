package ui

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/gdamore/tcell/v2"
	errorhand "github.com/mattemello/balanceTerminal/errorHand"
	"github.com/mattemello/balanceTerminal/sqlScript"
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

	pages.AddPage("Menu", menuCreation(), true, true)
	pages.AddPage("Insert", insertCreation(), true, false)

	return pages
}

func insertCreation() *tview.Form {
	//flex := flexCreation()
	form := tview.NewForm()

	var move sqlScript.Movement

	form.AddTextView("Insert the money used: \n", "", 0, 1, false, false).SetBorder(true)
	form.AddInputField("money", "", 20, func(textToCheck string, lastChar rune) bool {
		if unicode.IsDigit(lastChar) || lastChar == '.' {
			return true
		}

		return false

	}, func(text string) {
		m, err := strconv.ParseFloat(text, 32)
		errorhand.HandlerError(err)
		move.Money = float32(m)
	})
	form.AddInputField("date (format: dd/mm/yyyy) ", "", 4, func(textToCheck string, lastChar rune) bool {
		if !unicode.IsDigit(lastChar) && lastChar != '/' {
			return false
		}

		if lastChar == '/' {
			dat := strings.Split(textToCheck, "/")

			if len(dat) > 3 {
				return false
			}

			if len(dat) == 1 {
				m, _ := strconv.ParseInt(dat[0], 10, 64)

				if m < 0 || m > 31 {
					return false
				}
			}

			if len(dat) == 2 {
				m, _ := strconv.ParseInt(dat[1], 10, 64)

				if m < 0 || m > 12 {
					return false
				}
			}

			if len(dat) == 3 {
				//m, _ := strconv.ParseInt(dat[0], 10, 64)
				// to do the year

			}
		}

		return true

	}, func(text string) {
		move.Date = text
	})
	var prova = []string{"ciao", "due"}

	form.AddDropDown("money", prova, 0, func(option string, optionIndex int) {
		move.Tags = option
	})

	//flex.AddItem(form, 0, 8, false)
	//flex.AddItem(footSet(), 0, 1, false)

	return form

}

func menuCreation() *tview.Flex {
	flex := flexCreation()

	flex.AddItem(tview.NewBox().SetBorder(true), 0, 2, false)
	flex.AddItem(tview.NewBox().SetBorder(true), 0, 7, false)

	flex.AddItem(footSet(), 0, 1, false)

	return flex
}

func footSet() *tview.Flex {

	text := tview.NewTextView().SetText("(n) new pay \t (d) delet").SetTextColor(tcell.ColorSnow)
	text.SetTextAlign(tview.AlignBottom)
	text.SetTextAlign(tview.AlignCenter)

	text1 := tview.NewTextView().SetText("(q) quit").SetTextColor(tcell.ColorSnow)
	text1.SetTextAlign(tview.AlignBottom)

	keyboard := tview.NewFlex()
	keyboard.AddItem(text, 0, 10, false).AddItem(text1, 0, 1, false)
	keyboard.SetBorder(true)

	return keyboard
}

/*
for _, dimension := range dimensions {
	flex.AddItem(textA, 0, dimension, false)
}*/

func flexCreation() *tview.Flex {

	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	return flex

}
