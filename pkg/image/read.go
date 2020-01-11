package image

import (
	"image"
	"os"

	// Load file formats
	"image/png"
	_ "image/png"
)

func GetImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	image, _, err := image.Decode(f)
	return image, err
}

func SavePNG(img image.Image, filePath string) error {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	return png.Encode(out, img)
}
