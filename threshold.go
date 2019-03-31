package main

import (
	_ "fmt"
	_ "image"
	"math"
)

const (
	DEFAULT_THRESHOSD int = 127
)

func DefaultThreshosd() int {
	return DEFAULT_THRESHOSD
}

func MeanThreshosd(sgray [256]int) int {
	var total int = 0
	var num int = 0
	for gray, count := range sgray {
		total += gray * count
		num += count
	}
	return int(total / num)
}

func TwoPeakThreshosd(sgray [256]int) int {
	var h1, h2, t1, t2, th int = 0, 0, 0, 0, 0
	for i := 0; i < 256; i++ {
		if i <= 128 {
			if sgray[i] > t1 {
				t1 = sgray[i]
				h1 = i
			}
		} else {
			if sgray[i] > t2 {
				t2 = sgray[i]
				h2 = i
			}
		}
	}
	var button int = int(math.Min(float64(t1), float64(t2)))
	for k := h1; k <= h2; k++ {
		if sgray[k] < button {
			th = k
			button = sgray[k]
		}
	}
	return th
}
