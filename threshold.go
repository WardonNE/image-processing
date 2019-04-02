package main

import (
	_ "fmt"
	"image"
	"math"
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

func OneDimensionalMaxEntropyThreshosd(sgray [256]int, img image.Image) int {
	var p, q [256]float64
	var c0, c1 [256][256]float64
	var h0, h1 [256][256]float64
	var H [256]float64
	var hsum0, hsum1 [256]float64
	var dx, dy int = GetImageBounds(img)
	for i := 0; i < 256; i++ {
		p[i] = 0.0
		q[i] = 0.0
		hsum0[i] = 0.0
		hsum1[i] = 0.0
		H[i] = 0.0
		for j := 0; j < 256; j++ {
			c0[i][j] = 0.0
			c1[i][j] = 0.0
			h0[i][j] = 0.0
			h1[i][j] = 0.0
		}
	}
	for gray, count := range sgray {
		p[gray] = float64(count) / float64(dx*dy)
		q[gray] = 1 - p[gray]
	}
	for i := 0; i < 256; i++ {
		for j := 1; j < i; j++ {
			if p[i] > 0.0 {
				c0[i][j] = p[j] / p[i]
			} else {
				c0[i][j] = 0
			}
			for k := i + 1; k < 256; k++ {
				if q[i] > 0.0 {
					c1[i][k] = p[k] / q[i]
				} else {
					c1[i][k] = 0
				}
			}
		}
	}
	for i := 1; i < 256; i++ {
		for j := 1; j < i; j++ {
			if c0[i][j] != 0.0 {
				h0[i][j] = -c0[i][j] * math.Log10(c0[i][j])
			}
			for k := i + 1; k < 256; k++ {
				if c1[i][k] != 0.0 {
					h1[i][k] = -c1[i][k] * math.Log10(c1[i][k])
				}
			}
		}
	}
	for k, arr := range h0 {
		for _, value := range arr {
			hsum0[k] += value
		}
	}
	for k, arr := range h1 {
		for _, value := range arr {
			hsum1[k] += value
		}
	}
	for i := 0; i < 256; i++ {
		H[i] = (hsum0[i] + hsum1[i])
	}
	var th int
	var max float64 = 0.0
	for gray, value := range H {
		if value > max {
			max = value
			th = gray
		}
	}
	return th
}

func PParamMethodThreshosd(sgray [256]int, p float64, t float64) int {
	var th int = 0
	var P float64 = 0.0
	for g := 1; g < 254; g++ {
		var Pasum, Pbsum int = 0, 0
		for gray, count := range sgray {
			if gray > g {
				Pasum += count
			} else {
				Pbsum += count
			}
		}
		P = float64(Pasum) / float64(Pasum+Pbsum)
		if math.Abs(P-p) < t {
			th = g
			break
		} else {
			Pasum = 0
			Pbsum = 0
			P = 0.0
		}
	}
	return th
}
