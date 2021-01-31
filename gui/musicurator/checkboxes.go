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
	newNames = make(map[core.Music]string)
)

func createCheckBox(file core.Music) *widget.Check {
	oldName := file.Name
	newName, err := core.NewName(file, tmpl)
	if err != nil {
		dialog.ShowError(err, w)
	}
	checkBox := widget.NewCheck(fmt.Sprintf("%s -> %s", oldName, newName), func(changed bool) {

	})
	newNames[file] = newName
	checkBox.SetChecked(true)
	checkBox.Bind(chosedFiles[file])
	return checkBox
}

func selectAllCheckBox() (out *widget.Check) {
	out = widget.NewCheck("select all", func(b bool) {
		for _, check := range chosedFiles {
			_ = check.Set(b)
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
		chosedFiles[file] = binding.NewBool()
		chosedFiles[file].Set(true)
		checkBox := createCheckBox(file)
		checkBoxWidget.Add(checkBox)
	}
	return checkBoxWidget
}
