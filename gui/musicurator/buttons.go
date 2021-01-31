package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
	"regexp"
)

var (
	srcButton = widget.NewButton("src", func(){
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				log.Fatalf("error while openning %s\n", uri.String())
			}
			if uri != nil{
				srcEntry.SetText(uri.String()[7:])
				srcEntry.Refresh()
				dstEntry.SetText(uri.String()[7:])
				dstEntry.Refresh()
			}
		}, w)
	})
	dstButton = widget.NewButton("dst", func(){
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				log.Fatalf("Error occured while trying to choose destination folder")
			}
			if uri != nil {
				dstEntry.SetText(uri.String()[7:])
				dstEntry.Refresh()
			}
		}, w)
	})

)

var (
	artistButton = widget.NewButton("artist name", func(){
		tmplEntry.SetText(removeExtTempl(tmplEntry.Text) + "$artist")
	})
	titleButton = widget.NewButton("title of song", func(){
		tmplEntry.SetText(removeExtTempl(tmplEntry.Text) + "$title")
	})
	albumButton = widget.NewButton("album name", func(){
		tmplEntry.SetText(removeExtTempl(tmplEntry.Text) + "$album")
	})
	dashButton = widget.NewButton("-", func(){
		tmplEntry.SetText(removeExtTempl(tmplEntry.Text) + "-")
	})
	underscoreButton = widget.NewButton("_", func() {
		tmplEntry.SetText(removeExtTempl(tmplEntry.Text) + "_")
	})
	spaceButton = widget.NewButton("SPACE", func() {
		tmplEntry.SetText(removeExtTempl(tmplEntry.Text) + " ")
	})
	removeButton = widget.NewButton("<-", func(){
		lastPartOfTmplRegex := regexp.MustCompile("(\\$artist|\\$album|\\$name|\\$title|[ _-])$")
		newTemplate := lastPartOfTmplRegex.ReplaceAllString(removeExtTempl(tmplEntry.Text), "")
		tmplEntry.SetText(newTemplate + ".$ext")
	})
)

var (
	submitWidget *fyne.Container
	submitButton = widget.NewButton("submit", submitFunc)
	searchButton = widget.NewButton("search", func(){
		if err := srcEntry.Validate(); err != nil {
			dialog.ShowError(err, w)
		}
		if err := tmplEntry.Validate(); err != nil {
			dialog.ShowError(err, w)
		}
		if len(submitWidget.Objects) == 1 {
			submitWidget.Add(submitButton)
		}
		checkboxWidget := createCheckBoxWidget(src)
		if len(mainContainer.Objects) > 4 {
			mainContainer.Remove(mainContainer.Objects[len(mainContainer.Objects) - 1])
		}
		mainContainer.Add(checkboxWidget)
	})
)