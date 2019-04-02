package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"
)

const (
	MAX_CHANNEL_VALUE uint8 = 255
	MIN_CHANNEL_VALUE uint8 = 0
)

var SaltAndPeper [6]color.RGBA = [6]color.RGBA{
	{MIN_CHANNEL_VALUE, MIN_CHANNEL_VALUE, MIN_CHANNEL_VALUE, MAX_CHANNEL_VALUE},
	{MAX_CHANNEL_VALUE, MIN_CHANNEL_VALUE, MIN_CHANNEL_VALUE, MAX_CHANNEL_VALUE},
	{MAX_CHANNEL_VALUE, MAX_CHANNEL_VALUE, MIN_CHANNEL_VALUE, MAX_CHANNEL_VALUE},
	{MIN_CHANNEL_VALUE, MIN_CHANNEL_VALUE, MAX_CHANNEL_VALUE, MAX_CHANNEL_VALUE},
	{MIN_CHANNEL_VALUE, MAX_CHANNEL_VALUE, MAX_CHANNEL_VALUE, MAX_CHANNEL_VALUE},
	{MAX_CHANNEL_VALUE, MAX_CHANNEL_VALUE, MAX_CHANNEL_VALUE, MAX_CHANNEL_VALUE},
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func AddPeperNoise(img image.Image, p float64) image.RGBA {
	var dx, dy int = GetImageBounds(img)
	var n int = int(math.Ceil(float64(dx*dy) * p))
	var mat [][]color.RGBA
	for i := 0; i < dx; i++ {
		y_mat := make([]color.RGBA, 0, dy)
		for j := 0; j < dy; j++ {
			r, g, b, a := GetUint8RGBA(img, i, j)
			y_mat = append(y_mat, color.RGBA{r, g, b, a})
		}
		mat = append(mat, y_mat)
	}
	for i := 0; i < n; i++ {
		x := randInt(0, dx)
		y := randInt(0, dy)
		mat[x][y] = white
	}
	canvas := CreateCanvas(dx, dy)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			canvas.SetRGBA(i, j, mat[i][j])
		}
	}
	return canvas
}

func AddSaltNoise(img image.Image, p float64) image.RGBA {
	var dx, dy int = GetImageBounds(img)
	var n int = int(math.Ceil(float64(dx*dy) * p))
	var mat [][]color.RGBA
	for i := 0; i < dx; i++ {
		y_mat := make([]color.RGBA, 0, dy)
		for j := 0; j < dy; j++ {
			r, g, b, a := GetUint8RGBA(img, i, j)
			y_mat = append(y_mat, color.RGBA{r, g, b, a})
		}
		mat = append(mat, y_mat)
	}
	for i := 0; i < n; i++ {
		x := randInt(0, dx)
		y := randInt(0, dy)
		mat[x][y] = black
	}
	canvas := CreateCanvas(dx, dy)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			canvas.SetRGBA(i, j, mat[i][j])
		}
	}
	return canvas
}

func AddSaleAndPeperNoist(img image.Image, p float64) image.RGBA {
	var dx, dy int = GetImageBounds(img)
	var n int = int(math.Ceil(float64(dx*dy) * p))
	var mat [][]color.RGBA
	for i := 0; i < dx; i++ {
		y_mat := make([]color.RGBA, 0, dy)
		for j := 0; j < dy; j++ {
			r, g, b, a := GetUint8RGBA(img, i, j)
			y_mat = append(y_mat, color.RGBA{r, g, b, a})
		}
		mat = append(mat, y_mat)
	}
	for i := 0; i < n; i++ {
		x := randInt(0, dx)
		y := randInt(0, dy)
		key := randInt(0, len(SaltAndPeper)-1)
		mat[x][y] = SaltAndPeper[key]
	}
	canvas := CreateCanvas(dx, dy)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			canvas.SetRGBA(i, j, mat[i][j])
		}
	}
	return canvas
}
