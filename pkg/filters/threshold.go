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

func BinaryThreshold(img *types.GrayImage, threshold int, mode int) {
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
				newValue = img.MinValue
			} else {
				newValue = value
			}
		}

		img.SetI(i, newValue)
	}
}

func BinaryThreshold2(img *types.GrayImage, t1 int, t2 int, mode int) {
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

		img.SetI(i, newValue)
	}
}

func AdaptiveThreshold(img *types.GrayImage, blockSize int, c float64) {
	means := Convolve(img, kernels.GaussKernel(blockSize), CLOSEST)
	pixelCount := img.GetPixelCount()
	for i := 0; i < pixelCount; i++ {
		value := img.GetI(i)
		thresh := int(means.GetI(i) - c)

		if value >= thresh {
			img.SetI(i, img.MaxValue)
		} else {
			img.SetI(i, img.MinValue)
		}
	}
}
