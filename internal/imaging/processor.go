package imaging

import (
	"image"
	"log"
	"os"
	"path/filepath"

	"github.com/sunshineplan/imgconv"
)

func Resize(srcImage image.Image) error {
	resizedImage := imgconv.Resize(srcImage, &imgconv.ResizeOption{Width: 200})

	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Failed get get working directory: %v\n", err)
		return err
	}

	fp := filepath.Join(wd, "output_images", "pic.jpg")

	file, err := os.Create(fp)
	if err != nil {
		log.Printf("Failed to create file: %v\n", err)
		return err
	}

	log.Println("Writing image to: ", fp)

	if err := imgconv.Write(file, resizedImage, &imgconv.FormatOption{Format: imgconv.JPEG}); err != nil {
		log.Printf("Failed to write image: %v\n", err)
		return err
	}

	return nil
}
