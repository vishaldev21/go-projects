package internal

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

type GetData struct {
	url     string
	headers map[string]string
	params  map[string]string
	query   string
}

func GetRequest(data GetData) (*string, error) {
	client := resty.New()
	if data.headers != nil {
		client = client.SetHeaders(data.headers)
	}
	response, err := client.R().Get(data.url)
	if err != nil {
		return nil, err
	}
	byteData := response.Body()
	stringData := string(byteData)
	return &stringData, nil
}

func GetQueryRequest(data GetData) (*string, error) {
	if data.query == "" {
		return nil, errors.New("Query string is not found")
	}
	client := resty.New()
	if data.headers != nil {
		client = client.SetHeaders(data.headers)
	}
	response, err := client.R().SetQueryString(data.query).Get(data.url)
	if err != nil {
		return nil, err
	}
	byteData := response.Body()
	stringData := string(byteData)
	return &stringData, err
}

func GetParamRequest(data GetData) (*string, error) {
	if data.params == nil {
		return nil, errors.New("Path params is not found")
	}
	client := resty.New()
	if data.headers != nil {
		client = client.SetHeaders(data.headers)
	}
	response, err := client.R().SetPathParams(data.params).Get(data.url)
	if err != nil {
		return nil, err
	}
	byteData := response.Body()
	stringData := string(byteData)
	return &stringData, err
}
