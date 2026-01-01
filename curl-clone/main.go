package main

import (
	"curl-clone/internal"
	"fmt"
)

// import "curl-clone/cmd"

func main() {
	// cmd.Execute()
	data, err := internal.GetRequest("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*data)
}
