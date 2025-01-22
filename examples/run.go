package main

import (
	"fmt"

	loading "github.com/jijikhal/GoDIP/internal/loading"
	saving "github.com/jijikhal/GoDIP/internal/saving"
	"github.com/jijikhal/GoDIP/pkg/filters"
)

func main() {
	image, err := loading.Load("ada_threshold.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	blue := image.GetChannel(0)

	filters.AdaptiveThreshold(blue, 15, 2)

	//saving.SaveAsPPM("test3.ppm", image)
	saving.Save("test3.jpg", blue.ToColor())
}
