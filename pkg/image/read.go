package image

import (
	"image"
	"os"

	"image/png"
)

// GetPNG is able to load a PNG file from the disk into memory.
func GetPNG(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	image, _, err := image.Decode(f)
	return image, err
}

// SavePNG encodes the image to the desired png file.
func SavePNG(img image.Image, filePath string) error {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	return png.Encode(out, img)
}
