package filesystem

import (
	"fmt"
	"image"
	"os"
)

func OpenImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, format, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Detected image type: %s", format)

	return img, nil
}
