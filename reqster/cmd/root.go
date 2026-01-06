package cmd

import (
	"fmt"
	"os"
	"reqster/internal/flags"
	"reqster/internal/httpClient"
	"reqster/internal/utils"

	"github.com/spf13/cobra"
)

var method, body string
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
		headerData := flags.HeaderParser(header)
		queryData := flags.ParamParser(query)
		paramsData := flags.ParamParser(params)
		if !utils.MethodValidator(method) {
			fmt.Println("Method not allowed")
			return
		}
		if method == "get" {
			data := httpClient.GetData{
				URL:     url,
				Headers: headerData,
				Params:  paramsData,
				Query:   queryData,
			}
			response, err := httpClient.GetRequest(data)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(*response)
			return
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
	rootCmd.PersistentFlags().StringArrayVarP(&params, "params", "P", []string{}, "takes params of the url")
	rootCmd.PersistentFlags().StringVarP(&body, "body", "b", "", "takes body of the request")
}
