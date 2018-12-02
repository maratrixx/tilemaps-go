package utils

import (
	"io/ioutil"
	"path/filepath"
	"sync"
)

var (
	rootPath string
	once     sync.Once
)

func GetRootPath() (path string) {
	once.Do(func() {
		rootPath, _ = ioutil.TempDir("", "")
	})
	return rootPath
}

func GetPublicPath() (path string) {
	path = filepath.Join(GetRootPath(), "public")
	return
}

func GetStoragePath() (path string) {
	path = filepath.Join(GetRootPath(), "storage")
	return
}
