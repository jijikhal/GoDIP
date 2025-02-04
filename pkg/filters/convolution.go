package filters

import (
	"github.com/jijikhal/GoDIP/pkg/types"
)

const (
	ZEROS = iota
	CLOSEST
	REPEAT
)

func pmod(x, d int) int {
	x = x % d
	if x >= 0 {
		return x
	}
	return x + d
}

func getValueInt(image *types.GrayImage, x int, y int, mode int) int {
	if x >= 0 && x < image.Width && y >= 0 && y < image.Height {
		return image.GetXY(x, y)
	}
	switch mode {
	case ZEROS:
		return 0
	case CLOSEST:
		return image.GetXY(max(0, min(x, image.Width-1)), max(0, min(y, image.Height-1)))
	case REPEAT:
		return image.GetXY(pmod(x, image.Width), pmod(y, image.Height))
	}
	return 0
}

// Applies convolution to image. You can also choose how to treat edges using the `mode` parameter.
// ZEROS means out-of-image coordinates are filled with zeros, CLOSEST uses the closest pixel value
// and REPET repets the image cyclicaly
func Convolve(image *types.GrayImage, kernel *types.FloatImage, mode int) *types.FloatImage {
	result := types.MakeFloatImage(image.Height, image.Width, float64(image.MinValue), float64(image.MaxValue))

	for y := 0; y < image.Height; y++ {
		for x := 0; x < image.Width; x++ {
			var sum float64
			offsetX := (kernel.Width - 1) / 2
			offsetY := (kernel.Height - 1) / 2
			for ky := 0; ky < kernel.Height; ky++ {
				for kx := 0; kx < kernel.Width; kx++ {
					sum += kernel.GetXY(kx, ky) * float64(getValueInt(image, x-kx+offsetX, y-ky+offsetY, mode))
				}
			}
			result.SetXY(x, y, sum)
		}
	}
	return result
}
