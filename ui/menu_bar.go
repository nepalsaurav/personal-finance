package ui

import (
	"fyne.io/fyne/v2"
)

func CreateMenuBar(window fyne.Window, app fyne.App) {

	// file menu

	quitItem := fyne.NewMenuItem("Quit", func() {
		app.Quit()
	})

	fileMenu := fyne.NewMenu(
		"File",
		quitItem,
	)

	// accounts menu
	accountMenu := fyne.NewMenu("Accounts",
		fyne.NewMenuItem("Add Account", func() {
			addAccount(app)
		}),
		fyne.NewMenuItem("Chart of Accounts", func() {
		}),
	)

	// dashboard
	dashboardItem := fyne.NewMenuItem(
		"Dashboard", func() {},
	)
	DashboardMenu := fyne.NewMenu(
		"Dashboard",
		dashboardItem,
	)

	// transaction
	transactionMenu := fyne.NewMenu(
		"Transactions",
		fyne.NewMenuItem(
			"Transactions List",
			func() {},
		),

		fyne.NewMenuItem(
			"Add Journal Entry",
			func() {
				addJournalEntry(app)
			},
		),
	)

	menu := fyne.NewMainMenu(
		fileMenu,
		DashboardMenu,
		accountMenu,
		transactionMenu,
	)

	window.SetMainMenu(menu)
}
