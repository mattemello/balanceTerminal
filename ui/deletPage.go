package ui

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	errorhand "github.com/mattemello/balanceTerminal/errorHand"
	"github.com/mattemello/balanceTerminal/sqlScript"
	"github.com/rivo/tview"
)

var ToEliminate = make(map[int]bool)

func deletMain() *tview.Flex /*i want to do this?*/ {

	flex := flexCreation()

	flex.AddItem(addMoneywithCheck(), 0, 1, true)

	return flex
}

func createCheck(n int, form *tview.Form) *tview.Form {

	form.AddCheckbox(strconv.Itoa(n), false, func(checked bool) {

		errorhand.Controll(n)
		if checked == true {
			ToEliminate[n] = true
		} else {
			_, contained := ToEliminate[n]

			if !contained {

			} else if ToEliminate[n] {
				ToEliminate[n] = false
			}
		}
	}).SetFieldBackgroundColor(tcell.Color(tcell.ColorValues[12]))

	form.SetFieldTextColor(tcell.Color(tcell.ColorValues[11]))

	return form

}

func addMoneywithCheck() *tview.Flex {

	flex := flexCreation()
	form := tview.NewForm()

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Key() == tcell.KeyEsc {
			pages.SwitchToPage("Main")
		}

		if event.Key() == 257 {
			SwitchFocus(form, 0)
		} else if event.Key() == 256 {
			SwitchFocus(form, -2)
		}

		return event
	})

	max := 10

	if len(sqlScript.Movements) < max {
		max = len(sqlScript.Movements)
	}

	for i := 0; i < max; i++ {
		if i == 0 {

			flex.AddItem(writeMoneywithCheck(sqlScript.Movements[len(sqlScript.Movements)-i-1], form, true), 0, 1, true)
		} else {
			flex.AddItem(writeMoneywithCheck(sqlScript.Movements[len(sqlScript.Movements)-i-1], form, false), 0, 1, false)
		}
	}

	for i := 10 - len(sqlScript.Movements); i > 0; i-- {
		flex.AddItem(tview.NewBox().SetBorder(true), 0, 1, false)
	}

	return flex
}

func writeMoneywithCheck(mon sqlScript.MovementRow, form *tview.Form, hasFocus bool) *tview.Flex {
	flex := tview.NewFlex()

	t := tview.NewTextView()
	t1 := tview.NewTextView()
	t2 := tview.NewTextView()

	flex.SetBorder(true)

	flex.AddItem(createCheck(mon.Id, form), 0, 1, hasFocus)

	t.SetText(strconv.FormatFloat(float64(mon.Mov.Money), 'f', 2, 32))
	t.SetTextAlign(tview.AlignCenter)
	flex.AddItem(t, 0, 1, false)

	t1.SetText(mon.Mov.Date.Format("02/01/2006"))
	t1.SetTextAlign(tview.AlignCenter)
	flex.AddItem(t1, 0, 1, false)

	t2.SetText(mon.Mov.Tags)
	t2.SetTextAlign(tview.AlignCenter)
	flex.AddItem(t2, 0, 1, false)

	t.SetDrawFunc(func(screen tcell.Screen, x, y, w, h int) (int, int, int, int) {
		y += h / 2
		return x, y, w, h
	})
	t1.SetDrawFunc(func(screen tcell.Screen, x, y, w, h int) (int, int, int, int) {
		y += h / 2
		return x, y, w, h
	})
	t2.SetDrawFunc(func(screen tcell.Screen, x, y, w, h int) (int, int, int, int) {
		y += h / 2
		return x, y, w, h
	})

	return flex
}
