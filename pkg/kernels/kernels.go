package kernels

import (
	"math"

	"github.com/jijikhal/GoDIP/pkg/types"
)

// Generates a box kernel for mean filtering
func BoxKernel(width int, height int) *types.FloatImage {
	result := types.MakeFloatImage(height, width, 0, 1)
	for i := range width * height {
		result.SetI(i, 1.0/float64(width*height))
	}

	return result
}

// Generates gaussian kernel of size using some variance.
// If sigma is zero or less, suitable sigma is calculated
func GaussKernel(size int, sigma float64) *types.FloatImage {
	if sigma <= 0 {
		sigma = 0.3*(float64(size-1)*0.5-1) + 0.8
	}
	result := types.MakeFloatImage(size, size, 0, 1)
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

// Generates binary rectangle kernel
func OnesKernel(width int, height int) *types.GrayImage {
	result := types.MakeGrayImage(height, width, 0, 1)
	for i := range width * height {
		result.SetI(i, 1)
	}

	return result
}

// Generates a binary circular kernel
func CircleKernel(size int) *types.GrayImage {
	result := types.MakeGrayImage(size, size, 0, 1)

	center := float64(size-1) / 2.0

	for y := range size {
		for x := range size {
			if math.Pow(float64(y)-center, 2)+math.Pow(float64(x)-center, 2) <= math.Pow(float64(size)/2, 2) {
				result.SetPixelXY(x, y, 1)
			}
		}
	}

	return result
}
