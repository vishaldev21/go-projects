package httpClient

import (
	"errors"
	"reqster/internal/types"

	"github.com/go-resty/resty/v2"
)

func PutRequest(data types.DataP) (*string, error) {
	if data.Body == "" {
		return nil, errors.New("Body is not provided")
	}
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
	response, err := client.R().SetBody(data.Body).Put(data.URL)
	if err != nil {
		return nil, err
	}
	byteData := response.Body()
	stringData := string(byteData)
	return &stringData, nil
}
