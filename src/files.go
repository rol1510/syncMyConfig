package src

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func MakeDirIfNeeded(dir string) {
	dir = SanitizePath(dir)

	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.MkdirAll(dir, 775)
	}
}

func Copy(src, dst string) (int64, error) {
	src = SanitizePath(src)
	dst = SanitizePath(dst)

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	MakeDirIfNeeded(path.Dir(dst))
	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func SanitizePath(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}
