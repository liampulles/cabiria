package image

import (
	"image"
	"os"

	// Allows us to import PNG files.
	_ "image/png"
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
