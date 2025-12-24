package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/nepalsaurav/personal-finance/ui"
)

func main() {
	app := app.New()

	imgByte, err := os.ReadFile("assets/logo.png")
	if err != nil {
		log.Fatal(err)
	}

	iconRes := fyne.NewStaticResource("app-icon", imgByte)

	app.SetIcon(iconRes)

	window := app.NewWindow("Hledger GUI")
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
