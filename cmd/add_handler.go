package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/itszeeshan/fiberx/internal/generator"
	"github.com/spf13/cobra"
)

var handlerMethods []string

var addHandlerCmd = &cobra.Command{
	Use:   "handler [name]",
	Short: "Add a new HTTP handler",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		handlerName := strings.ToLower(args[0])
		validMethods := map[string]bool{
			"GET":    true,
			"POST":   true,
			"PUT":    true,
			"DELETE": true,
			"PATCH":  true,
		}

		// Validate methods
		var methods []string
		for _, m := range handlerMethods {
			method := strings.ToUpper(m)
			if !validMethods[method] {
				log.Fatalf("Invalid method: %s. Valid methods: GET, POST, PUT, DELETE, PATCH", m)
			}
			methods = append(methods, method)
		}

		if len(methods) == 0 {
			methods = []string{"GET"} // Default to GET
		}

		// Generate handler file
		if err := generator.GenerateHandler(handlerName, methods); err != nil {
			log.Fatalf("Failed to generate handler: %v", err)
		}

		fmt.Printf("✅ Handler '%s' created with methods: %v\n", handlerName, methods)
	},
}

func init() {
	addCmd.AddCommand(addHandlerCmd)
	addHandlerCmd.Flags().StringSliceVarP(
		&handlerMethods,
		"methods",
		"m",
		[]string{"GET"},
		"Comma-separated list of HTTP methods",
	)
}
