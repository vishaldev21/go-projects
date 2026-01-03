/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
)

var query, header, params string

type Input struct {
	Website string `validate:"required,url"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Reqster",
	Short: "A brief description of your application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("URL is required")
			return
		}
		input := Input{Website: args[0]}
		v := validator.New()
		if err := v.Struct(input); err != nil {
			fmt.Println("Provide valid URL")
			return
		}
		if query != "" {
			fmt.Println(query)
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
}
