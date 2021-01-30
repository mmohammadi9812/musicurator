package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
	"log"
)

var (
	srcButton = widget.NewButton("src", func(){
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				log.Fatalf("error while openning %s\n", uri.String())
			}
			srcEntry.SetText(uri.String()[7:])
			srcEntry.Refresh()
		}, w)
	})
	dstButton = widget.NewButton("dst", func(){
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				log.Fatalf("Error occured while trying to choose destination folder")
			}
			dstEntry.SetText(uri.String()[7:])
			dstEntry.Refresh()
		}, w)
	})
	submit = widget.NewButton("search", func(){
		if err := srcEntry.Validate(); err != nil {
			dialog.ShowError(err, w)
		}
		if err := dstEntry.Validate(); err != nil {
			dialog.ShowError(err, w)
		}
	})
)
