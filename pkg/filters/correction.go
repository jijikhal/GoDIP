package filters

import (
	"math"

	"github.com/jijikhal/GoDIP/pkg/types"
)

func ChangeBrightness(image *types.GrayImage, change int) *types.GrayImage {
	result := image.Duplicate()
	for i := range result.GetPixelCount() {
		result.SetI(i, min(max(image.GetI(i)+change, image.MinValue), image.MaxValue))
	}
	return result
}

func ChangeContrast(image *types.GrayImage, coefficent float64) *types.GrayImage {
	result := image.Duplicate()
	for i := range result.GetPixelCount() {
		result.SetI(i, min(max(int(math.Round(float64(image.GetI(i))*coefficent)), image.MinValue), image.MaxValue))
	}
	return result
}

func GammaCorrection(image *types.GrayImage, gamma float64) *types.GrayImage {
	result := image.Duplicate()
	for i := range result.GetPixelCount() {
		val := result.GetI(i)
		valNormalized := float64(val-image.MinValue) / float64(image.MaxValue)
		gammaApplied := math.Pow(valNormalized, gamma)

		result.SetI(i, min(max(int(math.Round(gammaApplied*float64(image.MaxValue)+float64(image.MinValue))), image.MinValue), image.MaxValue))
	}
	return result
}
