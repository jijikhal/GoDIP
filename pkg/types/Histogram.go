package types

// Structure for representing histogram
type Histogram struct {
	MinValue     int
	MaxValue     int
	ColorsPerBin int
	Values       []int
}

// Calculates a histogram of a single channel
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
