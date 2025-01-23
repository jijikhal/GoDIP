package types

import "errors"

type GrayImage struct {
	Width    int
	Height   int
	Data     []int
	MaxValue int
	MinValue int
}

func MakeGrayImage(height int, width int, min int, max int) *GrayImage {
	return &GrayImage{Height: height, Width: width, Data: make([]int, height*width), MaxValue: max, MinValue: min}
}

func (image *GrayImage) GetPixelI(i int) (int, error) {

	if i >= image.Width*image.Height || i < 0 {
		return 0, errors.New("the cooridnates are out of bounds")
	}
	return image.Data[i], nil
}

func (image *GrayImage) GetI(i int) int {
	return image.Data[i]
}

func (image *GrayImage) GetPixelXY(x int, y int) (int, error) {

	if x >= image.Width || x < 0 || y >= image.Height || y < 0 {
		return 0, errors.New("the cooridnates are out of bounds")
	}

	return image.Data[y*image.Width+x], nil
}

func (image *GrayImage) GetXY(x int, y int) int {
	return image.Data[y*image.Width+x]
}

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

func (image *GrayImage) SetI(i int, value int) {
	image.Data[i] = value
}

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

func (image *GrayImage) SetXY(x int, y int, value int) {
	image.Data[y*image.Width+x] = value
}

func (image *GrayImage) GetPixelCount() int {

	return image.Width * image.Height
}

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

func (image *GrayImage) Duplicate() *GrayImage {
	result := MakeGrayImage(image.Height, image.Width, image.MinValue, image.MaxValue)
	pixelCount := image.GetPixelCount()

	for i := 0; i < pixelCount; i++ {
		result.SetI(i, image.GetI(i))
	}

	return result
}
