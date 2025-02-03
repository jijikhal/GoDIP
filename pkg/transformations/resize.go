package transformations

import (
	"math"

	"github.com/jijikhal/GoDIP/pkg/types"
)

func ResizeNearest(image *types.GrayImage, width int, height int) *types.GrayImage {
	result := types.MakeGrayImage(height, width, image.MinValue, image.MaxValue)

	for y := range height {
		for x := range width {
			xInOld := math.Round(float64(x) / float64(width) * float64(image.Width-1))
			yInOld := math.Round(float64(y) / float64(height) * float64(image.Height-1))
			result.SetPixelXY(x, y, image.GetXY(int(xInOld), int(yInOld)))
		}
	}

	return result
}
