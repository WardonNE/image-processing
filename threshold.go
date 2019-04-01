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

func IterationThreshosd(sgray [256]int, t int) int {
	var th0, th1 int = 256, 0
	var c0, c1 int = 0, 0
	var s0, s1 int = 0, 0
	var th int = MeanThreshosd(sgray)
	for true {
		for gray, count := range sgray {
			if gray <= th {
				s0 += (gray * count)
				c0 += count
			} else {
				s1 += (gray * count)
				c1 += count
			}
		}
		th1 = int((s0/c0 + s1/c1) / 2.0)
		if int(math.Abs(float64(th1-th0))) <= t {
			break
		} else {
			th0 = th1
			c1 = 0
			c0 = 0
			s1 = 0
			s0 = 0
			th = th1
		}
	}
	return th
}

func SimpleCensusThreshosd(mat [][]int) int {
	var Ex, Ey, Exy, th, Esum, Efsum int
	for i := 1; i < len(mat)-1; i++ {
		for j := 1; j < len(mat[i])-1; j++ {
			Ex = mat[i-1][j] - mat[i+1][j]
			Ey = mat[i][j-1] - mat[i][j+1]
			Exy = int(math.Max(float64(Ex), float64(Ey)))
			Esum += Exy
			Efsum += Exy * mat[i][j]
		}
	}
	th = int(Efsum / Esum)
	return th
}

func OstuThreshosd(sgray [256]int) int {
	var th, n0, n1, s0, s1 int
	var maxotsu int = 0
	var otsu int = 0
	for g := 0; g < 255; g++ {
		for gray, count := range sgray {
			if g > gray {
				n0 += count
				s0 += gray * count
			} else {
				n1 += count
				s1 += gray * count
			}
		}
		p0 := float64(n0) / float64(n0+n1)
		p1 := 1.0 - p0
		var avg0, avg1 float64
		if n0 == 0 {
			avg0 = 0.0
			avg1 = float64(s1 / n1)
		} else if n1 == 0 {
			avg0 = float64(s0 / n0)
			avg1 = 0.0
		} else {
			avg0 = float64(s0 / n0)
			avg1 = float64(s1 / n1)
		}
		otsu = int(p0 * p1 * math.Pow((avg0-avg1), 2.0))
		if otsu > maxotsu {
			maxotsu = otsu
			th = g
		}
		n0 = 0
		n1 = 0
		s0 = 0
		s1 = 0
		p0 = 0.0
		p1 = 0.0
		avg0 = 0.0
		avg1 = 0.0
		otsu = 0
	}
	return th
}

func OneDimensionalMaxEntropyThreshosd(sgray [256]int) int {
	for g := 0; g < 256; i++ {
		for gray, count := range sgray {
			if true {

			}
		}
	}
}
