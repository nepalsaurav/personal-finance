package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AccountPosting struct {
	Account      string
	DebitAmount  string
	CreditAmount string
}

func addJournalEntry(app fyne.App) {
	w := app.NewWindow("Add Account Transactions")
	w.Resize(fyne.NewSize(800, 600))

	// create skeleton form
	form := &widget.Form{
		OnCancel: func() {
			w.Close()
		},
		OnSubmit: func() {
			fmt.Println("form submitted")
		},
	}

	date := datePicker(w)
	form.AppendItem(&widget.FormItem{
		Text: "Date", Widget: date, HintText: "Transaction Date",
	})

	description := widget.NewEntry()
	form.AppendItem(&widget.FormItem{
		Text: "Description", Widget: description, HintText: "Transaction description",
	})

	var postings []*AccountPosting
	var postingRows []fyne.CanvasObject

	postingContainer := container.NewVBox()
	form.AppendItem(&widget.FormItem{
		Text: "Journal Entries", Widget: postingContainer,
	})

	addPostingRow := func() {

		accountEntry := accountSelectWidget(func(account string) {})

		accountSection := container.NewVBox(
			widget.NewLabel("Account"),
			accountEntry,
		)

		debitEntry := widget.NewEntry()
		debitEntry.SetText("0")
		debitSection := container.NewVBox(
			widget.NewLabel("Debit"),
			debitEntry,
		)

		creditEntry := widget.NewEntry()
		creditEntry.SetText("0")
		creditSection := container.NewVBox(
			widget.NewLabel("Credit"),
			creditEntry,
		)

		row := container.NewGridWithColumns(3, accountSection, debitSection, creditSection)
		postingContainer.Add(row)
		postingContainer.Refresh()

		postings = append(postings, &AccountPosting{})
		postingRows = append(postingRows, row)

		accountEntry.OnChanged = func(s string) {
			postings[len(postings)-1].Account = s
		}
		debitEntry.OnChanged = func(s string) {
			postings[len(postings)-1].DebitAmount = s
		}
		creditEntry.OnChanged = func(s string) {
			postings[len(postings)-1].CreditAmount = s
		}
	}

	removeLastPosting := func() {
		if len(postingRows) <= 2 {
			return
		}
		lastRow := postingRows[len(postingRows)-1]
		postingContainer.Remove(lastRow)
		postingContainer.Refresh()

		postingRows = postingRows[:len(postingRows)-1]
		postings = postings[:len(postings)-1]
	}

	// Add initial two rows
	addPostingRow()
	addPostingRow()

	controls := container.NewHBox(
		widget.NewButton("Add More", addPostingRow),
		widget.NewButton("Clear Last", removeLastPosting),
	)

	form.AppendItem(&widget.FormItem{
		Widget: controls,
	})

	scrollable := container.NewScroll(form)
	w.SetContent(scrollable)
	w.Show()
}
