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
	pages.AddPage("Add", addMoney(), true, false)

	return pages
}

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

	form.AddInputField("Insert money to add: ", "", 20, func(textToCheck string, lastChar rune) bool {
		if unicode.IsDigit(lastChar) || lastChar == '.' {
			return true
		}

		return false
	}, func(text string) {
		_ = text
	})

	form.AddButton("Save money", func() {
		//TO-DO: save the money in the db
	})

	return form
}

func SwitchFocus(form *tview.Form, val int) {
	id, bid := form.GetFocusedItemIndex()

	if bid != -1 {
		form.SetFocus(form.GetFormItemCount() - 2)
	}

	form.SetFocus(id + val)
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
	form.AddInputField("date (format: dd/mm/yyyy) ", "", 20, func(textToCheck string, lastChar rune) bool {
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
	var prova = []string{"ciao", "due"}

	form.AddDropDown("tags", prova, 0, func(option string, optionIndex int) {
		move.Tags = option
	})

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

func menuCreation() *tview.Flex {
	flex := flexCreation()

	flex.AddItem(topBar(), 0, 2, false)
	flex.AddItem(addMoneyUi(), 0, 7, false)

	flex.AddItem(footSet(), 0, 1, false)

	return flex
}

func topBar() *tview.Flex {
	flex := flexCreation()

	//flex.AddItem()

	/*
		money    |    spesi  	| 	a graphics?

	*/

	flex.SetDirection(tview.FlexColumn)

	flex.AddItem(money(), 0, 1, false)
	flex.AddItem(minusMoney(), 0, 1, false)
	flex.AddItem(minusMoney(), 0, 1, false)

	return flex
}

func money() *tview.TextView {
	t := tview.NewTextView()

	return t
}

func minusMoney() *tview.TextView {

	t := tview.NewTextView()

	var m float32

	for _, tot := range sqlScript.Movements {
		m = m + tot.Mov.Money
	}

	t.SetBorder(true)
	t.SetTitle("expenses")
	t.SetText(strconv.FormatFloat(float64(m), 'f', 2, 32))
	t.SetTextAlign(tview.AlignCenter)

	return t

}

func addMoneyUi() *tview.Flex {

	flex := flexCreation()

	max := 5

	if len(sqlScript.Movements) < 5 {
		max = len(sqlScript.Movements)
	}

	for i := max; i > 0; i-- {
		flex.AddItem(writeMoney(sqlScript.Movements[i-1].Mov), 0, 1, false)
	}

	for i := 5 - len(sqlScript.Movements); i > 0; i-- {
		flex.AddItem(tview.NewBox().SetBorder(true), 0, 1, false)
	}

	return flex
}

func writeMoney(mon sqlScript.Movement) *tview.TextView {

	t := tview.NewTextView()

	t.SetBorder(true)

	t.SetText(strconv.FormatFloat(float64(mon.Money), 'f', 2, 32) + "\t \t \t \t \t \t \t \t \t \t \t \t " + mon.Date.Format("02/01/2006") + " \t \t \t \t \t \t \t \t \t \t \t \t " + mon.Tags)

	t.SetDrawFunc(func(screen tcell.Screen, x, y, w, h int) (int, int, int, int) {
		y += h / 2
		return x, y, w, h
	})

	t.SetTextAlign(tview.AlignCenter)

	return t
}

func footSet() *tview.Flex {

	text := tview.NewTextView().SetText("(n) new pay \t (a) add money \t (d) delet \n (" + string(tcell.RuneRArrow) + ") change box forward (in input) \t (" + string(tcell.RuneLArrow) + ") change box backwards (in input)").SetTextColor(tcell.ColorSnow)
	text.SetTextAlign(tview.AlignBottom)
	text.SetTextAlign(tview.AlignCenter)

	text1 := tview.NewTextView().SetText("(q) quit \n (b) back ").SetTextColor(tcell.ColorSnow)
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
