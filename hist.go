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
	return int(math.Pow(math.Pow(r, 2.2)*0.2973+math.Pow(g, 2.2)*0.6427+math.Pow(b, 2.2)*0.0753, 1.0/2.2))
}

func GetUint8GrayValuePs(img image.Image, x int, y int) uint8 {
	r, g, b, _ := GetFloat64RGBA(img, x, y)
	return uint8(math.Pow(math.Pow(r, 2.2)*0.2973+math.Pow(g, 2.2)*0.6427+math.Pow(b, 2.2)*0.0753, 1.0/2.2))
}

func GetFloat64GrayValuePs(img image.Image, x int, y int) float64 {
	r, g, b, _ := GetFloat64RGBA(img, x, y)
	return math.Pow(math.Pow(r, 2.2)*0.2973+math.Pow(g, 2.2)*0.6427+math.Pow(b, 2.2)*0.0753, 1.0/2.2)
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

func GrayLevelHistWeightedMean(img image.Image) [256]int {
	dx, dy := GetImageBounds(img)
	var sgray [256]int
	for k := 0; k < 256; k++ {
		sgray[k] = 0
	}
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			gray := GetIntGrayValueWeightedMean(img, i, j)
			sgray[gray]++
		}
	}
	return sgray
}

func GrayLevelHistAvg(img image.Image) [256]int {
	dx, dy := GetImageBounds(img)
	var sgray [256]int
	for k := 0; k < 256; k++ {
		sgray[k] = 0
	}
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			gray := GetIntGrayValueAvg(img, i, j)
			sgray[gray]++
		}
	}
	return sgray
}

func GrayLevelHistMax(img image.Image) [256]int {
	dx, dy := GetImageBounds(img)
	var sgray [256]int
	for k := 0; k < 256; k++ {
		sgray[k] = 0
	}
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			gray := GetIntGrayValueMax(img, i, j)
			sgray[gray]++
		}
	}
	return sgray
}

func GrayLevelHistRChannel(img image.Image) [256]int {
	dx, dy := GetImageBounds(img)
	var sgray [256]int
	for k := 0; k < 256; k++ {
		sgray[k] = 0
	}
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			gray := GetIntGrayValueRChannel(img, i, j)
			sgray[gray]++
		}
	}
	return sgray
}

func GrayLevelHistGChannel(img image.Image) [256]int {
	dx, dy := GetImageBounds(img)
	var sgray [256]int
	for k := 0; k < 256; k++ {
		sgray[k] = 0
	}
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			gray := GetIntGrayValueGChannel(img, i, j)
			sgray[gray]++
		}
	}
	return sgray
}

func GetGrayMat(img image.Image, mode int) [][]int {
	var dx, dy = GetImageBounds(img)
	var mat [][]int
	for i := 0; i < dx; i++ {
		y_mat := make([]int, 0, dy)
		for i := 0; i < dy; i++ {
			y_mat = append(y_mat, 0)
		}
		mat = append(mat, y_mat)
	}
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
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
			mat[i][j] = gray
		}
	}
	return mat
}
