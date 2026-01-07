package utils

import (
	"fmt"
	"reqster/internal/types"
)

func Handler(data types.DataP, requestHandler types.Method) {
	response, err := requestHandler(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*response)
}
