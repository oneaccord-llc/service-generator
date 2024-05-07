package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	gen "github.com/oneaccord-llc/service-generator/generator"
)

//go:embed templates/*
var templates embed.FS

var rootCmd = &cobra.Command{
	Use:   "service-generator",
	Short: "A CLI tool to generate a service for oneaccord",
	Run: func(cmd *cobra.Command, args []string) {
		projectName := ""
		fmt.Print("Enter your service name: ")
		fmt.Scanln(&projectName)

		generator := gen.NewGenerator(projectName)
		if err := generator.Generate(templates); err != nil {
			fmt.Println("Error generating service:", err)
			os.Exit(1)
		}
		fmt.Println("Service generated successfully!")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
