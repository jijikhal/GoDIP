package main

import (
	"fmt"

	"github.com/jijikhal/GoDIP/internal"
)

func main() {
	image, err := internal.LoadPPM("test2.ppm")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(image.Data)

	fmt.Printf("Loaded PPM Image: %dx%d\n", image.Width, image.Height)

	internal.SaveAsPPM("test3.ppm", image)
}
