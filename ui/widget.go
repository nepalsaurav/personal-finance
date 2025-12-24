package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	communityWidget "fyne.io/x/fyne/widget"

	datepicker "github.com/sdassow/fyne-datepicker"
)

func datePicker(window fyne.Window) *widget.Entry {
	date := widget.NewEntry()
	date.SetPlaceHolder("YYYY-MM-DD")
	date.ActionItem = widget.NewButtonWithIcon("", theme.MoreHorizontalIcon(), func() {
		when := time.Now()
		if date.Text != "" {
			t, err := time.Parse("2006/01/02", date.Text)
			if err == nil {
				when = t
			}
		}
		datepicker := datepicker.NewDatePicker(when, time.Monday, func(when time.Time, ok bool) {
			if ok {
				date.SetText(when.Format("2006-01-02"))
			}
		})
		dialog.ShowCustomConfirm(
			"Choose date",
			"Ok",
			"Cancel",
			datepicker,
			datepicker.OnActioned,
			window,
		)
	})
	return date
}

func accountSelectWidget(
	onSelect func(account string),
) *communityWidget.CompletionEntry {

	entry := communityWidget.NewCompletionEntry(nil)

	dummyAccounts := []string{
		"Assets:Cash",
		"Assets:Bank",
		"Expenses:Food",
		"Expenses:Rent",
		"Income:Salary",
	}

	entry.SetOptions(dummyAccounts)
	entry.SetPlaceHolder("Account")

	entry.OnChanged = func(s string) {
		if len(s) < 2 {
			entry.HideCompletion()
			return
		}
		entry.ShowCompletion()
	}

	entry.OnSubmitted = func(s string) {
		onSelect(s)
		entry.HideCompletion()
	}

	return entry
}
