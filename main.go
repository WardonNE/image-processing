package main

import (
	"fmt"
	"image"
)

var (
	filename string = "E:/goworkspace/bin/lena.jpg"
	// filename string = "F:/GoWorkspace/bin/demo.jpg"
	canvas image.RGBA
	dir    string = "E:/goworkspace/bin/images/"
	// dir     string  = "F:/GoWorkspace/bin/images/"
	mode            = ASH_WEIGHTED_MEAN
	P_Param float64 = 0.75
)

func main() {
	fmt.Println("Loading Picture...")
	img, err := NewImage(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Loaded!!")

	// fmt.Println("Starting Quick Cut Picture...")
	// canvas = QuickCut(img, 1)
	// SaveImage(dir+"quick_cut/left_top.png", canvas)
	// canvas = QuickCut(img, 2)
	// SaveImage(dir+"quick_cut/center_top.png", canvas)
	// canvas = QuickCut(img, 3)
	// SaveImage(dir+"quick_cut/right_top.png", canvas)
	// canvas = QuickCut(img, 4)
	// SaveImage(dir+"quick_cut/left_middle.png", canvas)
	// canvas = QuickCut(img, 5)
	// SaveImage(dir+"quick_cut/center_middle.png", canvas)
	// canvas = QuickCut(img, 6)
	// SaveImage(dir+"quick_cut/right_middle.png", canvas)
	// canvas = QuickCut(img, 7)
	// SaveImage(dir+"quick_cut/left_button.png", canvas)
	// canvas = QuickCut(img, 8)
	// SaveImage(dir+"quick_cut/center_button.png", canvas)
	// canvas = QuickCut(img, 9)
	// SaveImage(dir+"quick_cut/right_button.png", canvas)
	// fmt.Println("Quick Cut Picture Done!!")

	// fmt.Println("Starting Simple Cut Picture...")
	// canvas = SimpleCut(img, 360, 240, 1)
	// SaveImage(dir+"simple_cut/left_top.png", canvas)
	// canvas = SimpleCut(img, 360, 240, 2)
	// SaveImage(dir+"simple_cut/center_top.png", canvas)
	// canvas = SimpleCut(img, 360, 240, 3)
	// SaveImage(dir+"simple_cut/right_top.png", canvas)
	// canvas = SimpleCut(img, 360, 240, 4)
	// SaveImage(dir+"simple_cut/left_middle.png", canvas)
	// canvas = SimpleCut(img, 360, 240, 5)
	// SaveImage(dir+"simple_cut/center_middle.png", canvas)
	// canvas = SimpleCut(img, 360, 240, 6)
	// SaveImage(dir+"simple_cut/right_middle.png", canvas)
	// canvas = SimpleCut(img, 360, 240, 7)
	// SaveImage(dir+"simple_cut/left_button.png", canvas)
	// canvas = SimpleCut(img, 360, 240, 8)
	// SaveImage(dir+"simple_cut/center_button.png", canvas)
	// canvas = SimpleCut(img, 360, 240, 9)
	// SaveImage(dir+"simple_cut/right_button.png", canvas)
	// fmt.Println("Simple Cut Picture Done!!")

	// fmt.Println("Starting Cut Picture...")
	// canvas = Cut(img, 720, 540, 200, 100)
	// SaveImage(dir+"cut/cut.png", canvas)
	// fmt.Println("Cut Picture Done!!")

	// fmt.Println("Starting Ash Picture...")
	// canvas = AvgAsh(img)
	// SaveImage(dir+"ash/avgash.png", canvas)
	// canvas = PsAsh(img)
	// SaveImage(dir+"ash/psash.png", canvas)
	// canvas = WeightedMeanAsh(img)
	// SaveImage(dir+"ash/weightedmeanash.png", canvas)
	// canvas = RChannelAsh(img)
	// SaveImage(dir+"ash/rchannelash.png", canvas)
	// canvas = GChannelAsh(img)
	// SaveImage(dir+"ash/gchannelash.png", canvas)
	// canvas = BChannelAsh(img)
	// SaveImage(dir+"ash/bchannelash.png", canvas)
	// fmt.Println("Ash Picture Done!!")

	// var threshosd int
	// var sgray [256]int
	// fmt.Println("Starting Binary Picture...")
	// sgray = GrayLevelHistWeightedMean(img)
	// threshosd = DefaultThreshosd()
	// canvas = ImageBinary(img, threshosd, mode)
	// SaveImage(dir+"binaryzation/defaultbinaryzation.png", canvas)
	// threshosd = MeanThreshosd(sgray)
	// canvas = ImageBinary(img, threshosd, mode)
	// SaveImage(dir+"binaryzation/meanbinaryzation.png", canvas)
	// threshosd = TwoPeakThreshosd(sgray)
	// canvas = ImageBinary(img, threshosd, mode)
	// SaveImage(dir+"binaryzation/2peakbinaryzation.png", canvas)
	// threshosd = IterationThreshosd(sgray, 2)
	// canvas = ImageBinary(img, threshosd, mode)
	// SaveImage(dir+"binaryzation/iterationbinaryzation.png", canvas)
	// mat := GetGrayMat(img, 4)
	// threshosd = SimpleCensusThreshosd(mat)
	// canvas = ImageBinary(img, threshosd, mode)
	// SaveImage(dir+"binaryzation/simplecensusbinaryzation.png", canvas)
	// threshosd = OstuThreshosd(sgray)
	// canvas = ImageBinary(img, threshosd, mode)
	// SaveImage(dir+"binaryzation/otsubinaryzation.png", canvas)
	// threshosd = OneDimensionalMaxEntropyThreshosd(sgray, img)
	// canvas = ImageBinary(img, threshosd, mode)
	// SaveImage(dir+"binaryzation/onedimensionalmaxentropybinaryzation.png", canvas)
	// threshosd = PParamMethodThreshosd(sgray, P_Param, 0.01)
	// canvas = ImageBinary(img, threshosd, mode)
	// SaveImage(dir+"binaryzation/pparammethodbinaryzation.png", canvas)
	// fmt.Println("Binary Picture Done!!")

	// fmt.Println("Add Salt Peper Noise Start...")
	// canvas = AddSaltNoise(img, 0.30)
	// SaveImage(dir+"salt/30salt.png", canvas)
	// canvas = AddSaltNoise(img, 0.50)
	// SaveImage(dir+"salt/50salt.png", canvas)
	// canvas = AddSaltNoise(img, 0.80)
	// SaveImage(dir+"salt/80salt.png", canvas)
	// canvas = AddPeperNoise(img, 0.30)
	// SaveImage(dir+"peper/30peper.png", canvas)
	// canvas = AddPeperNoise(img, 0.50)
	// SaveImage(dir+"peper/50peper.png", canvas)
	// canvas = AddPeperNoise(img, 0.80)
	// SaveImage(dir+"peper/80peper.png", canvas)
	// canvas = AddSaleAndPeperNoist(img, 0.30)
	// SaveImage(dir+"saltandpeper/30saltandpeper.png", canvas)
	// canvas = AddSaleAndPeperNoist(img, 0.50)
	// SaveImage(dir+"saltandpeper/50saltandpeper.png", canvas)
	// canvas = AddSaleAndPeperNoist(img, 0.80)
	// SaveImage(dir+"saltandpeper/80saltandpeper.png", canvas)
	// fmt.Println("Add Salt Peper Noise Done!!")

	// fmt.Println("Add Gauss Noise Start...")
	// canvas = AddGaussNoise(img, 0.0, 1.0, 8)
	// SaveImage(dir+"gaussnoise/8gauss.png", canvas)
	// canvas = AddGaussNoise(img, 0.0, 1.0, 16)
	// SaveImage(dir+"gaussnoise/16gauss.png", canvas)
	// canvas = AddGaussNoise(img, 0.0, 1.0, 32)
	// SaveImage(dir+"gaussnoise/32gauss.png", canvas)
	// canvas = AddGaussNoise(img, 0.0, 1.0, 64)
	// SaveImage(dir+"gaussnoise/64gauss.png", canvas)
	// canvas = AddGaussNoise(img, 0.0, 1.0, 128)
	// SaveImage(dir+"gaussnoise/128gauss.png", canvas)
	// canvas = AddGaussNoise(img, 0.0, 1.0, 256)
	// SaveImage(dir+"gaussnoise/256gauss.png", canvas)
	// fmt.Println("Add Gauss Noise Done!!")

	// fmt.Println("Add Poisson Noise Start...")
	// canvas = AddPoissonNoise(img, 9, 2)
	// SaveImage(dir+"poissonnoise/9poisson.png", canvas)
	// canvas = AddPoissonNoise(img, 16, 2)
	// SaveImage(dir+"poissonnoise/16poisson.png", canvas)
	// canvas = AddPoissonNoise(img, 25, 2)
	// SaveImage(dir+"poissonnoise/25poisson.png", canvas)
	// canvas = AddPoissonNoise(img, 36, 2)
	// SaveImage(dir+"poissonnoise/36poisson.png", canvas)
	// canvas = AddPoissonNoise(img, 49, 2)
	// SaveImage(dir+"poissonnoise/49poisson.png", canvas)
	// canvas = AddPoissonNoise(img, 100, 2)
	// SaveImage(dir+"poissonnoise/100poisson.png", canvas)
	// fmt.Println("Add Poisson Noise Done!!")

	fmt.Println("Resize Picture Start...")
	// canvas = ScaleImage(img, 0.1, 1)
	// SaveImage(dir+"scale/0.1nearestneightor.png", canvas)
	// canvas = ScaleImage(img, 0.5, 1)
	// SaveImage(dir+"scale/0.5nearestneightor.png", canvas)
	// canvas = ScaleImage(img, 1.5, 1)
	// SaveImage(dir+"scale/1.5nearestneightor.png", canvas)
	// canvas = ScaleImage(img, 2.0, 1)
	// SaveImage(dir+"scale/2.0nearestneightor.png", canvas)

	// canvas = ScaleImage(img, 0.1, 2)
	// SaveImage(dir+"scale/0.1bilinear.png", canvas)
	// canvas = ScaleImage(img, 0.5, 2)
	// SaveImage(dir+"scale/0.5bilinear.png", canvas)
	// canvas = ScaleImage(img, 1.5, 2)
	// SaveImage(dir+"scale/1.5bilinear.png", canvas)
	// canvas = ScaleImage(img, 2.0, 2)
	// SaveImage(dir+"scale/2.0bilinear.png", canvas)

	canvas = ScaleImage(img, 0.1, 3)
	SaveImage(dir+"scale/0.1bicubic.png", canvas)
	canvas = ScaleImage(img, 0.5, 3)
	SaveImage(dir+"scale/0.5bicubic.png", canvas)
	canvas = ScaleImage(img, 1.5, 3)
	SaveImage(dir+"scale/1.5bicubic.png", canvas)
	canvas = ScaleImage(img, 2.0, 3)
	SaveImage(dir+"scale/2.0bicubic.png", canvas)
	fmt.Println("Resize Picture Done!!")
}
