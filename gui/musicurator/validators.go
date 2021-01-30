package main

import (
	"errors"
	"os"
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
	tmplReg := regexp.MustCompile("([a-zA-Z0-9 -_]*)(\\$artist|\\$album|\\$title|\\$ext|[ _-])+([a-zA-Z0-9 -_]*)\\.\\$ext")
	ok := tmplReg.MatchString(tmpl)
	if !ok {
		return errors.New("template didn't match")
	}
	return nil
}