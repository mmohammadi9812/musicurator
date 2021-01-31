package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	mainContainer *fyne.Container
	srcEntry *widget.Entry
	dstEntry *widget.Entry
	tmplEntry *widget.Entry
	w fyne.Window
	src, dst, tmpl string
)

func main() {
	a := app.New()
	w = a.NewWindow("Musicurator")
	w.Resize(fyne.Size{
		Width:  950,
		Height: 650,
	})
	w.SetFixedSize(true)
	srcPath := binding.BindString(&src)
	dstPath := binding.BindString(&dst)
	tmplString := binding.BindString(&tmpl)
	srcEntry = widget.NewEntryWithData(srcPath)
	dstEntry = widget.NewEntryWithData(dstPath)
	tmplEntry = widget.NewEntryWithData(tmplString)

	srcEntry.Validator = pathValidator
	tmplEntry.Validator = tmplValidator

	tmplEntry.OnChanged = func(s string){
		tmplEntry.SetText(removeExtTempl(tmplEntry.Text) + ".$ext")
		err := tmplString.Set(removeExtTempl(tmplEntry.Text) + ".$ext")
		if err != nil {
			dialog.ShowError(err, w)
		}
	}
	tmplEntry.Refresh()

	srcContainer := container.New(layout.NewFormLayout(), srcButton, srcEntry)
	dstContainer := container.New(layout.NewFormLayout(), dstButton, dstEntry)
	tmplContainer := container.NewVBox(
		container.NewHBox(titleButton, artistButton, albumButton, dashButton, underscoreButton, spaceButton, removeButton),
		tmplEntry,
	)

	submitWidget = container.NewHBox()
	submitWidget.Add(searchButton)

	mainContainer = container.NewVBox(
		srcContainer,
		dstContainer,
		tmplContainer,
		submitWidget,
	)
	w.SetContent(mainContainer)
	w.ShowAndRun()
}