package main

import (
	"image"
	"image/color"
)

/**
 * GRAY = (RED+BLUE+GREEN)/3
 */
func AvgAsh(img image.Image) image.RGBA {
	w, h := GetImageBounds(img)
	canvas := CreateCanvas(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			q := GetUint8GrayValueAvg(img, i, j)
			canvas.SetRGBA(i, j, color.RGBA{q, q, q, uint8(255)})
		}
	}
	return canvas
}

/**
 * Gray = R*0.299 + G*0.587 + B*0.114
 */
func WeightedMeanAsh(img image.Image) image.RGBA {
	w, h := GetImageBounds(img)
	canvas := CreateCanvas(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			q := GetUint8GrayValueWeightedMean(img, i, j)
			canvas.SetRGBA(i, j, color.RGBA{q, q, q, uint8(255)})
		}
	}
	return canvas
}

/**
 * Adobe RGB (1998) [gamma=2.20]
 * Gray = (R^2.2 * 0.2973 + G^2.2 * 0.6274 + B^2.2 * 0.0753)^(1/2.2)
 */
func PsAsh(img image.Image) image.RGBA {
	w, h := GetImageBounds(img)
	canvas := CreateCanvas(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			q := GetUint8GrayValuePs(img, i, j)
			canvas.SetRGBA(i, j, color.RGBA{q, q, q, uint8(255)})
		}
	}
	return canvas
}

func RChannelAsh(img image.Image) image.RGBA {
	w, h := GetImageBounds(img)
	canvas := CreateCanvas(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			q := GetUint8GrayValueRChannel(img, i, j)
			canvas.SetRGBA(i, j, color.RGBA{q, q, q, uint8(255)})
		}
	}
	return canvas
}

func GChannelAsh(img image.Image) image.RGBA {
	w, h := GetImageBounds(img)
	canvas := CreateCanvas(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			q := GetUint8GrayValueGChannel(img, i, j)
			canvas.SetRGBA(i, j, color.RGBA{q, q, q, uint8(255)})
		}
	}
	return canvas
}

func BChannelAsh(img image.Image) image.RGBA {
	w, h := GetImageBounds(img)
	canvas := CreateCanvas(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			q := GetUint8GrayValueBChannel(img, i, j)
			canvas.SetRGBA(i, j, color.RGBA{q, q, q, uint8(255)})
		}
	}
	return canvas
}

/**
 * gray = max(r,g,b)
 */
func MaxAsh(img image.Image) image.RGBA {
	w, h := GetImageBounds(img)
	canvas := CreateCanvas(w, h)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			q := GetUint8GrayValueMax(img, i, j)
			canvas.SetRGBA(i, j, color.RGBA{q, q, q, uint8(255)})
		}
	}
	return canvas
}
