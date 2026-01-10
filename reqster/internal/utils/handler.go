package utils

import (
	"fmt"
	"reqster/internal/types"
)

func Handler(data types.DataP, requestHandler types.Method) {
	response, err := requestHandler(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	if data.OutputPath != "" {
		stringResponse := []byte(*response)
		isWritten, writeErr := WriteFile(data.OutputPath, stringResponse)
		if !isWritten {
			fmt.Println(writeErr)
			return
		}
		fmt.Printf("Response saved in file: %s\n", data.OutputPath)
		return
	}
	fmt.Println(*response)
}
