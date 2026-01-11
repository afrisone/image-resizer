package imaging

import (
	"image"
	"log"
	"os"
)

type FileSystemProvider struct{}

func (f FileSystemProvider) GetImage(srcPath string) (image.Image, error) {
	file, err := os.Open(srcPath)
	if err != nil {
		log.Printf("Failed to open file %v\n", err)
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Printf("Failed to decode image &v\n", err)
		return nil, err
	}

	return img, nil
}
