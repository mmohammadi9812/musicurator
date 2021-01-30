package main

import (
	"fmt"
	"git.sr.ht/~mmohammadi9812/musicaurator/core"
	"github.com/enescakir/emoji"
	flag "github.com/spf13/pflag"
	"log"
	"os"
	"path"
)

var (
	src, dst string
	template string
	copyIt, dryRun, help    bool
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	flag.StringVarP(&src, "source", "s", path.Join(home, "Music"),
		"It's the source from which the program will start searching for music files")
	flag.StringVarP(&dst, "destination", "d", path.Join(home, "Music"),
		"It's where the music files will be move to (renamed to) after renaming files")
	flag.StringVarP(&template, "template", "t", "$artist__$title.$ext",
		"This template is used for renaming files, valid variables are `$artist`, `$title`, `$album` and `$ext`")
	flag.BoolVarP(&copyIt, "copy", "c", false, "Whether to copy renamed files or move them")
	flag.BoolVarP(&dryRun, "dryrun", "n", false, "just show what it does, don't do them")
	flag.BoolVarP(&help, "help", "h", false, "show this help")
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	files, err := core.Search(src)
	if err != nil {
		log.Fatal("Something happened while searching for music files: ", err)
	}
	for _, file := range files {
		name, err := core.NewName(file, template)
		if err != nil {
			log.Fatal("Something happened while trying to rename file based on template:", err)
		}
		if dryRun {
			fmt.Printf("%v %s -> %s\n", emoji.MusicalScore, file.Name, name)
			continue
		}
		dstPath := path.Join(dst, name)
		_, err = os.Stat(dstPath)
		if os.IsExist(err) {
			log.Fatalf("%s already exists\n", dstPath)
		}
		_, err = os.Stat(dst)
		if os.IsNotExist(err) {
			err = os.MkdirAll(dst, 0774)
			if err != nil {
				log.Fatalf("Error while creating %s:\n\t%v\n", dst, err)
			}
		}
		err = os.Rename(file.Path, dstPath)
		if err != nil {
			log.Fatalf("Error while moving %s to %s\n", file.Name, name)
		}
	}
}
