package httpClient

import (
	"github.com/go-resty/resty/v2"
)

type GetData struct {
	URL     string
	Headers map[string]string
	Params  map[string]string
	Query   map[string]string
}

func GetRequest(data GetData) (*string, error) {
	client := resty.New()
	if len(data.Headers) > 0 {
		client = client.SetHeaders(data.Headers)
	}
	if len(data.Query) > 0 {
		client = client.SetQueryParams(data.Query)
	}
	if len(data.Params) > 0 {
		client = client.SetPathParams(data.Params)
	}
	response, err := client.R().Get(data.URL)
	if err != nil {
		return nil, err
	}
	byteData := response.Body()
	stringData := string(byteData)
	return &stringData, nil
}
