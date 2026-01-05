package filesystem

import (
	"image"
	"os"
)

func OpenImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}
