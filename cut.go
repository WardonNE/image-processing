package main

import (
	_ "fmt"
	"image"
	"image/color"
	"math"
)

const (
	POSITION_LEFT_TOP      int = 1
	POSITION_CENTER_TOP    int = 2
	POSITION_RIGHT_TOP     int = 3
	POSITION_LEFT_MIDDLE   int = 4
	POSITION_CENTER_MIDDLE int = 5
	POSITION_RIGHT_MIDDLE  int = 6
	POSITION_LEFT_BUTTON   int = 7
	POSITION_CENTER_BUTTON int = 8
	POSITION_RIGHT_BUTTON  int = 9
)

func QuickCut(img image.Image, position int) image.RGBA {
	x, y := GetImageBounds(img)
	w := int(math.Ceil(float64(x / 3)))
	h := int(math.Ceil(float64(y / 3)))
	canvas := CreateCanvas(w, h)
	var x_offset, y_offset int
	switch {
	case position == POSITION_LEFT_TOP:
		x_offset = 0
		y_offset = 0
	case position == POSITION_CENTER_TOP:
		x_offset = w
		y_offset = 0
	case position == POSITION_RIGHT_TOP:
		x_offset = 2 * w
		y_offset = 0
	case position == POSITION_LEFT_MIDDLE:
		x_offset = 0
		y_offset = h
	case position == POSITION_CENTER_MIDDLE:
		x_offset = w
		y_offset = h
	case position == POSITION_RIGHT_MIDDLE:
		x_offset = 2 * w
		y_offset = h
	case position == POSITION_LEFT_BUTTON:
		x_offset = 0
		y_offset = 2 * h
	case position == POSITION_CENTER_BUTTON:
		x_offset = w
		y_offset = 2 * h
	case position == POSITION_RIGHT_BUTTON:
		x_offset = 2 * w
		y_offset = 2 * h
	}
	for i := x_offset; i < x_offset+w; i++ {
		if i > x {
			break
		} else {
			for j := y_offset; j < y_offset+h; j++ {
				if j > y {
					break
				} else {
					r, g, b, a := GetUint8RGBA(img, i, j)
					canvas.SetRGBA(i-x_offset, j-y_offset, color.RGBA{r, g, b, a})
				}
			}
		}
	}
	return canvas
}

func SimpleCut(img image.Image, x int, y int, position int) image.RGBA {
	width, height := GetImageBounds(img)
	half_w := int(math.Ceil(float64(width / 2)))
	half_h := int(math.Ceil(float64(height / 2)))
	half_x := int(math.Ceil(float64(x / 2)))
	half_y := int(math.Ceil(float64(y / 2)))

	var x_offset, y_offset int

	switch {
	case position == POSITION_LEFT_TOP:
		x_offset = 0
		y_offset = 0
		if x > width {
			x = width
		}
		if y > height {
			y = height
		}
	case position == POSITION_CENTER_TOP:
		if x > width {
			x_offset = 0
			x = width
		} else {
			x_offset = half_w - half_x
		}
		if y > height {
			y = height
		}
		y_offset = 0
	case position == POSITION_RIGHT_TOP:
		if x > width {
			x_offset = 0
			x = width
		} else {
			x_offset = width - x
		}
		if y > height {
			y = height
		}
		y_offset = 0
	case position == POSITION_LEFT_MIDDLE:
		x_offset = 0
		if x > width {
			x = width
		}
		if y > height {
			y_offset = 0
			y = height
		} else {
			y_offset = half_h - half_y
		}
	case position == POSITION_CENTER_MIDDLE:
		if x > width {
			x_offset = 0
			x = width
		} else {
			x_offset = half_w - half_x
		}
		if y > height {
			y_offset = 0
			y = height
		} else {
			y_offset = half_h - half_y
		}
	case position == POSITION_RIGHT_MIDDLE:
		if x > width {
			x_offset = 0
			x = width
		} else {
			x_offset = width - x
		}
		if y > height {
			y_offset = 0
			y = height
		} else {
			y_offset = height - y
		}
	case position == POSITION_LEFT_BUTTON:
		x_offset = 0
		if x > width {
			x = width
		}
		if y > height {
			y_offset = 0
			y = height
		} else {
			y_offset = height - y
		}
	case position == POSITION_CENTER_BUTTON:
		if x > width {
			x_offset = 0
			x = width
		} else {
			x_offset = half_w - half_x
		}
		if y > height {
			y_offset = 0
			y = height
		} else {
			y_offset = height - y
		}
	case position == POSITION_RIGHT_BUTTON:
		if x > width {
			x_offset = 0
			x = width
		} else {
			x_offset = width - x
		}
		if y > height {
			y_offset = 0
			y = height
		} else {
			y_offset = height - y
		}
	}
	canvas := CreateCanvas(x, y)
	for i := x_offset; i < x_offset+x; i++ {
		for j := y_offset; j < y_offset+y; j++ {
			r, g, b, a := GetUint8RGBA(img, i, j)
			canvas.SetRGBA(i-x_offset, j-y_offset, color.RGBA{r, g, b, a})
		}
	}
	return canvas
}

func Cut(img image.Image, x int, y int, x_offset int, y_offset int) image.RGBA {
	width, height := GetImageBounds(img)
	if x_offset > width {
		x_offset = x_offset % width
	}
	if y_offset > height {
		y_offset = y_offset % height
	}
	if x_offset+x > width {
		x = width - x_offset
	}
	if y_offset+y > height {
		y = height - y_offset
	}
	canvas := CreateCanvas(x, y)
	for i := x_offset; i < x_offset+x; i++ {
		for j := y_offset; j < y_offset+y; j++ {
			r, g, b, a := GetUint8RGBA(img, i, j)
			canvas.SetRGBA(i-x_offset, j-y_offset, color.RGBA{r, g, b, a})
		}
	}
	return canvas
}
