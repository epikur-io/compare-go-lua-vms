package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func FailOnError(b testing.TB, err error) {
	if err == nil {
		return
	}
	b.Error(err)
}

func ListFiles(dir string) (list []string, err error) {
	var files []os.DirEntry
	dn, _ := filepath.Abs(dir)
	files, err = os.ReadDir(dn)
	if err != nil {
		return
	}
	for _, f := range files {
		if !f.IsDir() {
			//fn, _ := filepath.Abs(f.Name())
			fn := filepath.Join(dn, f.Name())
			list = append(list, fn)
		}
	}
	return
}
