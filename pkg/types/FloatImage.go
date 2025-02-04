package types

import (
	"errors"
	"math"
)

// Structure for representing single-channel image of floats
type FloatImage struct {
	Width    int
	Height   int
	Data     []float64
	MaxValue float64
	MinValue float64
}

// Creates a single-channel image of specified size and specified allowed bit depth
func MakeFloatImage(height int, width int, min float64, max float64) *FloatImage {
	return &FloatImage{Height: height, Width: width, Data: make([]float64, height*width), MaxValue: max, MinValue: min}
}

// Safe method for getting i-th pixel of image
// Checks for invalid access. If you know what you are doing use `GetI` instead.
func (image *FloatImage) GetPixelI(i int) (float64, error) {

	if i >= image.Width*image.Height || i < 0 {
		return 0, errors.New("the cooridnates are out of bounds")
	}
	return image.Data[i], nil
}

// Gets value of i-th pixel
func (image *FloatImage) GetI(i int) float64 {
	return image.Data[i]
}

// Safe method for getting pixel at coordinates x, y
// Checks for invalid access. If you know what you are doing use `GetXY` instead.
func (image *FloatImage) GetPixelXY(x int, y int) (float64, error) {

	if x >= image.Width || x < 0 || y >= image.Height || y < 0 {
		return 0, errors.New("the cooridnates are out of bounds")
	}

	return image.Data[y*image.Width+x], nil
}

// Gets value of pixel at coordinates x, y
func (image *FloatImage) GetXY(x int, y int) float64 {
	return image.Data[y*image.Width+x]
}

// Safe method for setting i-th pixel of image
// Checks for invalid access. If you know what you are doing use `SetI` instead.
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

// Sets value of i-th pixel
func (image *FloatImage) SetI(i int, value float64) {
	image.Data[i] = value
}

// Safe method for setting pixel at coordinates x, y
// Checks for invalid access. If you know what you are doing use `SetXY` instead.
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

// Sets value of pixel at coordinates x, y
func (image *FloatImage) SetXY(x int, y int, value float64) {
	image.Data[y*image.Width+x] = value
}

// Returns total ammount of pixels in image
func (image *FloatImage) GetPixelCount() int {

	return image.Width * image.Height
}

// Converts this float image to a int image
func (image *FloatImage) ToGray(minVal int, maxVal int) *GrayImage {
	newImage := MakeGrayImage(image.Height, image.Width, minVal, maxVal)
	pixelCount := image.GetPixelCount()
	for i := 0; i < pixelCount; i++ {
		newImage.SetI(i, max(minVal, min(maxVal, int(math.Round(image.GetI(i))))))
	}

	return newImage
}
