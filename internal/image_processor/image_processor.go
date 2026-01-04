package image_processor

import (
	"fmt"
	"io"
)

type imageProcessor interface {
	Resize(srcImage io.Reader, width, height int) (io.Reader, error)
	Process(src io.Reader, width, height int) (io.Reader, error)
}

type storer interface {
	Store(image io.Reader) error
}

func Process() {
	fmt.Println("Hello from processor!")
}
