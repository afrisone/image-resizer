package imaging

import "image"

type ImageProvider interface {
	GetImage(source string) (image.Image, error)
}
