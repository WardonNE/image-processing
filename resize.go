package main

import (
	// "fmt"
	"image"
	"image/color"
	"math"
)

func round(x float64) int {
	return int(math.Floor(x + 0.5))
}

func ScaleImage(img image.Image, r float64, t int) image.RGBA {
	var dx, dy int = GetImageBounds(img)
	var w, h int
	w = int(float64(dx) * r)
	h = int(float64(dy) * r)
	var canvas = CreateCanvas(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			var R, G, B, A uint8
			switch t {
			case NEAREST:
				R, G, B, A = nearest(img, i, j, r, r)
			case BILINEAR:
				R, G, B, A = bilinear(img, i, j, r, r)
			}
			canvas.SetRGBA(i, j, color.RGBA{R, G, B, A})
		}
	}
	return canvas
}

func nearest(img image.Image, x, y int, xscale, yscale float64) (uint8, uint8, uint8, uint8) {
	x0 := int(float64(x) * (1.0 / xscale))
	y0 := int(float64(y) * (1.0 / yscale))
	r, g, b, a := GetUint8RGBA(img, x0, y0)
	return r, g, b, a
}

func bilinear(img image.Image, x, y int, xscale, yscale float64) (uint8, uint8, uint8, uint8) {
	var x0, y0 float64
	x0 = (float64(x)+0.5)*(1.0/xscale) - 0.5
	y0 = (float64(y)+0.5)*(1.0/yscale) - 0.5
	ltR, ltG, ltB, _ := GetFloat64RGBA(img, int(math.Floor(x0)), int(math.Floor(y0)))
	lbR, lbG, lbB, _ := GetFloat64RGBA(img, int(math.Floor(x0)), int(math.Ceil(y0)))
	rtR, rtG, rtB, _ := GetFloat64RGBA(img, int(math.Ceil(x0)), int(math.Floor(y0)))
	rbR, rbG, rbB, _ := GetFloat64RGBA(img, int(math.Ceil(x0)), int(math.Ceil(y0)))
	var v, u float64
	v = y0 - float64(math.Floor(y0))
	u = x0 - float64(math.Floor(x0))
	var res color.RGBA
	res.R = uint8((1.0-u)*(1.0-v)*ltR + (1.0-u)*v*lbR + u*(1-v)*rtR + u*v*rbR)
	res.G = uint8((1.0-u)*(1.0-v)*ltG + (1.0-u)*v*lbG + u*(1-v)*rtG + u*v*rbG)
	res.B = uint8((1.0-u)*(1.0-v)*ltB + (1.0-u)*v*lbB + u*(1-v)*rtB + u*v*rbB)
	res.A = uint8(255)
	return res.R, res.G, res.B, res.A
}
