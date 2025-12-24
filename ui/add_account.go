package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func addAccount(app fyne.App) {
	window := app.NewWindow("Add Account")
	window.Resize(fyne.NewSize(600, 400))

	form := &widget.Form{}

	account := widget.NewEntry()
	form.Append("Account", account)
	helpText := canvas.NewText("Account format: Assets:Investments:Stocks,Use : to separate levels, from general to specific.", color.RGBA{R: 150, G: 150, B: 150, A: 255})

	form.Append("", helpText)
	form.Append("", layout.NewSpacer())
	save := widget.NewButton("Save Account", func() {})
	save.Importance = widget.HighImportance
	cancel := widget.NewButton("Cancel", func() {
		window.Close()
	})
	form.Append("", container.NewHBox(save, cancel))

	window.SetContent(
		container.NewScroll(form),
	)
	window.Show()
}
