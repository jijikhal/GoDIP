package filters

import (
	"math"

	"github.com/jijikhal/GoDIP/pkg/transformations"
	"github.com/jijikhal/GoDIP/pkg/types"
)

// Applies minimum filter to image using a binary kernel
func MinFilter(image *types.GrayImage, kernel *types.GrayImage) *types.GrayImage {
	result := types.MakeGrayImage(image.Height, image.Width, image.MinValue, image.MaxValue)

	for y := 0; y < image.Height; y++ {
		for x := 0; x < image.Width; x++ {
			var minValue = math.MaxInt
			offsetX := (kernel.Width - 1) / 2
			offsetY := (kernel.Height - 1) / 2

			minX := max(0, x-offsetX)
			maxX := min(image.Width-1, x+offsetX)
			minY := max(0, y-offsetY)
			maxY := min(image.Height-1, y+offsetY)

			for ky := 0; ky < kernel.Height; ky++ {
				for kx := 0; kx < kernel.Width; kx++ {
					if kernel.GetXY(kx, ky) <= 0 || y+ky-offsetY > maxY || y+ky-offsetY < minY || x+kx-offsetX > maxX || x+kx-offsetX < minX {
						continue
					}

					value := image.GetXY(x+kx-offsetX, y+ky-offsetY)
					if value < minValue {
						minValue = value
					}
				}
			}
			result.SetXY(x, y, minValue)
		}
	}
	return result
}

// Applies maximum filter to image using a binary kernel
func MaxFilter(image *types.GrayImage, kernel *types.GrayImage) *types.GrayImage {
	result := types.MakeGrayImage(image.Height, image.Width, image.MinValue, image.MaxValue)

	for y := 0; y < image.Height; y++ {
		for x := 0; x < image.Width; x++ {
			var maxValue = math.MinInt
			offsetX := (kernel.Width - 1) / 2
			offsetY := (kernel.Height - 1) / 2

			minX := max(0, x-offsetX)
			maxX := min(image.Width-1, x+offsetX)
			minY := max(0, y-offsetY)
			maxY := min(image.Height-1, y+offsetY)

			for ky := 0; ky < kernel.Height; ky++ {
				for kx := 0; kx < kernel.Width; kx++ {
					if kernel.GetXY(kx, ky) <= 0 || y+ky-offsetY > maxY || y+ky-offsetY < minY || x+kx-offsetX > maxX || x+kx-offsetX < minX {
						continue
					}

					value := image.GetXY(x+kx-offsetX, y+ky-offsetY)
					if value > maxValue {
						maxValue = value
					}
				}
			}
			result.SetXY(x, y, maxValue)
		}
	}
	return result
}

// Applies erosion to binary image using binary kernel
func Erode(image *types.GrayImage, kernel *types.GrayImage) *types.GrayImage {
	return binaryMorphology(image, kernel, 0)
}

// Applies dilatation to binary image using binary kernel
func Dilatate(image *types.GrayImage, kernel *types.GrayImage) *types.GrayImage {
	return binaryMorphology(image, kernel, 1)
}

func binaryMorphology(image *types.GrayImage, kernel *types.GrayImage, criticalValue int) *types.GrayImage {
	result := types.MakeGrayImage(image.Height, image.Width, image.MinValue, image.MaxValue)

	for y := 0; y < image.Height; y++ {
		for x := 0; x < image.Width; x++ {
			offsetX := (kernel.Width - 1) / 2
			offsetY := (kernel.Height - 1) / 2

			minX := max(0, x-offsetX)
			maxX := min(image.Width-1, x+offsetX)
			minY := max(0, y-offsetY)
			maxY := min(image.Height-1, y+offsetY)

			for ky := 0; ky < kernel.Height; ky++ {
				for kx := 0; kx < kernel.Width; kx++ {
					if kernel.GetXY(kx, ky) <= 0 || y+ky-offsetY > maxY || y+ky-offsetY < minY || x+kx-offsetX > maxX || x+kx-offsetX < minX {
						continue
					}

					value := image.GetXY(x+kx-offsetX, y+ky-offsetY)
					if value == criticalValue {
						result.SetXY(x, y, criticalValue)
						ky = kernel.Height
						break
					}
				}
			}
		}
	}
	return result
}

// Applies closing to binary image using binary kernel
func Close(image *types.GrayImage, kernel *types.GrayImage) *types.GrayImage {
	result := Dilatate(image, kernel)
	return Erode(result, transformations.FlipXY(kernel))
}

// Applies opening to binary image using binary kernel
func Open(image *types.GrayImage, kernel *types.GrayImage) *types.GrayImage {
	result := Erode(image, kernel)
	return Dilatate(result, transformations.FlipXY(kernel))
}
