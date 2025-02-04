package filters

import (
	"math"

	"github.com/jijikhal/GoDIP/pkg/types"
)

// Calculates the gradient size for each pixel. Edge pixels are skipped.
func GradientMagnitude(image *types.GrayImage) *types.FloatImage {
	result := types.MakeFloatImage(image.Height, image.Width, float64(image.MinValue), float64(image.MaxValue))

	for y := 1; y < image.Height-1; y++ {
		for x := 1; x < image.Width-1; x++ {
			xg := image.GetXY(x+1, y) - image.GetXY(x-1, y)
			yg := image.GetXY(x, y+1) - image.GetXY(x, y-1)
			result.SetXY(x, y, math.Sqrt(float64(xg*xg+yg*yg)))
		}
	}

	return result
}

// Calculates the gradient orientation for each pixel. Edge pixels are skipped.
func GradientOrientation(image *types.GrayImage) *types.FloatImage {
	result := types.MakeFloatImage(image.Height, image.Width, -math.Pi, math.Pi)

	for y := 1; y < image.Height-1; y++ {
		for x := 1; x < image.Width-1; x++ {
			xg := image.GetXY(x+1, y) - image.GetXY(x-1, y)
			yg := image.GetXY(x, y+1) - image.GetXY(x, y-1)
			result.SetXY(x, y, math.Atan2(float64(yg), float64(xg)))
		}
	}

	return result
}
