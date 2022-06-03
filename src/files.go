package src

import (
	"fmt"
	"io"
	"os"
	"path"
)

func MakeDirIfNeeded(dir string) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.MkdirAll(dir, 775)
	}
}

func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		panic(err)
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		panic(err)
		return 0, err
	}
	defer source.Close()

	MakeDirIfNeeded(path.Dir(dst))
	destination, err := os.Create(dst)
	if err != nil {
		panic(err)
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
