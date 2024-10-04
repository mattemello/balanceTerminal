package ui

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/mattemello/balanceTerminal/sqlScript"
	"github.com/rivo/tview"
)

func menuCreation() *tview.Flex {
	flex := flexCreation()

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 110 {
			pages.AddAndSwitchToPage("Insert", insertFCreation(), true)
		} else if event.Rune() == 97 {
			pages.AddAndSwitchToPage("Add", addFMoney(), true)
		} else if event.Rune() == 98 {
			pages.AddAndSwitchToPage("Main", menuCreation(), true)
		} else if event.Rune() == 116 {
			pages.AddAndSwitchToPage("Tags", insertFTags(), true)
		} else if event.Key() == 257 {
			//TO-DO UP
		} else if event.Key() == 258 {
			//TO-DO DOWN
		}

		if event.Rune() == 113 {
			stopApp()
		}

		return event
	})

	flex.AddItem(topBar(), 0, 2, false)
	flex.AddItem(addMoneyUi(), 0, 7, true)

	flex.AddItem(footSet(), 0, 1, false)

	return flex
}

func topBar() *tview.Flex {
	flex := flexCreation()

	flex.SetDirection(tview.FlexColumn)

	flex.AddItem(money(), 0, 1, false)
	flex.AddItem(minusMoney(), 0, 1, false)

	return flex
}

func addMoneyUi() *tview.Flex {

	flex := flexCreation()

	max := 5

	if len(sqlScript.Movements) < 5 {
		max = len(sqlScript.Movements)
	}

	for i := 0; i < max; i++ {
		flex.AddItem(writeMoney(sqlScript.Movements[len(sqlScript.Movements)-i-1].Mov), 0, 1, true)
	}

	for i := 5 - len(sqlScript.Movements); i > 0; i-- {
		flex.AddItem(tview.NewBox().SetBorder(true), 0, 1, false)
	}

	return flex
}

func footSet() *tview.Flex {

	text := tview.NewTextView().SetText("(n) new pay \t (a) add money \t (d) delet \t (t) add tag \n (" + string(tcell.RuneLArrow) + "/" + string(tcell.RuneRArrow) + ") navigate ").SetTextColor(tcell.ColorSnow)
	text.SetTextAlign(tview.AlignBottom)
	text.SetTextAlign(tview.AlignCenter)

	text1 := tview.NewTextView().SetText("(q) quit").SetTextColor(tcell.ColorSnow)
	text1.SetTextAlign(tview.AlignBottom)

	keyboard := tview.NewFlex()
	keyboard.AddItem(text, 0, 10, false).AddItem(text1, 0, 1, false)
	keyboard.SetBorder(true)

	return keyboard
}

func money() *tview.TextView {
	t := tview.NewTextView()

	t.SetBorder(true)
	t.SetTitle("total money")
	t.SetTextAlign(tview.AlignCenter)
	t.SetSize(1000, 1000)

	t.SetText(strconv.FormatFloat(float64(sqlScript.TotalMoney.Total), 'f', 2, 32))

	return t
}

func minusMoney() *tview.TextView {

	t := tview.NewTextView()

	var m float32

	for _, tot := range sqlScript.Movements {
		if !tot.Mov.Add {
			m = m + tot.Mov.Money
		}
	}

	t.SetBorder(true)
	t.SetTitle("expenses")
	t.SetText(strconv.FormatFloat(float64(m), 'f', 2, 32))
	t.SetTextAlign(tview.AlignCenter)

	return t

}

func writeMoney(mon sqlScript.Movement) *tview.Flex {
	flex := tview.NewFlex()

	t := tview.NewTextView()
	t1 := tview.NewTextView()
	t2 := tview.NewTextView()

	flex.SetBorder(true)

	if mon.Add {
		flex.SetBorderColor(tcell.ColorGreen)
		//t.SetTextColor(tcell.ColorGreen)
	} else {
		flex.SetBorderColor(tcell.ColorRed)
		//t.SetTextColor(tcell.ColorRed)
	}

	flex.AddItem(tview.NewForm().AddCheckbox("", false, nil), 0, 1, true)

	t.SetText(strconv.FormatFloat(float64(mon.Money), 'f', 2, 32))
	t.SetTextAlign(tview.AlignCenter)
	flex.AddItem(t, 0, 1, false)

	t1.SetText(mon.Date.Format("02/01/2006"))
	t1.SetTextAlign(tview.AlignCenter)
	flex.AddItem(t1, 0, 1, false)

	t2.SetText(mon.Tags)
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
