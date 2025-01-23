package main

import (
	"fmt"

	loading "github.com/jijikhal/GoDIP/internal/loading"
	saving "github.com/jijikhal/GoDIP/internal/saving"
	"github.com/jijikhal/GoDIP/pkg/filters"
	"github.com/jijikhal/GoDIP/pkg/kernels"
)

func main() {
	image, err := loading.Load("doge.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	red := image.GetChannel(0)

	//thresh := filters.AdaptiveThreshold(red, 15, 0.5)

	saving.SaveGray("test4.jpg", red)

	kernel := kernels.CircleKernel(11)
	fmt.Println(kernel.Data)

	dilat := filters.MinFilter(red, kernel)

	//saving.SaveAsPPM("test3.ppm", image)
	saving.SaveGray("test3.jpg", dilat)
}
