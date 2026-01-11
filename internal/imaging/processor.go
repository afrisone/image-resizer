package imaging

import (
	"image"
	"log"
	"os"
	"path/filepath"

	"github.com/sunshineplan/imgconv"
)

func ProcessImage(p ImageProvider, src string, outputPath string) error {
	data, err := p.GetImage(src)

	if err != nil {
		log.Printf("Failed to fetch image from %s: %v\n", src, err)
		return err
	}

	log.Println("Successfully fetched image! Processing...")

	img := resize(data)

	return store(img, outputPath)
}

func resize(srcImage image.Image) image.Image {
	return imgconv.Resize(srcImage, &imgconv.ResizeOption{Width: 200})

}

func store(img image.Image, outputPath string) error {
	fp := filepath.Join(outputPath, "pic.jpg")

	file, err := os.Create(fp)
	if err != nil {
		log.Printf("Failed to create file: %v\n", err)
		return err
	}

	log.Println("Writing image to: ", fp)

	if err := imgconv.Write(file, img, &imgconv.FormatOption{Format: imgconv.JPEG}); err != nil {
		log.Printf("Failed to write image: %v\n", err)
		return err
	}

	return nil
}
