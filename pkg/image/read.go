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

// GetPNGs can load multiple PNG files from disk into memory
func GetPNGs(filePaths []string) ([]image.Image, error) {
	images := make([]image.Image, len(filePaths))
	for i, filePath := range filePaths {
		img, err := GetPNG(filePath)
		if err != nil {
			return nil, err
		}
		images[i] = img
	}
	return images, nil
}
