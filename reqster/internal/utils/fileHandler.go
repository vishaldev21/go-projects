package utils

import "os"

func ReadFile(filePath string) (*string, error) {
	byteData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	data := string(byteData)
	return &data, nil
}
