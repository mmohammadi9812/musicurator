package core

import (
	id3 "github.com/bogem/id3v2"
)

type Music struct {
	Name string
	Path string
	ext  string
}

func (m *Music)Artist() (artist string, err error) {
	tag, err := id3.Open(m.Path, id3.Options{Parse: true})
	if err != nil {
		return
	}
	defer tag.Close()
	artist = tag.Artist()
	return
}

func (m *Music)Title() (title string, err error) {
	tag, err := id3.Open(m.Path, id3.Options{Parse: true})
	if err != nil {
		return
	}
	defer tag.Close()
	title = tag.Title()
	return
}

func (m *Music)Album() (album string, err error) {
	tag, err := id3.Open(m.Path, id3.Options{Parse: true})
	if err != nil {
		return
	}
	defer tag.Close()
	album = tag.Album()
	return
}