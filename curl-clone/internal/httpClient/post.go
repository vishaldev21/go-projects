package internal

import "github.com/go-resty/resty/v2"

func PostRequest(url string) (*string, error) {
	client := resty.New()
	response, err := client.R().Post(url)
	if err != nil {
		return nil, err
	}
	byteData := response.Body()
	data := string(byteData)
	return &data, nil
}
