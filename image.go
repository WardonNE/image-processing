package main

import (
	"bytes"
	// "fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
)

func NewImage(path string) (image.Image, error) {
	if f, err := ioutil.ReadFile(path); err == nil {
		b := bytes.NewBuffer(f)
		if pic, err := jpeg.Decode(b); err == nil {
			return pic, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func SaveImage(filename string, canvas image.RGBA) error {
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
