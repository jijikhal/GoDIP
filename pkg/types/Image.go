package types

import "errors"

type UnsignedInteger interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type SignedInteger interface {
	int | int8 | int16 | int32 | int64
}

type Integer interface {
	SignedInteger | UnsignedInteger
}

type Float interface {
	float32 | float64
}

type Number interface {
	Integer | Float
}

type Image[T Number] struct {
	Width    int
	Height   int
	Channels int
	Data     []T
}

func MakeImageMultiChannel[T Number](height int, width int, channels int) *Image[T] {
	return &Image[T]{Height: height, Width: width, Channels: channels, Data: make([]T, height*width*channels)}
}

func MakeImageSingleChannel[T Number](height int, width int) *Image[T] {
	return &Image[T]{Height: height, Width: width, Channels: 1, Data: make([]T, height*width)}
}

func (image *Image[T]) GetPixelI(i int) (T, error) {

	if i >= image.Width*image.Height*image.Channels || i < 0 {
		return 0, errors.New("the cooridnates are out of bounds")
	}

	return image.Data[i], nil
}

func (image *Image[T]) GetPixelXY(x int, y int) (T, error) {
	if image.Channels != 1 {
		return 0, errors.New("the Get(x, y) method can only be used on single channel images")
	}

	if x >= image.Width || x < 0 || y >= image.Height || y < 0 {
		return 0, errors.New("the cooridnates are out of bounds")
	}

	return image.Data[y*image.Width+x], nil
}

func (image *Image[T]) GetPixelXYC(x int, y int, c int) (T, error) {

	if x >= image.Width || x < 0 || y >= image.Height || y < 0 || c >= image.Channels || c < 0 {
		return 0, errors.New("the cooridnates are out of bounds")
	}

	return image.Data[y*image.Width*image.Channels+x*image.Channels+c], nil
}

func (image *Image[T]) GetXYC(x int, y int, c int) T {
	return image.Data[y*image.Width*image.Channels+x*image.Channels+c]
}

func (image *Image[T]) SetPixelI(i int, value T) error {

	if i >= image.Width*image.Height*image.Channels || i < 0 {
		return errors.New("the cooridnates are out of bounds")
	}

	image.Data[i] = value
	return nil
}

func (image *Image[T]) SetPixelXY(x int, y int, value T) error {
	if image.Channels != 1 {
		return errors.New("the Get(x, y) method can only be used on single channel images")
	}

	if x >= image.Width || x < 0 || y >= image.Height || y < 0 {
		return errors.New("the cooridnates are out of bounds")
	}

	image.Data[y*image.Width+x] = value
	return nil
}

func (image *Image[T]) SetPixelXYC(x int, y int, c int, value T) error {

	if x >= image.Width || x < 0 || y >= image.Height || y < 0 || c >= image.Channels || c < 0 {
		return errors.New("the cooridnates are out of bounds")
	}

	image.Data[y*image.Width*image.Channels+x*image.Channels+c] = value

	return nil
}

func (image *Image[T]) GetPixelCount() int {

	return image.Channels * image.Width * image.Height
}
