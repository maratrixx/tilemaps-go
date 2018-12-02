package utils

import (
	"os"
)

// 检查目录是否存在，不存在需要创建
func CheckDirAndMkdir(pathName string) (err error) {
	_, r := os.Stat(pathName)

	if r != nil && os.IsNotExist(r) {
		err = os.MkdirAll(pathName, os.ModePerm)
	}
	return
}

// 检查文件是否存在
func CheckFileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}
