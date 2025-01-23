package filters

import (
	"github.com/jijikhal/GoDIP/pkg/kernels"
	"github.com/jijikhal/GoDIP/pkg/types"
)

func GaussianBlur(img *types.GrayImage, size int, sigma float64) *types.FloatImage {
	return Convolve(img, kernels.GaussKernel(size, sigma), CLOSEST)
}

func BoxBlur(img *types.GrayImage, size int) *types.FloatImage {
	return Convolve(img, kernels.BoxKernel(size, size), CLOSEST)
}
