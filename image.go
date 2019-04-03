package main

import (
	"bytes"
	// "fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
)

func NewImage(path string) (image.Image, error) {
	if f, err := ioutil.ReadFile(path); err == nil {
		b := bytes.NewBuffer(f)
		if pic, _, err := image.Decode(b); err == nil {
			return pic, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func SaveImage(filename string, canvas image.RGBA) error {
	var dirname string = filepath.Dir(filename)
	if !isExist(dirname) {
		err := os.MkdirAll(dirname, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if f, err := os.Create(filename); err == nil {
		defer f.Close()
		err := png.Encode(f, &canvas)
		return err
	} else {
		return err
	}
}

func CreateCanvas(w int, h int) image.RGBA {
	return *image.NewRGBA(image.Rect(0, 0, w, h))
}

func isExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		if os.IsExist(err) {
			return true
		} else {
			return false
		}
	}
}
