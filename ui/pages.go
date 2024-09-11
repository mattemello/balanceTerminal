package ui

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/gdamore/tcell/v2"
	errorhand "github.com/mattemello/balanceTerminal/errorHand"
	"github.com/mattemello/balanceTerminal/sqlScript"
	"github.com/rivo/tview"
)

var pages = tview.NewPages()

func AppCreation() *tview.Application {

	app := tview.NewApplication()
	return app
}

func PageCreation() *tview.Pages {

	textA := tview.NewTextView().SetText("(q) for quit").SetTextColor(tcell.ColorSnow)
	textA.SetBackgroundColor(tcell.ColorBlack)

	pages.AddPage("Main", menuCreation(), true, true)
	pages.AddPage("Insert", insertCreation(), true, false)

	return pages
}

func SwitchFocus(val int) {
	id, bid := form.GetFocusedItemIndex()

	if bid != -1 {
		form.SetFocus(form.GetFormItemCount() - 2)
	}

	form.SetFocus(id + val)
}

var form = tview.NewForm()

func insertCreation() *tview.Form {

	form.SetFieldBackgroundColor(tcell.Color(tcell.ColorValues[12]))
	form.SetFieldTextColor(tcell.ColorSnow)
	form.SetLabelColor(tcell.ColorWhiteSmoke)
	form.SetButtonBackgroundColor(tcell.Color(tcell.ColorValues[12]))

	var move sqlScript.Movement

	form.AddTextView("Insert the money used: \n", "", 0, 1, false, false).SetBorder(true)
	form.AddInputField("money", "", 20, func(textToCheck string, lastChar rune) bool {
		if unicode.IsDigit(lastChar) || lastChar == '.' {
			return true
		}

		return false

	}, func(text string) {
		if text != "" {
			m, err := strconv.ParseFloat(text, 32)
			errorhand.HandlerError(err)
			move.Money = float32(m)
		}
	})
	form.AddInputField("date (format: dd-mm-yyyy) ", "", 20, func(textToCheck string, lastChar rune) bool {
		if unicode.IsDigit(lastChar) {
			return true
		}

		if lastChar == '-' {
			dat := strings.Split(textToCheck, "-")

			if len(dat) > 3 {
				return false
			}

			/*if len(dat) == 4 {
				m, _ := strconv.ParseInt(dat[0], 10, 64)

				if int(m) > 0 && int(m) < 32 {
					return true
				}
			}

			if len(dat) == 3 {
				m, _ := strconv.ParseInt(dat[1], 10, 64)

				if m > 0 && m < 13 {
					return true
				}
			}

			if len(dat) == 4 {
				//m, _ := strconv.ParseInt(dat[0], 10, 64)
				// to do the year

			}*/
		}

		return true

	}, func(text string) {
		move.Date = text
	})
	var prova = []string{"ciao", "due"}

	form.AddDropDown("tags", prova, 0, func(option string, optionIndex int) {
		move.Tags = option
	})

	form.AddButton("save", func() {
		err := sqlScript.SaveValue(move)
		if err != nil {
			errorhand.HandlerError(err)
		} else {
			pages.SwitchToPage("Main")
		}
	})

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

	text := tview.NewTextView().SetText("(n) new pay \t (d) delet \n (" + string(tcell.RuneRArrow) + ") change box forward (in input) \t (" + string(tcell.RuneLArrow) + ") change box backwards (in input)").SetTextColor(tcell.ColorSnow)
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
