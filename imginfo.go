package main

import (
	"fmt"
	"image"
	"os"
)

func GetImageBounds(img image.Image) (int, int) {
	bounds := img.Bounds()
	return bounds.Dx(), bounds.Dy()
}

func GetImageExt(file string) (string, error) {
	var header = make([]byte, 8)
	var ext string
	if f, err := os.Open(file); err == nil {
		defer f.Close()
		if _, err := f.Read(header); err == nil {
			xstr := fmt.Sprintf("%x", header)
			switch {
			case xstr == "89504e470d0a1a0a":
				ext = ".png"
			case xstr == "0000010001002020":
				ext = ".ico"
			case xstr == "0000020001002020":
				ext = ".cur"
			case xstr[:12] == "474946383961" || xstr[:12] == "474946383761":
				ext = ".gif"
			case xstr[:10] == "0000020000" || xstr[:10] == "0000100000":
				ext = ".tga"
			case xstr[:8] == "464f524d":
				ext = ".iff"
			case xstr[:8] == "52494646":
				ext = ".ani"
			case xstr[:4] == "4d4d" || xstr[:4] == "4949":
				ext = ".tiff"
			case xstr[:4] == "424d":
				ext = ".bmp"
			case xstr[:4] == "ffd8":
				ext = ".jpg"
			case xstr[:2] == "0a":
				ext = ".pcx"
			default:
				ext = ""
			}
			return ext, nil
		} else {
			return "", err
		}
	} else {
		return "", err
	}

}
