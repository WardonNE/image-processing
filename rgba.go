package main

import (
	"image"
)

// GetIntRGBA ...
func GetIntRGBA(img image.Image, x int, y int) (int, int, int, int) {
	r, g, b, a := img.At(x, y).RGBA()
	return int(r >> 8), int(g >> 8), int(b >> 8), int(a >> 8)
}

// GetUint8RGBA ...
func GetUint8RGBA(img image.Image, x int, y int) (uint8, uint8, uint8, uint8) {
	r, g, b, a := img.At(x, y).RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)
}

// GetFloat64RGBA ...
func GetFloat64RGBA(img image.Image, x int, y int) (float64, float64, float64, float64) {
	r, g, b, a := img.At(x, y).RGBA()
	return float64(r >> 8), float64(g >> 8), float64(b >> 8), float64(a >> 8)
}
