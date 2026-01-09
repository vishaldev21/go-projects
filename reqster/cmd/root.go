package cmd

import (
	"fmt"
	"os"
	"reqster/internal/flags"
	"reqster/internal/httpClient"
	"reqster/internal/types"
	"reqster/internal/utils"

	"github.com/spf13/cobra"
)

var method, body, bodyPath string
var query, header, params []string

var rootCmd = &cobra.Command{
	Use:   "reqster",
	Short: "A command-line HTTP request tool with header, query, and parameter support",
	Long: `Reqster is a lightweight command-line HTTP client that allows you to send HTTP requests 
with custom headers, query parameters, and URL parameters. It supports multiple HTTP methods 
and provides a simple interface for testing APIs and interacting with web services from the terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("method and URL is required")
			return
		}
		method := args[0]
		url := args[1]
		headerData := flags.Parser(header, ":")
		queryData := flags.Parser(query, "=")
		paramsData := flags.Parser(params, "=")
		if !utils.MethodValidator(method) {
			fmt.Println("Method not allowed")
			return
		}
		data := types.DataP{
			URL:      url,
			Headers:  headerData,
			Params:   paramsData,
			Query:    queryData,
			Body:     body,
			BodyPath: bodyPath,
		}
		switch method {
		case "get":
			utils.Handler(data, httpClient.GetRequest)
		case "post":
			utils.Handler(data, httpClient.PostRequest)
		case "put":
			utils.Handler(data, httpClient.PutRequest)
		case "delete":
			utils.Handler(data, httpClient.DeleteRequest)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringArrayVarP(&query, "query", "q", []string{}, "takes query string of the url")
	rootCmd.PersistentFlags().StringArrayVarP(&header, "header", "H", []string{}, "takes headers of the request")
	rootCmd.PersistentFlags().StringArrayVarP(&params, "params", "p", []string{}, "takes params of the url")
	rootCmd.PersistentFlags().StringVarP(&body, "body", "b", "", "takes body of the request")
	rootCmd.PersistentFlags().StringVarP(&bodyPath, "body-file", "B", "", "takes body of the request from a file")
}
