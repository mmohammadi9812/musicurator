package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var (
	mainContainer *fyne.Container
)

func main() {
	a := app.New()
	w := a.NewWindow("Musicurator")
	w.Resize(fyne.Size{
		Width:  800,
		Height: 600,
	})
	w.SetFixedSize(true)
	srcPath := binding.NewString()
	dstPath := binding.NewString()
	tmplPath := binding.NewString()
	srcEntry := widget.NewEntryWithData(srcPath)
	dstEntry := widget.NewEntryWithData(dstPath)
	tmplEntry := widget.NewEntryWithData(tmplPath)
	mainContainer = container.NewVBox(
		srcEntry,
		dstEntry,
		tmplEntry,
		submit,
	)
	w.SetContent(mainContainer)
	w.ShowAndRun()
}