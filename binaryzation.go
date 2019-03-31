package main

import (
	"image"
	"image/color"
	_ "math"
)

var (
	black = color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)}
	white = color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)}
)

const (
	ASH_AVG           int = 1
	ASH_MAX           int = 2
	ASH_PS            int = 3
	ASH_WEIGHTED_MEAN int = 4
	ASH_RED_CHANNEL   int = 5
	ASH_GREEN_CHANNEL int = 6
	ASH_BLUE_CHANNEL  int = 7
)

func ImageBinary(img image.Image, threshosd int, mode int) image.RGBA {
	width, height := GetImageBounds(img)
	canvas := CreateCanvas(width, height)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			var gray int
			switch mode {
			case 1:
				gray = GetIntGrayValueAvg(img, i, j)
			case 2:
				gray = GetIntGrayValueMax(img, i, j)
			case 3:
				gray = GetIntGrayValuePs(img, i, j)
			case 4:
				gray = GetIntGrayValueWeightedMean(img, i, j)
			case 5:
				gray = GetIntGrayValueRChannel(img, i, j)
			case 6:
				gray = GetIntGrayValueGChannel(img, i, j)
			case 7:
				gray = GetIntGrayValueBChannel(img, i, j)
			}
			if gray >= threshosd {
				canvas.SetRGBA(i, j, black)
			} else {
				canvas.SetRGBA(i, j, white)
			}
		}
	}
	return canvas
}
