package main

import (
	"flag"
	"log"

	_ "image/jpeg"

	"github.com/afrisone/image_resizer/internal/imaging"
)

func main() {
	mode := flag.String("mode", "file", "the mode of image retrieval - ex: 'file' or 'url'")
	src := flag.String("src", "", "the path to the images - ex: '~/Desktop/' or 'https://yoururlhere'")
	out := flag.String("out", "output_images", "the path to write the resized images")
	width := flag.Int("w", 200, "the width of the new image")

	flag.Parse()

	var provider imaging.ImageProvider

	switch *mode {
	case "url":
		provider = imaging.WebScraperImageProvider{}
	case "file":
		provider = imaging.FileSystemProvider{}
	default:
		log.Fatal("unknown mode")
	}

	imaging.ProcessImage(provider, *src, *out, *width)

}
