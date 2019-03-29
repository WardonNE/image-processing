package main

import (
	"image"
	"math"
)

func GetIntGrayValueWeightedMean(img image.Image, x int, y int) int {
	r, g, b, _ := GetFloat64RGBA(img, x, y)
	return int(r*0.299 + g*0.587 + b*0.114)
}

func GetFloat64GrayValueWeightedMean(img image.Image, x int, y int) float64 {
	r, g, b, _ := GetFloat64RGBA(img, x, y)
	return r*0.299 + g*0.587 + b*0.114
}

func GetUint8GrayValueWeightedMean(img image.Image, x int, y int) uint8 {
	r, g, b, _ := GetFloat64RGBA(img, x, y)
	return uint8(r*0.299 + g*0.587 + b*0.114)
}

func GetIntGrayValueAvg(img image.Image, x int, y int) int {
	r, g, b, _ := GetIntRGBA(img, x, y)
	return int((r + g + b) / 3)
}

func GetUint8GrayValueAvg(img image.Image, x int, y int) uint8 {
	r, g, b, _ := GetIntRGBA(img, x, y)
	return uint8((r + g + b) / 3)
}

func GetFloat64GrayValueAvg(img image.Image, x int, y int) float64 {
	r, g, b, _ := GetIntRGBA(img, x, y)
	return float64((r + g + b) / 3)
}

func GetIntGrayValueMax(img image.Image, x int, y int) int {
	r, g, b, _ := GetFloat64RGBA(img, x, y)
	return int(math.Max(r, math.Max(g, b)))
}

func GetUint8GrayValueMax(img image.Image, x int, y int) uint8 {
	r, g, b, _ := GetFloat64RGBA(img, x, y)
	return uint8(math.Max(r, math.Max(g, b)))
}

func GetFloat64GrayValueMax(img image.Image, x int, y int) float64 {
	r, g, b, _ := GetFloat64RGBA(img, x, y)
	return math.Max(r, math.Max(g, b))
}

func GetIntGrayValuePs(img image.Image, x int, y int) int {
	r, g, b, _ := GetFloat64RGBA(img, x, y)
	return int(math.Pow(math.Pow(r, 2.2)+math.Pow(g, 2.2)*0.6427+math.Pow(b, 2.2)*0.0753, 1.0/2.2))
}

func GetUint8GrayValuePs(img image.Image, x int, y int) uint8 {
	r, g, b, _ := GetFloat64RGBA(img, x, y)
	return uint8(math.Pow(math.Pow(r, 2.2)+math.Pow(g, 2.2)*0.6427+math.Pow(b, 2.2)*0.0753, 1.0/2.2))
}

func GetFloat64GrayValuePs(img image.Image, x int, y int) float64 {
	r, g, b, _ := GetFloat64RGBA(img, x, y)
	return math.Pow(math.Pow(r, 2.2)+math.Pow(g, 2.2)*0.6427+math.Pow(b, 2.2)*0.0753, 1.0/2.2)
}

func GetIntGrayValueRChannel(img image.Image, x int, y int) int {
	r, _, _, _ := GetIntRGBA(img, x, y)
	return r
}

func GetUint8GrayValueRChannel(img image.Image, x int, y int) uint8 {
	r, _, _, _ := GetUint8RGBA(img, x, y)
	return r
}

func GetFloat64GrayValueRChannel(img image.Image, x int, y int) float64 {
	r, _, _, _ := GetFloat64RGBA(img, x, y)
	return r
}

func GetIntGrayValueGChannel(img image.Image, x int, y int) int {
	_, g, _, _ := GetIntRGBA(img, x, y)
	return g
}

func GetUint8GrayValueGChannel(img image.Image, x int, y int) uint8 {
	_, g, _, _ := GetUint8RGBA(img, x, y)
	return g
}

func GetFloat64GrayValueGChannel(img image.Image, x int, y int) float64 {
	_, g, _, _ := GetFloat64RGBA(img, x, y)
	return g
}

func GetIntGrayValueBChannel(img image.Image, x int, y int) int {
	_, _, b, _ := GetIntRGBA(img, x, y)
	return b
}

func GetUint8GrayValueBChannel(img image.Image, x int, y int) uint8 {
	_, _, b, _ := GetUint8RGBA(img, x, y)
	return b
}

func GetFloat64GrayValueBChannel(img image.Image, x int, y int) float64 {
	_, _, b, _ := GetFloat64RGBA(img, x, y)
	return b
}

func GrayLevelStatWeightedMean(img image.Image) [256]int {
	dx, dy := GetImageBounds(img)
	var sgray [256]int
	for k := 0; k < 256; k++ {
		sgray[k] = 0
	}
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			gary := GetIntGrayValueWeightedMean(img, i, j)
			sgray[gray]++
		}
	}
	return sgray
}

func GrayLevelStatAvg(img image.Image) [256]int {
	dx, dy := GetImageBounds(img)
	var sgray [256]int
	for k := 0; k < 256; k++ {
		sgray[k] = 0
	}
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			gary := GetIntGrayValueAvg(img, i, j)
			sgray[gray]++
		}
	}
	return sgary
}

func GrayLevelStatMax(img image.Image) [256]int {
	dx, dy := GetImageBounds(img)
	var sgray [256]int
	for k := 0; k < 256; k++ {
		sgray[k] = 0
	}
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			gary := GetIntGrayValueMax(img, i, j)
			sgray[gray]++
		}
	}
	return sgary
}

func GrayLevelStatRChannel(img image.Image) [256]int {
	dx, dy := GetImageBounds(img)
	var sgray [256]int
	for k := 0; k < 256; k++ {
		sgray[k] = 0
	}
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			gary := GetIntGrayValueRChannel(img, i, j)
			sgray[gray]++
		}
	}
	return sgary
}

func GrayLevelStatGChannel(img image.Image) [256]int {
	dx, dy := GetImageBounds(img)
	var sgray [256]int
	for k := 0; k < 256; k++ {
		sgray[k] = 0
	}
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			gary := GetIntGrayValueGChannel(img, i, j)
			sgray[gray]++
		}
	}
	return sgary
}
