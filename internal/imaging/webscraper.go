package imaging

import "image"

type WebScraperImageProvider struct{}

func (f WebScraperImageProvider) GetImage(url string) (image.Image, error) {

	return nil, nil
}
