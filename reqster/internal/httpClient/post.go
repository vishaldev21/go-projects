package httpClient

import (
	"reqster/internal/types"

	"github.com/go-resty/resty/v2"
)

func PostRequest(data types.DataP) (*string, error) {
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
	response, err := client.R().SetBody(data.Body).Post(data.URL)
	if err != nil {
		return nil, err
	}
	byteData := response.Body()
	stringData := string(byteData)
	return &stringData, nil
}
