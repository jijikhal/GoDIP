package internal

import (
	"fmt"
	"os"

	"github.com/jijikhal/GoDIP/pkg/types"
)

func SaveAsPPM(filePath string, img *types.Image[uint8]) error {
	// Open the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write the PPM header
	header := fmt.Sprintf("P6\n%d %d\n255\n", img.Width, img.Height)
	_, err = file.Write([]byte(header))
	if err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	// Write the pixel data
	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			_, err := file.Write([]uint8{img.GetXYC(x, y, 0), img.GetXYC(x, y, 1), img.GetXYC(x, y, 2)})
			if err != nil {
				return fmt.Errorf("failed to write pixel data: %w", err)
			}
		}
	}

	return nil
}
