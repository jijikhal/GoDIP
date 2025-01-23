package types

import (
	"errors"
	"math"
)

type FloatImage struct {
	Width    int
	Height   int
	Data     []float64
	MaxValue float64
	MinValue float64
}

func MakeFloatImage(height int, width int, min float64, max float64) *FloatImage {
	return &FloatImage{Height: height, Width: width, Data: make([]float64, height*width), MaxValue: max, MinValue: min}
}

func (image *FloatImage) GetPixelI(i int) (float64, error) {

	if i >= image.Width*image.Height || i < 0 {
		return 0, errors.New("the cooridnates are out of bounds")
	}
	return image.Data[i], nil
}

func (image *FloatImage) GetI(i int) float64 {
	return image.Data[i]
}

func (image *FloatImage) GetPixelXY(x int, y int) (float64, error) {

	if x >= image.Width || x < 0 || y >= image.Height || y < 0 {
		return 0, errors.New("the cooridnates are out of bounds")
	}

	return image.Data[y*image.Width+x], nil
}

func (image *FloatImage) GetXY(x int, y int) float64 {
	return image.Data[y*image.Width+x]
}

func (image *FloatImage) SetPixelI(i int, value float64) error {

	if i >= image.Width*image.Height || i < 0 {
		return errors.New("the cooridnates are out of bounds")
	}

	if value-image.MaxValue > 0.0001 || value-image.MinValue < -0.00001 {
		return errors.New("the value is out of allowed range")
	}

	image.Data[i] = value
	return nil
}

func (image *FloatImage) SetI(i int, value float64) {
	image.Data[i] = value
}

func (image *FloatImage) SetPixelXY(x int, y int, value float64) error {

	if x >= image.Width || x < 0 || y >= image.Height || y < 0 {
		return errors.New("the cooridnates are out of bounds")
	}

	if value-image.MaxValue > 0.0001 || value-image.MinValue < -0.00001 {
		return errors.New("the value is out of allowed range")
	}

	image.Data[y*image.Width+x] = value

	return nil
}

func (image *FloatImage) SetXY(x int, y int, value float64) {
	image.Data[y*image.Width+x] = value
}

func (image *FloatImage) GetPixelCount() int {

	return image.Width * image.Height
}

func (image *FloatImage) ToGray(minVal int, maxVal int) *GrayImage {
	newImage := MakeGrayImage(image.Height, image.Width, minVal, maxVal)
	pixelCount := image.GetPixelCount()
	for i := 0; i < pixelCount; i++ {
		newImage.SetI(i, max(minVal, min(maxVal, int(math.Round(image.GetI(i))))))
	}

	return newImage
}
