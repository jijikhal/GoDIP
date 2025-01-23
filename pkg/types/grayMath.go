package types

import (
	"errors"
)

func (image *GrayImage) Invert() {
	pixelCount := image.GetPixelCount()
	for i := range pixelCount {
		image.SetI(i, image.MaxValue-(image.GetI(i)-image.MinValue))
	}
}

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

func (image *GrayImage) Add(img2 *GrayImage) error {
	if image.Width != img2.Width || image.Height != img2.Height {
		return errors.New("the images have different size")
	}

	for i := range image.GetPixelCount() {
		image.SetI(i, min(max(image.GetI(i)+img2.GetI(i), image.MinValue), image.MaxValue))
	}

	return nil
}

func (image *GrayImage) Subtract(img2 *GrayImage) error {
	if image.Width != img2.Width || image.Height != img2.Height {
		return errors.New("the images have different size")
	}

	for i := range image.GetPixelCount() {
		image.SetI(i, min(max(image.GetI(i)-img2.GetI(i), image.MinValue), image.MaxValue))
	}

	return nil
}

func (image *GrayImage) Multiply(img2 *GrayImage) error {
	if image.Width != img2.Width || image.Height != img2.Height {
		return errors.New("the images have different size")
	}

	for i := range image.GetPixelCount() {
		image.SetI(i, min(max(image.GetI(i)*img2.GetI(i), image.MinValue), image.MaxValue))
	}

	return nil
}
