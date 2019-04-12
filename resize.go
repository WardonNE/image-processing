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
	for i := 0; i <= w; i++ {
		for j := 0; j <= h; j++ {
			var R, G, B, A uint8
			switch t {
			case NEAREST:
				R, G, B, A = nearest(img, i, j, r, r)
			case BILINEAR:
				R, G, B, A = bilinear(img, i, j, r, r)
			case BICUBIC:
				R, G, B, A = bicubic(img, i, j, r, r)
			}
			canvas.SetRGBA(i, j, color.RGBA{R, G, B, A})
		}
	}
	return canvas
}

func Resize(img image.Image, x, y, t int) image.RGBA {
	var dx, dy int = GetImageBounds(img)
	var canvas = CreateCanvas(x, y)
	var xscale float64 = float64(x) / float64(dx)
	var yscale float64 = float64(y) / float64(dy)
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			var R, G, B, A uint8
			switch t {
			case NEAREST:
				R, G, B, A = nearest(img, i, j, xscale, yscale)
			case BILINEAR:
				R, G, B, A = bilinear(img, i, j, xscale, yscale)
			case BICUBIC:
				R, G, B, A = bicubic(img, i, j, xscale, yscale)
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

func getW_xy(point_x, point_y [4]int, x0, y0 float64) [4][4]float64 {
	var k [4][4]float64
	w_x := getW_x(point_x, x0)
	w_y := getW_y(point_y, y0)
	for i, val := range w_x {
		for j, v := range w_y {
			k[i][j] = v * val
		}
	}
	return k
}

func bicubic(img image.Image, x, y int, xscale, yscale float64) (uint8, uint8, uint8, uint8) {
	var w, h = GetImageBounds(img)
	var x0, y0 float64 = float64(x) * (1.0 / xscale), float64(y) * (1.0 / yscale)
	var point_x, point_y [4]int
	for i := 0; i <= 3; i++ {
		point_x[i] = int(x0) + i - 1
	}
	for i := 0; i <= 3; i++ {
		point_y[i] = int(y0) + i - 1
	}
	var k = getW_xy(point_x, point_y, x0, y0)
	var r, g, b float64
	for i, val := range k {
		if point_x[i] < 0 {
			point_x[i] = -point_x[i]
		} else if point_x[i] > w {
			point_x[i] = w - (point_x[i] - w)
		}
		for j, v := range val {

			if point_y[j] < 0 {
				point_y[j] = -point_y[j]
			} else if point_y[j] > h {
				point_y[j] = h - (point_y[j] - h)
			}
			R, G, B, _ := GetFloat64RGBA(img, point_x[i], point_y[j])
			r += R * v
			g += G * v
			b += B * v
		}
	}
	return uint8(r), uint8(g), uint8(b), uint8(MAX_CHANNEL_VALUE)
}

func getW_x(point_x [4]int, x0 float64) [4]float64 {
	var a = -0.5
	X := float64(int(x0))
	var step_x, w_x [4]float64
	step_x[0] = 1.0 + (x0 - X)
	step_x[1] = x0 - X
	step_x[2] = 1.0 - (x0 - X)
	step_x[3] = 2.0 - (x0 - X)

	w_x[0] = a*math.Abs(step_x[0]*step_x[0]*step_x[0]) - 5.0*a*step_x[0]*step_x[0] + 8.0*a*math.Abs(step_x[0]) - 4.0*a
	w_x[1] = (a+2.0)*math.Abs(step_x[1]*step_x[1]*step_x[1]) - (a+3.0)*step_x[1]*step_x[1] + 1.0
	w_x[2] = (a+2.0)*math.Abs(step_x[2]*step_x[2]*step_x[2]) - (a+3.0)*step_x[2]*step_x[2] + 1.0
	w_x[3] = a*math.Abs(step_x[3]*step_x[3]*step_x[3]) - 5.0*a*step_x[3]*step_x[3] + 8.0*a*math.Abs(step_x[3]) - 4.0*a
	return w_x
}

func getW_y(point_y [4]int, y0 float64) [4]float64 {
	var a = -0.5
	Y := float64(int(y0))
	var step_y, w_y [4]float64
	step_y[0] = 1.0 + (y0 - Y)
	step_y[1] = y0 - Y
	step_y[2] = 1.0 - (y0 - Y)
	step_y[3] = 2.0 - (y0 - Y)

	w_y[0] = a*math.Abs(step_y[0]*step_y[0]*step_y[0]) - 5.0*a*step_y[0]*step_y[0] + 8.0*a*math.Abs(step_y[0]) - 4.0*a
	w_y[1] = (a+2.0)*math.Abs(step_y[1]*step_y[1]*step_y[1]) - (a+3.0)*step_y[1]*step_y[1] + 1.0
	w_y[2] = (a+2.0)*math.Abs(step_y[2]*step_y[2]*step_y[2]) - (a+3.0)*step_y[2]*step_y[2] + 1.0
	w_y[3] = a*math.Abs(step_y[3]*step_y[3]*step_y[3]) - 5.0*a*step_y[3]*step_y[3] + 8.0*a*math.Abs(step_y[3]) - 4.0*a
	return w_y
}

func mitchellNetravali(x0, y0 float64) (uint8, uint8, uint8, uint8) {
	var B, C float64 = 1.0 / 3.0
	var d float64 = math.Sqrt(math.Pow(x0-round(x0), 2) + math.Pow((y0-int(y0)), 2))
	type float64RGB struct {
		R, G, B, A float64
	}
	var P0, P1, P2, P3 float64RGBA
	P0.R, P0.G, P0.B, P0.A = GetFloat64RGBA(img, int(x0), int(y0))
	P1.R, P1.G, P1.B, P1.A = GetFloat64RGBA(img, round(x0), int(y0))
	R2.R, P2.G, P2.B, P2.A = GetFloat64RGBA(img, int(x0), round(y0))
	R3.R, P3.G, P3.B, P3.A = GetFloat64RGBA(img, round(x0), round(y0))
	var r, g, b float64
	r = ((-1.0/6.0*B-C)*P0.R+(-3.0/2.0*B-C+2.0)*P1.R+(3.0/2.0*B+C-2.0)*P2.R+(1.0/6.0*B+C)*P3.R)*math.Pow(d, 3) + ((1.0/2.0*B+2.0*C)*P0.R+(2.0*B+C-3.0)*P1.R+(-5.0/2.0*B-2.0*C+3.0)*P2.R-C*P3.R)*math.Pow(d, 2) + ((-1.0/2.0*B-C)*P0.R+(1.0/2.0*B+C)*P2.R)*d + 1.0/6.0*B*P0.R + (-1.0/3.0*B+1.0)*P1.R + 1.0/6.0*B*P2.R
	g = ((-1.0/6.0*B-C)*P0.G+(-3.0/2.0*B-C+2.0)*P1.G+(3.0/2.0*B+C-2.0)*P2.G+(1.0/6.0*B+C)*P3.G)*math.Pow(d, 3) + ((1.0/2.0*B+2.0*C)*P0.G+(2.0*B+C-3.0)*P1.G+(-5.0/2.0*B-2.0*C+3.0)*P2.G-C*P3.G)*math.Pow(d, 2) + ((-1.0/2.0*B-C)*P0.G+(1.0/2.0*B+C)*P2.G)*d + 1.0/6.0*B*P0.G + (-1.0/3.0*B+1.0)*P1.G + 1.0/6.0*B*P2.G
	b = ((-1.0/6.0*B-C)*P0.B+(-3.0/2.0*B-C+2.0)*P1.B+(3.0/2.0*B+C-2.0)*P2.B+(1.0/6.0*B+C)*P3.B)*math.Pow(d, 3) + ((1.0/2.0*B+2.0*C)*P0.B+(2.0*B+C-3.0)*P1.B+(-5.0/2.0*B-2.0*C+3.0)*P2.B-C*P3.B)*math.Pow(d, 2) + ((-1.0/2.0*B-C)*P0.B+(1.0/2.0*B+C)*P2.B)*d + 1.0/6.0*B*P0.B + (-1.0/3.0*B+1.0)*P1.B + 1.0/6.0*B*P2.B
	return uint8(r), uint8(g), uint8(b), MAX_CHANNEL_VALUE
}
