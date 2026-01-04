/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var query, header, params, method, body string

var rootCmd = &cobra.Command{
	Use:   "Reqster",
	Short: "A brief description of your application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("method and URL is required")
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
	rootCmd.PersistentFlags().StringVarP(&query, "query", "q", "", "takes query string of the url")
	rootCmd.PersistentFlags().StringVarP(&header, "header", "H", "", "takes headers of the request")
	rootCmd.PersistentFlags().StringVarP(&params, "params", "P", "", "takes params of the url")
	rootCmd.PersistentFlags().StringVarP(&body, "body", "b", "", "takes body of the request")
}
