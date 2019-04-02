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
	// dir    string = "F:/GoWorkspace/bin/images/"
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

	fmt.Println("Starting Quick Cut Picture...")
	canvas = QuickCut(img, 1)
	SaveImage(dir+"quick_left_top.png", canvas)
	canvas = QuickCut(img, 2)
	SaveImage(dir+"quick_center_top.png", canvas)
	canvas = QuickCut(img, 3)
	SaveImage(dir+"quick_right_top.png", canvas)
	canvas = QuickCut(img, 4)
	SaveImage(dir+"quick_left_middle.png", canvas)
	canvas = QuickCut(img, 5)
	SaveImage(dir+"quick_center_middle.png", canvas)
	canvas = QuickCut(img, 6)
	SaveImage(dir+"quick_right_middle.png", canvas)
	canvas = QuickCut(img, 7)
	SaveImage(dir+"quick_left_button.png", canvas)
	canvas = QuickCut(img, 8)
	SaveImage(dir+"quick_center_button.png", canvas)
	canvas = QuickCut(img, 9)
	SaveImage(dir+"quick_right_button.png", canvas)
	fmt.Println("Quick Cut Picture Done!!")

	fmt.Println("Starting Simple Cut Picture...")
	canvas = SimpleCut(img, 720, 540, 1)
	SaveImage(dir+"simple_left_top.png", canvas)
	canvas = SimpleCut(img, 720, 540, 2)
	SaveImage(dir+"simple_center_top.png", canvas)
	canvas = SimpleCut(img, 720, 540, 3)
	SaveImage(dir+"simple_right_top.png", canvas)
	canvas = SimpleCut(img, 720, 540, 4)
	SaveImage(dir+"simple_left_middle.png", canvas)
	canvas = SimpleCut(img, 720, 540, 5)
	SaveImage(dir+"simple_center_middle.png", canvas)
	canvas = SimpleCut(img, 720, 540, 6)
	SaveImage(dir+"simple_right_middle.png", canvas)
	canvas = SimpleCut(img, 720, 540, 7)
	SaveImage(dir+"simple_left_button.png", canvas)
	canvas = SimpleCut(img, 720, 540, 8)
	SaveImage(dir+"simple_center_button.png", canvas)
	canvas = SimpleCut(img, 720, 540, 9)
	SaveImage(dir+"simple_right_button.png", canvas)
	fmt.Println("Simple Cut Picture Done!!")

	fmt.Println("Starting Cut Picture...")
	canvas = Cut(img, 720, 540, 200, 100)
	SaveImage(dir+"cut.png", canvas)
	fmt.Println("Cut Picture Done!!")

	fmt.Println("Starting Ash Picture...")
	canvas = AvgAsh(img)
	SaveImage(dir+"avgash.png", canvas)
	canvas = PsAsh(img)
	SaveImage(dir+"psash.png", canvas)
	canvas = WeightedMeanAsh(img)
	SaveImage(dir+"weightedmeanash.png", canvas)
	canvas = RChannelAsh(img)
	SaveImage(dir+"rchannelash.png", canvas)
	canvas = GChannelAsh(img)
	SaveImage(dir+"gchannelash.png", canvas)
	canvas = BChannelAsh(img)
	SaveImage(dir+"bchannelash.png", canvas)
	fmt.Println("Ash Picture Done!!")

	var threshosd int
	var sgray [256]int
	fmt.Println("Starting Binary Picture...")
	sgray = GrayLevelHistWeightedMean(img)
	threshosd = DefaultThreshosd()
	canvas = ImageBinary(img, threshosd, mode)
	SaveImage(dir+"defaultbinaryzation.png", canvas)
	threshosd = MeanThreshosd(sgray)
	canvas = ImageBinary(img, threshosd, mode)
	SaveImage(dir+"meanbinaryzation.png", canvas)
	threshosd = TwoPeakThreshosd(sgray)
	canvas = ImageBinary(img, threshosd, mode)
	SaveImage(dir+"2peakbinaryzation.png", canvas)
	threshosd = IterationThreshosd(sgray, 2)
	canvas = ImageBinary(img, threshosd, mode)
	SaveImage(dir+"iterationbinaryzation.png", canvas)
	mat := GetGrayMat(img, 4)
	threshosd = SimpleCensusThreshosd(mat)
	canvas = ImageBinary(img, threshosd, mode)
	SaveImage(dir+"simplecensusbinaryzation.png", canvas)
	threshosd = OstuThreshosd(sgray)
	canvas = ImageBinary(img, threshosd, mode)
	SaveImage(dir+"otsubinaryzation.png", canvas)
	threshosd = OneDimensionalMaxEntropyThreshosd(sgray, img)
	canvas = ImageBinary(img, threshosd, mode)
	SaveImage(dir+"onedimensionalmaxentropybinaryzation.png", canvas)
	threshosd = PParamMethodThreshosd(sgray, P_Param, 0.01)
	canvas = ImageBinary(img, threshosd, mode)
	SaveImage(dir+"pparammethodbinaryzation.png", canvas)
	fmt.Println("Binary Picture Done!!")

	fmt.Println("Add Salt Peper Noise Start...")
	canvas = AddSaltNoise(img, 0.30)
	SaveImage(dir+"30salt.png", canvas)
	canvas = AddSaltNoise(img, 0.50)
	SaveImage(dir+"50salt.png", canvas)
	canvas = AddSaltNoise(img, 0.80)
	SaveImage(dir+"80salt.png", canvas)
	canvas = AddPeperNoise(img, 0.30)
	SaveImage(dir+"30peper.png", canvas)
	canvas = AddPeperNoise(img, 0.50)
	SaveImage(dir+"50peper.png", canvas)
	canvas = AddPeperNoise(img, 0.80)
	SaveImage(dir+"80peper.png", canvas)
	canvas = AddSaleAndPeperNoist(img, 0.30)
	SaveImage(dir+"30saltandpeper.png", canvas)
	canvas = AddSaleAndPeperNoist(img, 0.50)
	SaveImage(dir+"50saltandpeper.png", canvas)
	canvas = AddSaleAndPeperNoist(img, 0.80)
	SaveImage(dir+"80saltandpeper.png", canvas)
	fmt.Println("Add Salt Peper Noise Done!!")
}
