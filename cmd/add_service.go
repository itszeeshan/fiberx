package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/itszeeshan/fiberx/internal/generator"
	"github.com/spf13/cobra"
)

var addServiceCmd = &cobra.Command{
	Use:   "service [name]",
	Short: "Add a new service layer component",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := strings.ToLower(args[0])
		if err := generator.GenerateService(serviceName); err != nil {
			log.Fatalf("Failed to generate service: %v", err)
		}
		fmt.Printf("✅ Service '%s' created\n", serviceName)
	},
}

func init() {
	addCmd.AddCommand(addServiceCmd)
}
