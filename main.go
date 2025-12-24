package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/nepalsaurav/personal-finance/ui"
)

func main() {
	app := app.New()

	app.Settings().SetTheme(&customTheme{Theme: theme.DefaultTheme()})

	window := app.NewWindow("Personal Finance")
	window.Resize(fyne.NewSize(1000, 800))

	// create menu bar
	ui.CreateMenuBar(window, app)

	mainContent := widget.NewLabel("hello world")

	appLayout := container.New(
		layout.NewCustomPaddedLayout(10, 0, 0, 0),
		mainContent,
	)

	window.SetContent(appLayout)

	window.ShowAndRun()
}
