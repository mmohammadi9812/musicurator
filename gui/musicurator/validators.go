package main

import (
	"errors"
	"os"
	"regexp"
)

func pathValidator(path string) (err error) {
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return err
	}
	return nil
}

func tmplValidator(tmpl string) error {
	tmplReg := regexp.MustCompile("([a-zA-Z0-9 -_]*)(\\$artist|\\$album|\\$title|\\$ext)+([a-zA-Z0-9 -_]*)\\.\\$ext")
	ok := tmplReg.MatchString(tmpl)
	if !ok {
		return errors.New("template didn't match")
	}
	return nil
}