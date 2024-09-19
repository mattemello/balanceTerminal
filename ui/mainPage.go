package ui

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/mattemello/balanceTerminal/sqlScript"
	"github.com/rivo/tview"
)

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
