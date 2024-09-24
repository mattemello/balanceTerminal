package ui

import (
	"github.com/gdamore/tcell/v2"
	errorhand "github.com/mattemello/balanceTerminal/errorHand"
	"github.com/mattemello/balanceTerminal/sqlScript"
	"github.com/rivo/tview"
)

func insertFTags() *tview.Flex {
	f := flexCreation()

	f.AddItem(insertTags(), 0, 9, true)
	f.AddItem(footSet(), 0, 1, false)

	return f
}
func insertTags() *tview.Form {
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

	var fiel string

	form.SetBorder(true).SetTitle("Page to add a new tags")

	//TO-DO: stop the switch of page when press a n

	form.AddInputField("Tag: ", "", 20, nil, func(text string) {
		fiel = text
	})

	form.AddButton("Save tag", func() {
		//TO-DO: controll if the tags its already in
		err := sqlScript.SaveTags(fiel)

		if err != nil {
			errorhand.BadSaving(err)
		} else {
			sqlScript.SaveTag(fiel)
			pages.RemovePage("menu")
			pages.AddAndSwitchToPage("Main", menuCreation(), true)
		}

	})

	return form
}
