package filters

import (
	"github.com/jijikhal/GoDIP/pkg/kernels"
	"github.com/jijikhal/GoDIP/pkg/types"
)

const (
	BINARY = iota
	BINARY_INV
	TRUNC
	TRUNC_INV
	TO_ZERO
	TO_ZERO_INV
)

// Thresholds image using a single threshold. A mode can be chosen:
// - BINARY sets all pixels above threshold to max, all below to min
// - BINARY_INV sets all pixels above threshold to min, all below to max
// - TRUNC sets all pixels above threshold to threshold value
// - TRUNC_INV sets all pixels below threshold to threshold value
// - TO_ZERO sets all pixels below threshold to min and keeps others
// - TO_ZERO_INV sets all pixels above threshold to max and keeps others
func BinaryThreshold(img *types.GrayImage, threshold int, mode int) *types.GrayImage {
	result := types.MakeGrayImage(img.Height, img.Width, img.MinValue, img.MaxValue)
	pixelCount := img.GetPixelCount()
	for i := 0; i < pixelCount; i++ {
		value := img.GetI(i)
		var newValue int

		switch mode {
		case BINARY:
			if value >= threshold {
				newValue = img.MaxValue
			} else {
				newValue = img.MinValue
			}
		case BINARY_INV:
			if value >= threshold {
				newValue = img.MinValue
			} else {
				newValue = img.MaxValue
			}
		case TRUNC:
			if value >= threshold {
				newValue = threshold
			} else {
				newValue = value
			}
		case TRUNC_INV:
			if value < threshold {
				newValue = threshold
			} else {
				newValue = value
			}
		case TO_ZERO:
			if value >= threshold {
				newValue = value
			} else {
				newValue = img.MinValue
			}
		case TO_ZERO_INV:
			if value >= threshold {
				newValue = img.MaxValue
			} else {
				newValue = value
			}
		}

		result.SetI(i, newValue)
	}

	return result
}

// Thresholds image using a two thresholds. A mode can be chosen:
// - BINARY sets all pixels inside range to max, all outside to min
// - BINARY_INV sets all pixels in range to min, all outside to max
// - TRUNC clamps all pixel into the threshold range
// - TRUNC_INV sets all pixel outisde the range to min or max
// - TO_ZERO sets all pixels outside range to min and keeps others
// - TO_ZERO_INV sets all pixels inside to max and keeps others
func BinaryThreshold2(img *types.GrayImage, t1 int, t2 int, mode int) *types.GrayImage {
	result := types.MakeGrayImage(img.Height, img.Width, img.MinValue, img.MaxValue)
	pixelCount := img.GetPixelCount()
	for i := 0; i < pixelCount; i++ {
		value := img.GetI(i)
		var newValue int

		switch mode {
		case BINARY:
			if value >= t1 && value < t2 {
				newValue = img.MaxValue
			} else {
				newValue = img.MinValue
			}
		case BINARY_INV:
			if value >= t1 && value < t2 {
				newValue = img.MinValue
			} else {
				newValue = img.MaxValue
			}
		case TRUNC:
			if value < t1 {
				newValue = t1
			} else if value > t2 {
				newValue = t2
			} else {
				newValue = value
			}
		case TRUNC_INV:
			if value < t1 || value > t2 {
				newValue = value
			} else {
				newValue = img.MinValue
			}
		case TO_ZERO:
			if value >= t1 && value < t2 {
				newValue = value
			} else {
				newValue = img.MinValue
			}
		case TO_ZERO_INV:
			if value >= t1 && value < t2 {
				newValue = img.MinValue
			} else {
				newValue = value
			}
		}

		result.SetI(i, newValue)
	}

	return result
}

// Thresholds image using Gaussian adaptive thresholding using kernel of size `blockSize`
// All pixels that are at least `c` larger than average of neighbourhood are set to max
func AdaptiveThreshold(img *types.GrayImage, blockSize int, c float64) *types.GrayImage {
	means := Convolve(img, kernels.GaussKernel(blockSize, -1), CLOSEST)
	result := types.MakeGrayImage(img.Height, img.Width, img.MinValue, img.MaxValue)
	pixelCount := img.GetPixelCount()
	for i := 0; i < pixelCount; i++ {
		value := img.GetI(i)
		thresh := int(means.GetI(i) - c)

		if value >= thresh {
			result.SetI(i, img.MaxValue)
		} else {
			result.SetI(i, img.MinValue)
		}
	}

	return result
}
