package types

import (
	"errors"
)

// Inverts all pixels
func (image *GrayImage) Invert() {
	pixelCount := image.GetPixelCount()
	for i := range pixelCount {
		image.SetI(i, image.MaxValue-(image.GetI(i)-image.MinValue))
	}
}

// Applies absolute value to all pixels
func (image *GrayImage) Abs() {
	pixelCount := image.GetPixelCount()
	for i := range pixelCount {
		val := image.GetI(i)
		if val < 0 {
			val = val * (-1)
		}
		image.SetI(i, val)
	}
}

// Adds two images of same size
func (image *GrayImage) Add(img2 *GrayImage) error {
	if image.Width != img2.Width || image.Height != img2.Height {
		return errors.New("the images have different size")
	}

	for i := range image.GetPixelCount() {
		image.SetI(i, min(max(image.GetI(i)+img2.GetI(i), image.MinValue), image.MaxValue))
	}

	return nil
}

// Subtract image from this image (both must be the same size)
func (image *GrayImage) Subtract(img2 *GrayImage) error {
	if image.Width != img2.Width || image.Height != img2.Height {
		return errors.New("the images have different size")
	}

	for i := range image.GetPixelCount() {
		image.SetI(i, min(max(image.GetI(i)-img2.GetI(i), image.MinValue), image.MaxValue))
	}

	return nil
}

// Multiplies two images of same size
func (image *GrayImage) Multiply(img2 *GrayImage) error {
	if image.Width != img2.Width || image.Height != img2.Height {
		return errors.New("the images have different size")
	}

	for i := range image.GetPixelCount() {
		image.SetI(i, min(max(image.GetI(i)*img2.GetI(i), image.MinValue), image.MaxValue))
	}

	return nil
}
