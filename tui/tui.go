package tui

import (
	"github.com/rivo/tview"
)

// DisplayEmails shows a list of emails in a TUI
func DisplayEmails(emails []string) {
	app := tview.NewApplication()
	list := tview.NewList()

	// Add emails to the list
	for _, email := range emails {
		list.AddItem(email, "", 0, nil)
	}

	list.SetBorder(true).SetTitle(" 📬 Email Inbox ")

	// Run the TUI
	if err := app.SetRoot(list, true).Run(); err != nil {
		panic(err)
	}
}
