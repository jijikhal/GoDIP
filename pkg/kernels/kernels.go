package kernels

import (
	"math"

	"github.com/jijikhal/GoDIP/pkg/types"
)

func BoxKernel(width int, height int) *types.FloatImage {
	result := types.MakeFloatImage(height, width, 1, 0)
	for i := range width * height {
		result.SetI(i, 1.0/float64(width*height))
	}

	return result
}

func GaussKernel(size int) *types.FloatImage {
	sigma := 0.3*(float64(size-1)*0.5-1) + 0.8
	result := types.MakeFloatImage(size, size, 1, 0)
	offset := (size - 1) / 2
	sum := 0.0
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			xr := x - offset
			yr := y - offset
			value := math.Exp(-float64(xr*xr+yr*yr) / (2 * sigma * sigma))
			result.SetPixelXY(x, y, value)
			sum += value
		}
	}
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			value := result.GetXY(x, y)
			result.SetPixelXY(x, y, value/sum)
		}
	}

	return result
}
