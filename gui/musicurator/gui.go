package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	mainContainer *fyne.Container
	srcEntry *widget.Entry
	dstEntry *widget.Entry
	w fyne.Window
)

func main() {
	a := app.New()
	w = a.NewWindow("Musicurator")
	w.Resize(fyne.Size{
		Width:  800,
		Height: 600,
	})
	w.SetFixedSize(true)
	srcPath := binding.NewString()
	dstPath := binding.NewString()
	tmplString := binding.NewString()
	srcEntry = widget.NewEntryWithData(srcPath)
	dstEntry = widget.NewEntryWithData(dstPath)
	tmplEntry := widget.NewEntryWithData(tmplString)

	srcEntry.Validator = pathValidator
	dstEntry.Validator = pathValidator
	tmplEntry.Validator = tmplValidator

	srcContainer := container.New(layout.NewFormLayout(), srcButton, srcEntry)
	dstContainer := container.New(layout.NewFormLayout(), dstButton, dstEntry)

	mainContainer = container.NewVBox(
		srcContainer,
		dstContainer,
		tmplEntry,
		submit,
	)
	w.SetContent(mainContainer)
	w.ShowAndRun()
}