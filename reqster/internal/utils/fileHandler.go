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

func WriteFile(filePath string, data []byte) (bool, error) {
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return false, err
	}
	return true, nil
}
