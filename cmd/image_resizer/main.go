package main

import (
	"log"
	"os"
	"path/filepath"

	_ "image/jpeg"

	"github.com/afrisone/image_resizer/internal/imaging"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting working directory %v\n", err)
		os.Exit(1)
	}

	img, err := imaging.OpenImage(filepath.Join(wd, "src_images", "pic.jpg"))
	if err != nil {
		log.Printf("Error during import: %v\n", err)
		os.Exit(1)
	}

	if err := imaging.Resize(img); err != nil {
		log.Printf("Failed to resize image %v\n", err)
	}
}
