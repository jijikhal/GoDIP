package types

type Histogram struct {
	MinValue     int
	MaxValue     int
	ColorsPerBin int
	Values       []int
}

func (image *GrayImage) Histogram(colorsPerBin int) Histogram {
	result := Histogram{MinValue: image.MinValue, MaxValue: image.MaxValue, ColorsPerBin: colorsPerBin}
	bins := (image.MaxValue - image.MinValue) / colorsPerBin
	if (image.MaxValue-image.MinValue)%colorsPerBin != 0 {
		bins++
	}
	result.Values = make([]int, bins)

	for i := range image.GetPixelCount() {
		result.Values[image.GetI(i)/colorsPerBin]++
	}

	return result
}
