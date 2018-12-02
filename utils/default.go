package utils

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"sync"
)

// 错误检查，发现错误立刻抛出
func MustCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// 保存图片url地址到指定文件
func SaveImage(originUrl, target string) (err error) {

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			return
		}
	}()

	resp, _ := http.Get(originUrl)
	defer resp.Body.Close()

	imageData, _, _ := image.Decode(resp.Body)

	out, _ := os.Create(target)
	defer out.Close()

	png.Encode(out, imageData)

	return
}

// 异步保存图片url地址到指定文件
func SaveImageAsync(originUrl, target string, group *sync.WaitGroup) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[ERROR] SaveImageAsync error: %#v", r)
		}
		group.Done()
	}()

	if CheckFileExists(target) {
		return nil
	}

	resp, err := http.Get(originUrl)
	MustCheck(err)
	defer resp.Body.Close()

	imageData, _, err := image.Decode(resp.Body)
	MustCheck(err)

	out, err := os.Create(target)
	defer out.Close()
	MustCheck(err)

	err = png.Encode(out, imageData)
	MustCheck(err)

	return
}
