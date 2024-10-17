package ui

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/mattemello/balanceTerminal/sqlScript"
	"github.com/rivo/tview"
)

var toEliminate = make(map[int]bool)

func deletMain() *tview.Flex /*i want to do this?*/ {

	flex := flexCreation()

	flex.AddItem(addMoneywithCheck(), 0, 7, true)
	flex.AddItem(footSetDelet(), 0, 1, true)

	return flex
}

func createCheck(form *tview.Form, mon int) {
	form.AddCheckbox("", false, func(checked bool) {
		if checked == true {
			toEliminate[mon] = true
		} else {
			_, contained := toEliminate[mon]

			if !contained {

			} else if toEliminate[mon] {
				toEliminate[mon] = false
			}
		}
	})
}

func addMoneywithCheck() *tview.Flex {

	flex := flexCreation()
	form := tview.NewForm()
	form.SetDrawFunc(func(screen tcell.Screen, x, y, w, h int) (int, int, int, int) {
		y += h / 2
		return x, y, w, h
	})

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Key() == tcell.KeyEsc {
			pages.SwitchToPage("Main")
		} else if event.Key() == 257 {
			SwitchFocus(form, -1)
		} else if event.Key() == 258 {
			SwitchFocus(form, 1)
		} else if event.Rune() == 107 {
			sqlScript.DeletPay(toEliminate)
			pages.RemovePage("menu")
			pages.AddAndSwitchToPage("Main", menuCreation(), true)
		}

		return event
	})

	max := 10

	if len(sqlScript.Movements) < max {
		max = len(sqlScript.Movements)
	}

	for i := 0; i < max; i++ {
		createCheck(form, sqlScript.Movements[len(sqlScript.Movements)-i-1].Id)
	}

	for i := 0; i < max; i++ {
		flex.AddItem(writeMoneywithCheck(sqlScript.Movements[len(sqlScript.Movements)-i-1], form, i), 0, 1, true)
	}

	for i := 10 - len(sqlScript.Movements); i > 0; i-- {
		flex.AddItem(tview.NewBox().SetBorder(true), 0, 1, false)
	}

	return flex
}

func writeMoneywithCheck(mon sqlScript.MovementRow, form *tview.Form, i int) *tview.Flex {
	flex := tview.NewFlex()

	t := tview.NewTextView()
	t1 := tview.NewTextView()
	t2 := tview.NewTextView()

	flex.SetBorder(true)
	form.SetFieldBackgroundColor(tcell.Color(tcell.ColorValues[12])).SetFieldTextColor(tcell.Color(tcell.ColorValues[11]))
	flex.AddItem(form.GetFormItem(i), 0, 1, true)

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

func footSetDelet() *tview.TextView {

	text1 := tview.NewTextView().SetText("(esc) back \n (" + string(tcell.RuneUArrow) + "/" + string(tcell.RuneDArrow) + ") move  \n (k) save").SetTextColor(tcell.ColorSnow)
	text1.SetTextAlign(tview.AlignCenter).SetBorder(true)

	return text1
}
