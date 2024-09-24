package ui

import (
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/mattemello/balanceTerminal/errorHand"
	"github.com/mattemello/balanceTerminal/sqlScript"
	"github.com/rivo/tview"
)

func insertFCreation() *tview.Flex {
	f := flexCreation()

	f.AddItem(insertCreation(), 0, 9, true)
	f.AddItem(footSet(), 0, 1, true)

	return f
}

func insertCreation() *tview.Form {
	var form = tview.NewForm()

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

	var move sqlScript.Movement

	form.AddTextView("Insert the money used: \n", "", 0, 1, false, false).SetBorder(true)
	form.AddInputField("money", "", 20, func(textToCheck string, lastChar rune) bool {
		if unicode.IsDigit(lastChar) || lastChar == '.' {
			if len(strings.Split(textToCheck, ".")) > 2 {
				return false
			}
			return true
		}

		return false

	}, func(text string) {
		if text != "" {
			m, err := strconv.ParseFloat(text, 32)
			errorhand.HandlerError(err, errorhand.TakeFileLine()+"  error with the parseFloat")
			move.Money = float32(m)
		}
	})
	form.AddInputField("date (format: dd/mm/yyyy) ", time.Now().Format("02/01/2006"), 20, func(textToCheck string, lastChar rune) bool {
		if unicode.IsDigit(lastChar) {
			return true
		}

		if lastChar == '-' {
			dat := strings.Split(textToCheck, "-")

			if len(dat) > 3 {
				return false
			}

			if len(dat) == 2 {
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

			//if len(dat) == 4 {
			//m, _ := strconv.ParseInt(dat[0], 10, 64)
			// to do the year

			//}
		}

		return true

	}, func(text string) {
		if len(text) > 9 {
			var err error
			move.Date, err = time.Parse("02/01/2006", text)
			errorhand.HandlerError(err, errorhand.TakeFileLine()+" error in the parse of the date")
		}
	})

	form.AddDropDown("tags", sqlScript.AllTags, 0, func(option string, optionIndex int) {
		move.Tags = option
	})

	move.Add = false

	form.AddButton("save", func() {
		err := sqlScript.SaveTransaction(move)
		if err != nil {
			errorhand.BadSaving(err)
		} else {
			sqlScript.SaveMove(move)
			pages.RemovePage("menu")
			pages.AddAndSwitchToPage("Main", menuCreation(), true)
		}
	})

	return form

}
