package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"git.sr.ht/~mmohammadi9812/musicaurator/core"
)

var (
	chosedFiles = make(map[core.Music]binding.Bool)
	checkBoxes []*widget.Check
)

func createCheckBox(file core.Music) *widget.Check {
	oldName := file.Name
	newName, err := core.NewName(file, tmpl)
	if err != nil {
		dialog.ShowError(err, w)
	}
	checkBox := widget.NewCheck(fmt.Sprintf("%s -> %s", oldName, newName), func(changed bool) {})
	checkBox.Bind(chosedFiles[file])
	return checkBox
}

func selectAllCheckBox() (out *widget.Check) {
	out = widget.NewCheck("select all", func(b bool) {
		for _, check := range checkBoxes {
			check.Checked = b
		}
	})
	out.Checked = true
	return
}

func createCheckBoxWidget(path string) *fyne.Container {
	files, err := core.Search(path)
	if err != nil {
		dialog.ShowError(err, w)
	}
	selectAll := selectAllCheckBox()
	checkBoxWidget := container.NewVBox(selectAll)
	for _, file := range files {
		checkBox := createCheckBox(file)
		checkBoxes = append(checkBoxes, checkBox)
		chosedFiles[file] = binding.NewBool()
		checkBoxWidget.Add(checkBoxWidget)
	}
	return checkBoxWidget
}
