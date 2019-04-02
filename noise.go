package main

import (
	_ "fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
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

func randFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
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

func boxmullerGaussNoise(mu float64, sigma float64) float64 {
	u := randFloat64(0.0, 1.0)
	v := randFloat64(0.0, 1.0)
	z := math.Sqrt(-2.0*math.Log10(u)) * math.Cos(2.0*math.Pi*v)
	return mu + z*sigma
}

func AddGaussNoise(img image.Image, mu float64, sigma float64, k int) image.RGBA {
	var dx, dy int = GetImageBounds(img)
	var canvas image.RGBA = CreateCanvas(dx, dy)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			r, g, b, _ := GetUint8RGBA(img, i, j)
			var temp color.RGBA
			var tempR, tempG, tempB uint8
			tempR = r + uint8(float64(k)*boxmullerGaussNoise(mu, sigma))
			if tempR > MAX_CHANNEL_VALUE {
				tempR = MAX_CHANNEL_VALUE
			} else if tempR < MIN_CHANNEL_VALUE {
				tempR = MIN_CHANNEL_VALUE
			}
			tempG = g + uint8(float64(k)*boxmullerGaussNoise(mu, sigma))
			if tempG > MAX_CHANNEL_VALUE {
				tempG = MAX_CHANNEL_VALUE
			} else if tempG < MIN_CHANNEL_VALUE {
				tempG = MIN_CHANNEL_VALUE
			}
			tempB = b + uint8(float64(k)*boxmullerGaussNoise(mu, sigma))
			if tempB > MAX_CHANNEL_VALUE {
				tempB = MAX_CHANNEL_VALUE
			} else if tempB < MIN_CHANNEL_VALUE {
				tempB = MIN_CHANNEL_VALUE
			}
			temp = color.RGBA{tempR, tempG, tempB, uint8(255)}
			canvas.SetRGBA(i, j, temp)
		}
	}
	return canvas
}

func makePoissonValue(pixel uint8, factor int, mean int) uint8 {
	var l float64 = math.Exp(float64(-factor * mean))
	var k float64 = 0
	var p float64 = 1.0
	for p >= l {
		k += 1.0
		p *= randFloat64(0.0, 1.0)
	}
	return uint8(math.Max(0.0, float64(pixel)+(k-1.0)/float64(mean)-float64(factor)))
}

func AddPoissonNoise(img image.Image, factor int, mean int) image.RGBA {
	var dx, dy = GetImageBounds(img)
	var canvas image.RGBA = CreateCanvas(dx, dy)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			r, g, b, _ := GetUint8RGBA(img, i, j)
			tempR := makePoissonValue(r, factor, mean)
			tempG := makePoissonValue(g, factor, mean)
			tempB := makePoissonValue(b, factor, mean)
			canvas.SetRGBA(i, j, color.RGBA{tempR, tempG, tempB, MAX_CHANNEL_VALUE})
		}
	}
	return canvas
}
