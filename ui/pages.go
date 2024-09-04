package ui

import (
	"github.com/gdamore/tcell/v2"
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
	pages.AddPage("Insert", flexCreation().AddItem(tview.NewBox().SetBorder(true), 0, 1, false).AddItem(tview.NewBox().SetBorder(true), 0, 1, false).AddItem(tview.NewBox().SetBorder(true), 0, 1, false), true, false)

	return pages
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
