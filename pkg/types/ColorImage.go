package types

import (
	"errors"
	"fmt"
)

type ColorImage struct {
	Width    int
	Height   int
	Channels int
	Data     []int
	MaxValue int
	MinValue int
}

func MakeColorImage(height int, width int, channels int, min int, max int) *ColorImage {
	return &ColorImage{Height: height, Width: width, Channels: channels, Data: make([]int, height*width*channels), MaxValue: max, MinValue: min}
}
func (image *ColorImage) GetPixelI(i int) (int, error) {

	if i >= image.Width*image.Height*image.Channels || i < 0 {
		return 0, errors.New("the cooridnates are out of bounds")
	}
	return image.Data[i], nil
}

func (image *ColorImage) GetI(i int) int {
	return image.Data[i]
}

func (image *ColorImage) GetPixelXYC(x int, y int, c int) (int, error) {

	if x >= image.Width || x < 0 || y >= image.Height || y < 0 || c >= image.Channels || c < 0 {
		return 0, errors.New("the cooridnates are out of bounds")
	}

	return image.Data[y*image.Width*image.Channels+x*image.Channels+c], nil
}

func (image *ColorImage) GetXYC(x int, y int, c int) int {
	return image.Data[y*image.Width*image.Channels+x*image.Channels+c]
}

func (image *ColorImage) SetPixelI(i int, value int) error {

	if i >= image.Width*image.Height*image.Channels || i < 0 {
		return errors.New("the cooridnates are out of bounds")
	}

	if value > image.MaxValue || value < image.MinValue {
		return errors.New("the value is out of allowed range")
	}

	image.Data[i] = value
	return nil
}

func (image *ColorImage) SetI(i int, value int) {
	image.Data[i] = value
}

func (image *ColorImage) SetPixelXYC(x int, y int, c int, value int) error {

	if x >= image.Width || x < 0 || y >= image.Height || y < 0 || c >= image.Channels || c < 0 {
		return errors.New("the cooridnates are out of bounds")
	}

	if value > image.MaxValue || value < image.MinValue {
		return errors.New("the value is out of allowed range")
	}

	image.Data[y*image.Width*image.Channels+x*image.Channels+c] = value

	return nil
}

func (image *ColorImage) SetXYC(x int, y int, c int, value int) {
	image.Data[y*image.Width*image.Channels+x*image.Channels+c] = value
}

func (image *ColorImage) GetPixelCount() int {

	return image.Channels * image.Width * image.Height
}

func (image *ColorImage) GetChannel(channel int) *GrayImage {
	newImage := MakeGrayImage(image.Height, image.Width, image.MinValue, image.MaxValue)
	for y := 0; y < image.Height; y++ {
		for x := 0; x < image.Width; x++ {
			newImage.SetXY(x, y, image.GetXYC(x, y, channel))
		}
	}

	return newImage
}

func MergeChannels(channels ...*GrayImage) (*ColorImage, error) {
	if len(channels) == 0 {
		return nil, errors.New("at least one channel must be provided")
	}

	width := channels[0].Width
	height := channels[0].Height
	maxValue := channels[0].MaxValue
	minValue := channels[0].MinValue

	for i, channel := range channels {
		if channel.Height != height || channel.Width != width || channel.MinValue != minValue || channel.MaxValue != maxValue {
			return nil, fmt.Errorf("channel %d has different format than others", i)
		}
	}

	result := MakeColorImage(height, width, len(channels), minValue, maxValue)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			for c, channel := range channels {
				result.SetXYC(x, y, c, channel.GetXY(x, y))
			}
		}
	}

	return result, nil
}
