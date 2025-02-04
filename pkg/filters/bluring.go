package filters

import (
	"github.com/jijikhal/GoDIP/pkg/kernels"
	"github.com/jijikhal/GoDIP/pkg/types"
)

// Applies the Gaussian filter to image `img` using a kernel of size `size` and
// variance of `sigma`. If `sigma` is zero, suitable variance is calculated from size
func GaussianBlur(img *types.GrayImage, size int, sigma float64) *types.FloatImage {
	return Convolve(img, kernels.GaussKernel(size, sigma), CLOSEST)
}

// Applies the Mean filter to image `img` using a kernel of size `size`.
func BoxBlur(img *types.GrayImage, size int) *types.FloatImage {
	return Convolve(img, kernels.BoxKernel(size, size), CLOSEST)
}
