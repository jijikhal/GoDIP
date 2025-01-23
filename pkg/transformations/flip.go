package transformations

import (
	"github.com/jijikhal/GoDIP/pkg/types"
)

func FlipX(img *types.GrayImage) *types.GrayImage {
	result := types.MakeGrayImage(img.Height, img.Width, img.MinValue, img.MaxValue)
	halfWidth := img.Width / 2
	for y := 0; y < img.Height; y++ {
		for x := 0; x < halfWidth; x++ {
			result.SetXY(x, y, img.GetXY(img.Width-x-1, y))
			result.SetXY(img.Width-x-1, y, img.GetXY(x, y))
		}
	}

	return result
}

func FlipY(img *types.GrayImage) *types.GrayImage {
	result := types.MakeGrayImage(img.Height, img.Width, img.MinValue, img.MaxValue)
	halfHeight := img.Height / 2
	for y := 0; y < halfHeight; y++ {
		for x := 0; x < img.Width; x++ {
			result.SetXY(x, y, img.GetXY(x, img.Height-y-1))
			result.SetXY(x, img.Height-y-1, img.GetXY(x, y))
		}
	}

	return result
}

func FlipXY(img *types.GrayImage) *types.GrayImage {
	result := types.MakeGrayImage(img.Height, img.Width, img.MinValue, img.MaxValue)
	pixelCount := img.GetPixelCount()
	half := pixelCount / 2

	for i := range half {
		result.SetI(i, img.GetI(pixelCount-i-1))
		result.SetI(pixelCount-i-1, img.GetI(i))
	}

	return result
}

func Rotate90(img *types.GrayImage) *types.GrayImage {
	result := types.MakeGrayImage(img.Width, img.Height, img.MinValue, img.MaxValue)
	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			result.SetXY(img.Height-1-y, x, img.GetXY(x, y))
		}
	}

	return result
}

func Rotate270(img *types.GrayImage) *types.GrayImage {
	result := types.MakeGrayImage(img.Width, img.Height, img.MinValue, img.MaxValue)
	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			result.SetXY(y, img.Width-1-x, img.GetXY(x, y))
		}
	}

	return result
}

func Rotate180(img *types.GrayImage) *types.GrayImage {
	result := types.MakeGrayImage(img.Height, img.Width, img.MinValue, img.MaxValue)
	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			result.SetXY(img.Width-1-x, img.Height-1-y, img.GetXY(x, y))
		}
	}

	return result
}
