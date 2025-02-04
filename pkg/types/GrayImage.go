package types

import "errors"

// Structure for representing single-channel image of whole numbers
type GrayImage struct {
	Width    int
	Height   int
	Data     []int
	MaxValue int
	MinValue int
}

// Creates a single-channel image of specified size and specified allowed bit depth
func MakeGrayImage(height int, width int, min int, max int) *GrayImage {
	return &GrayImage{Height: height, Width: width, Data: make([]int, height*width), MaxValue: max, MinValue: min}
}

// Safe method for getting i-th pixel of image
// Checks for invalid access. If you know what you are doing use `GetI` instead.
func (image *GrayImage) GetPixelI(i int) (int, error) {

	if i >= image.Width*image.Height || i < 0 {
		return 0, errors.New("the cooridnates are out of bounds")
	}
	return image.Data[i], nil
}

// Gets value of i-th pixel
func (image *GrayImage) GetI(i int) int {
	return image.Data[i]
}

// Safe method for getting pixel at coordinates x, y
// Checks for invalid access. If you know what you are doing use `GetXY` instead.
func (image *GrayImage) GetPixelXY(x int, y int) (int, error) {

	if x >= image.Width || x < 0 || y >= image.Height || y < 0 {
		return 0, errors.New("the cooridnates are out of bounds")
	}

	return image.Data[y*image.Width+x], nil
}

// Gets value of pixel at coordinates x, y
func (image *GrayImage) GetXY(x int, y int) int {
	return image.Data[y*image.Width+x]
}

// Safe method for setting i-th pixel of image
// Checks for invalid access. If you know what you are doing use `SetI` instead.
func (image *GrayImage) SetPixelI(i int, value int) error {

	if i >= image.Width*image.Height || i < 0 {
		return errors.New("the cooridnates are out of bounds")
	}

	if value > image.MaxValue || value < image.MinValue {
		return errors.New("the value is out of allowed range")
	}

	image.Data[i] = value
	return nil
}

// Sets value of i-th pixel
func (image *GrayImage) SetI(i int, value int) {
	image.Data[i] = value
}

// Safe method for setting pixel at coordinates x, y
// Checks for invalid access. If you know what you are doing use `SetXY` instead.
func (image *GrayImage) SetPixelXY(x int, y int, value int) error {

	if x >= image.Width || x < 0 || y >= image.Height || y < 0 {
		return errors.New("the cooridnates are out of bounds")
	}

	if value > image.MaxValue || value < image.MinValue {
		return errors.New("the value is out of allowed range")
	}

	image.Data[y*image.Width+x] = value

	return nil
}

// Sets value of pixel at coordinates x, y
func (image *GrayImage) SetXY(x int, y int, value int) {
	image.Data[y*image.Width+x] = value
}

// Returns total ammount of pixels in image
func (image *GrayImage) GetPixelCount() int {

	return image.Width * image.Height
}

// Creates a three channel image from a single channel (useful for saving graysacle images)
func (image *GrayImage) ToColor() *ColorImage {
	newImage := MakeColorImage(image.Height, image.Width, 3, image.MinValue, image.MaxValue)
	for y := 0; y < image.Height; y++ {
		for x := 0; x < image.Width; x++ {
			newImage.SetXYC(x, y, 0, image.GetXY(x, y))
			newImage.SetXYC(x, y, 1, image.GetXY(x, y))
			newImage.SetXYC(x, y, 2, image.GetXY(x, y))
		}
	}

	return newImage
}

// Creates a duplicate of the image
func (image *GrayImage) Duplicate() *GrayImage {
	result := MakeGrayImage(image.Height, image.Width, image.MinValue, image.MaxValue)
	pixelCount := image.GetPixelCount()

	for i := 0; i < pixelCount; i++ {
		result.SetI(i, image.GetI(i))
	}

	return result
}
