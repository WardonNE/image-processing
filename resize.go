package main

import (
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
			case NEAREST_NEIGHBOR:
				R, G, B, A = nearest(img, i, j, r, r)
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

// func bilinear(x, y int, xscale, yscale float64) (int, int) {
// 	var x0, y0 float64
// 	x0 = float64(x) * (1.0 / xscale)
// 	y0 = float64(y) * (1.0 / yscale)
// 	var vx, ux, vy, uy float64 = x0 - x, 1.0 - vx, y0 - y, 1.0 - vy
// 	var (
// 		lt = image.Point{int(math.Floor(x0)), int(math.Floor(y0))}
// 		lb = image.Point{int(math.Floor(x0)), math.Ceil(y0)}
// 		rt = image.Point{math.Ceil(x0), math.Ceil(y0)}
// 		rb = image.Point{math.Ceil(x0), math.Floor(y0)}
// 	)
// 	var r, g, b, a uint8
// 	ltR, ltG, ltB, ltA := GetFloat64RGBA(img, lt.X, lt.Y)
// 	lbR, lbG, lbB, lbA := GetFloat64RGBA(img, lb.X, lb.Y)
// 	rtR, rtG, rtB, rtA := GetFloat64RGBA(img, rt.X, rt.Y)
// 	rbR, rbG, rbB, rbA := GetFloat64RGBA(img, rt.X, rt.Y)
// }
