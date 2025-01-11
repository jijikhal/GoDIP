package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jijikhal/GoDIP/pkg/types"
)

func LoadPPM(filePath string) (*types.Image[uint8], error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	encoding, _ := reader.ReadString('\n')
	encoding = strings.TrimRight(encoding, "\n")

	var size string
	for size, _ = reader.ReadString('\n'); strings.HasPrefix(size, "#"); size, _ = reader.ReadString('\n') {
	}
	dimensions := strings.Fields(strings.TrimRight(size, "\n"))
	width, _ := strconv.Atoi(dimensions[0])
	height, _ := strconv.Atoi(dimensions[1])

	image := types.MakeImageMultiChannel[uint8](height, width, 3)

	maxValString, _ := reader.ReadString('\n')
	maxVal, _ := strconv.Atoi(strings.TrimRight(maxValString, "\n"))

	if maxVal > 255 {
		return nil, errors.New("too large max value")
	}

	if encoding == "P3" {
		fmt.Println("P3")
		return loadPPMText(image, reader)
	} else if encoding == "P6" {
		fmt.Println("P6")
		return loadPPMBinary(image, reader)
	} else {
		return nil, fmt.Errorf("unsupported format: only P6 and P3 is supported")
	}
}

func loadPPMText(image *types.Image[uint8], reader *bufio.Reader) (*types.Image[uint8], error) {

	row, col, ch := 0, 0, 0

	for {
		line, err := reader.ReadString('\n')

		values := strings.Fields(strings.TrimRight(line, "\n"))
		for _, v := range values {
			val, _ := strconv.Atoi(v)
			fmt.Println(v, row, col, ch)
			image.SetPixelXYC(col, row, ch, uint8(val))
			ch++

			if ch == 3 {
				col++
				ch = 0
			}

			if col == (image.Width) {
				row++
				col = 0
			}
		}

		if err != nil {
			// surely it is EOF
			break
		}
	}

	return image, nil
}

func loadPPMBinary(image *types.Image[uint8], reader *bufio.Reader) (*types.Image[uint8], error) {

	pixelData := make([]byte, 3*image.Width*image.Height)
	buffer := make([]byte, 256)
	read := 0

	// I hate buffering omg
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			return nil, err
		}
		for i := range n {
			if read+i >= 3*image.Width*image.Height {
				break
			}
			pixelData[read+i] = buffer[i]
		}
		read += n
		if read == 3*image.Width*image.Height {
			break
		}
		if n == 0 {
			return nil, errors.New("not enough data")
		}
	}

	fmt.Println(pixelData)

	idx := 0
	for y := 0; y < image.Height; y++ {
		for x := 0; x < image.Width; x++ {
			image.SetPixelXYC(x, y, 0, pixelData[idx])
			image.SetPixelXYC(x, y, 1, pixelData[idx+1])
			image.SetPixelXYC(x, y, 2, pixelData[idx+2])
			idx += 3
		}
	}

	return image, nil
}
