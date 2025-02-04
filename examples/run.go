package main

import (
	"fmt"

	loading "github.com/jijikhal/GoDIP/pkg/loading"
	saving "github.com/jijikhal/GoDIP/pkg/saving"
	"github.com/jijikhal/GoDIP/pkg/transformations"
)

func main() {
	image, err := loading.Load("doge.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	red := image.GetChannel(0)

	//thresh := filters.AdaptiveThreshold(red, 15, 0.5)

	resized := transformations.ResizeNearest(red, 10000, 10000)

	//saving.SaveAsPPM("test3.ppm", image)
	saving.SaveGray("test3.jpg", resized)
}
