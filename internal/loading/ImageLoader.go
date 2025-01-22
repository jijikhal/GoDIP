package loading

import (
	"bufio"
	"os"

	"github.com/jijikhal/GoDIP/pkg/types"

	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func Load(filePath string) (*types.ColorImage, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	rawImage, _, err := image.Decode(reader)

	if err != nil {
		return nil, err
	}

	bounds := rawImage.Bounds()
	image := types.MakeColorImage(bounds.Dy(), bounds.Dx(), 4, 255, 0)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := rawImage.At(x, y).RGBA()

			image.SetPixelXYC(x, y, 0, int(r>>8))
			image.SetPixelXYC(x, y, 1, int(g>>8))
			image.SetPixelXYC(x, y, 2, int(b>>8))
			image.SetPixelXYC(x, y, 3, int(a>>8))
		}
	}

	return image, nil

}
