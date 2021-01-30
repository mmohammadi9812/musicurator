package core

import (
	"os"
	"path/filepath"
)

func isMusicFile(path string) bool {
	extension := filepath.Ext(path)
	switch extension {
		case
			".mp3",
			".m4a",
			".flac",
			".aac",
			".opus":
				return true
	}
	return false
}

// search `root` for music files (based on extension)
func Search(root string) (out []Music,err error) {
	_, err = os.Stat(root)
	if os.IsNotExist(err) {
		return nil, err
	}
	err = filepath.Walk(root, func(path string, info os.FileInfo, perr error) error {
		if perr != nil {
			return perr
		}
		if info.IsDir() {
			return nil
		}
		if isMusicFile(path) {
			out = append(out, Music{
				Name: filepath.Base(path),
				Path: path,
				ext:  filepath.Ext(path),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}
