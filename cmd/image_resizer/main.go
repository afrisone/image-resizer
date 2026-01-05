package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "image/jpeg"

	"github.com/afrisone/image_resizer/internal/filesystem"
	"github.com/afrisone/image_resizer/internal/image_processor"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	img, err := filesystem.OpenImage(filepath.Join(wd, "src_images", "pic.jpg"))

	if err != nil {
		fmt.Printf("Error during import: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Image dimensions: %d x %d\n", img.Bounds().Dx(), img.Bounds().Dy())
	resizedImage := image_processor.Resize(img)

	fmt.Printf("Image dimensions: %d x %d\n", resizedImage.Bounds().Dx(), resizedImage.Bounds().Dy())
}
