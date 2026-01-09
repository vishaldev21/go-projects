package httpClient

import (
	"errors"
	"reqster/internal/types"
	"reqster/internal/utils"

	"github.com/go-resty/resty/v2"
)

func PostRequest(data types.DataP) (*string, error) {
	if data.Body == "" && data.BodyPath == "" {
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
	if data.BodyPath != "" {
		body, err := utils.ReadFile(data.BodyPath)
		if err != nil {
			return nil, err
		}
		data.Body = *body
	}
	response, err := client.R().SetBody(data.Body).Post(data.URL)
	if err != nil {
		return nil, err
	}
	byteData := response.Body()
	stringData := string(byteData)
	return &stringData, nil
}
