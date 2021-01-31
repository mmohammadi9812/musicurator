package main

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2/dialog"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func removeExtTempl(s string) string {
	if strings.HasSuffix(s, ".$ext") {
		return s[:len(s) - len(".$ext")]
	} else {
		return s
	}
}

func pathValidator(path string) (err error) {
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return err
	}
	return nil
}

func tmplValidator(tmpl string) error {
	tmplReg := regexp.MustCompile("([a-zA-Z0-9 -_]*)(\\$artist|\\$album|\\$title|\\$name|\\$ext|[ _-])+([a-zA-Z0-9 -_]*)\\.\\$ext")
	ok := tmplReg.MatchString(tmpl)
	if !ok {
		return errors.New("template didn't match")
	}
	return nil
}


func submitFunc(){
	_, err := os.Stat(dst)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dst, 0774)
		if err != nil {
			dialog.ShowError(err, w)
		}
	}
	numFiles := 0
	for file, isChosen := range chosedFiles {
		if ok ,err := isChosen.Get(); err != nil || !ok {
			continue
		}
		newPath := path.Join(dst, newNames[file])
		err = os.Rename(file.Path, newPath)
		if err != nil {
			continue
		}
		numFiles += 1
	}

	if len(mainContainer.Objects) > 4 {
		mainContainer.Remove(mainContainer.Objects[len(mainContainer.Objects) - 1])
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}
	relSrc, err := filepath.Rel(home, src)
	if err != nil {
		return
	}
	relDst, err := filepath.Rel(home, dst)
	if err != nil {
		return
	}
	dialog.ShowInformation("Done",
		fmt.Sprintf("Successfully renamed %d files\nmoved from %s to %s", numFiles, relSrc, relDst),
		w)
	srcEntry.SetText("")
	dstEntry.SetText("")
	tmplEntry.SetText("$name")
}