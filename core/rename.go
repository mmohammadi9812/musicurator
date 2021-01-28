package core

import (
	"regexp"
)

func NewName(file Music, template string) (name string, err error) {
	var (
		artistRegex = regexp.MustCompile("\\$artist")
		albumRegex = regexp.MustCompile("\\$album")
		titleRegex = regexp.MustCompile("\\$title")
		extRegex = regexp.MustCompile("\\$ext")
	)
	artist, err := file.Artist()
	if err != nil {
		return
	}
	album, err := file.Album()
	if err != nil {
		return
	}
	title, err := file.Title()
	if err != nil {
		return
	}
	name = template
	name = artistRegex.ReplaceAllString(name, artist)
	name = albumRegex.ReplaceAllString(name, album)
	name = titleRegex.ReplaceAllString(name, title)
	name = extRegex.ReplaceAllString(name, file.ext)
	return
}
