package ui

import (
	"github.com/gdamore/tcell/v2"
	//errorhand "github.com/mattemello/balanceTerminal/errorHand"
	"github.com/rivo/tview"
)

var pages = tview.NewPages()

func AppCreation() *tview.Application {

	app := tview.NewApplication()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 81 {
			app.Stop()
		} else if event.Rune() == 78 {
			pages.AddAndSwitchToPage("Insert", insertFCreation(), true)
		} else if event.Rune() == 65 {
			pages.AddAndSwitchToPage("Add", addFMoney(), true)
		} else if event.Rune() == 66 {
			pages.AddAndSwitchToPage("Main", menuCreation(), true)
		} else if event.Rune() == 84 {
			pages.AddAndSwitchToPage("Tags", insertFTags(), true)
		}

		return event
	})
	return app
}

func PageCreation() *tview.Pages {

	textA := tview.NewTextView().SetText("(q) for quit").SetTextColor(tcell.ColorSnow)
	textA.SetBackgroundColor(tcell.ColorBlack)

	pages.AddPage("Main", menuCreation(), true, true)

	return pages
}

func SwitchFocus(form *tview.Form, val int) {
	id, bid := form.GetFocusedItemIndex()

	if bid != -1 {
		//errorhand.Controll(form.GetFormItemIndex("Insert money to add: "))
		form.SetFocus(form.GetFormItemCount() - 2)
	}

	form.SetFocus(id + val)
}

func flexCreation() *tview.Flex {

	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	return flex

}
