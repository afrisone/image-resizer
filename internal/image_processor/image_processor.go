package image_processor

import (
	"image"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/sunshineplan/imgconv"
)

type imageProcessor interface {
	Resize(srcImage io.Reader, width, height int) (io.Reader, error)
	Process(src io.Reader, width, height int) (io.Reader, error)
}

type storer interface {
	Store(image io.Reader) error
}

func Resize(srcImage image.Image) image.Image {
	resizedImage := imgconv.Resize(srcImage, &imgconv.ResizeOption{Width: 200})
	log.Println("Attempting to write new image....")

	wd, err := os.Getwd()

	if err != nil {
		log.Fatalf("Failed get get working directory: %v\n", err)
	}

	fp := filepath.Join(wd, "output_images", "pic.jpg")
	file, err := os.Create(fp)

	if err != nil {
		log.Fatalf("Failed to create file: %v\n", err)
	}

	if err := imgconv.Write(file, resizedImage, &imgconv.FormatOption{Format: imgconv.JPEG}); err != nil {
		log.Fatalf("Failed to write image: %v\n", err)
	}

	return resizedImage
}
