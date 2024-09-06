package ui

import (
	"unicode"

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
	pages.AddPage("Insert", insertCreation(), true, false)

	return pages
}

func insertCreation() *tview.Flex {
	flex := flexCreation()
	form := tview.NewForm()

	form.AddTextView("Insert the money used: \n", "", 0, 1, false, false).SetBorder(true)
	form.AddTextArea("money", "00.00", 0, 1, 1, func(text string) {
		for _, c := range text {
			if unicode.IsDigit(c) {
				//do here
			}
		}
	})

	flex.AddItem(form, 0, 8, false)
	flex.AddItem(footSet(), 0, 1, false)

	return flex

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
