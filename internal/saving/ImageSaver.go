package saver

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"image"
	"image/color"

	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/jijikhal/GoDIP/pkg/types"
)

func Save(filePath string, img *types.ColorImage) error {
	// Open the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	rect := image.Rect(0, 0, img.Width, img.Height)
	rawImage := image.NewRGBA(rect)

	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			r := uint8(img.GetXYC(x, y, 0))
			g := uint8(img.GetXYC(x, y, 1))
			b := uint8(img.GetXYC(x, y, 2))
			a := uint8(0)
			if img.Channels == 4 {
				a = uint8(img.GetXYC(x, y, 3))
			}

			rawImage.Set(x, y, color.RGBA{R: r, G: g, B: b, A: a})
		}
	}

	if strings.HasSuffix(filePath, "png") {
		err = png.Encode(writer, rawImage)
		if err != nil {
			return err
		}
		return nil
	} else if strings.HasSuffix(filePath, "jpg") || strings.HasSuffix(filePath, "jpeg") {
		options := jpeg.Options{Quality: 95}
		err = jpeg.Encode(writer, rawImage, &options)
		if err != nil {
			return err
		}
		return nil
	} else if strings.HasSuffix(filePath, "gif") {
		options := gif.Options{NumColors: 256}
		err = gif.Encode(writer, rawImage, &options)
		if err != nil {
			return err
		}
		return nil
	}

	return fmt.Errorf("unsupported file format")
}

func SaveGray(filePath string, img *types.GrayImage) error {
	return Save(filePath, img.ToColor())
}
