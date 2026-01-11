package imaging

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

type WebScraperImageProvider struct{}

func (f WebScraperImageProvider) GetImage(srcUrl string) (image.Image, error) {
	response, err := http.Get(srcUrl)

	if err != nil {
		log.Printf("Failed fetching web content: %v", err)
		return nil, err
	}
	defer response.Body.Close()

	tokenizer := html.NewTokenizer(response.Body)
	baseUrl, _ := url.Parse(srcUrl)

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			err := tokenizer.Err()
			return nil, err
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "img" {
				for _, attr := range token.Attr {
					if attr.Key == "src" && attr.Val != "" {
						imgUrl, err := url.Parse(attr.Val)
						if err != nil {
							return nil, err
						}
						imgUrlPath := baseUrl.ResolveReference(imgUrl).String()

						return fetchImage(imgUrlPath)
					}
				}
			}
		}
	}
}

func fetchImage(imgUrl string) (image.Image, error) {
	imgResponse, err := http.Get(imgUrl)

	if err != nil {
		return nil, err
	}
	defer imgResponse.Body.Close()

	if imgResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch image: %s", imgUrl)
	}

	img, _, err := image.Decode(imgResponse.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}
