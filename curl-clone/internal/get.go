package internal

import (
	"github.com/go-resty/resty/v2"
)

func GetRequest(url string) (*string, error) {
	client := resty.New()
	response, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}
	byteData := response.Body()
	stringData := string(byteData)
	return &stringData, nil
}
