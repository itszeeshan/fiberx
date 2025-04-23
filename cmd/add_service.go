package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/itszeeshan/fiberx/internal/generator"
	"github.com/spf13/cobra"
)

var (
	serviceDB      string
	serviceMethods []string
	serviceRedis   bool
)

var addServiceCmd = &cobra.Command{
	Use:   "service [name]",
	Short: "Add a new service component",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := strings.ToLower(args[0])

		config := generator.ServiceConfig{
			Name:      serviceName,
			DBType:    serviceDB,
			Methods:   serviceMethods,
			WithRedis: serviceRedis,
		}

		if err := generator.GenerateService(config); err != nil {
			log.Fatalf("Failed to generate service: %v", err)
		}

		fmt.Printf("✅ Service '%s' created with:\n", serviceName)
		fmt.Printf("   Database: %s\n", serviceDB)
		fmt.Printf("   Methods: %v\n", serviceMethods)
		fmt.Printf("   Redis: %v\n", serviceRedis)
	},
}

func init() {
	addCmd.AddCommand(addServiceCmd)

	addServiceCmd.Flags().StringVarP(&serviceDB, "db", "d", "", // Changed default to empty
		"Database type (postgres, mysql, sqlite)")
	addServiceCmd.Flags().StringSliceVarP(&serviceMethods, "methods", "m", []string{}, // Empty default
		"Comma-separated list of methods to generate (crud, create, read, update, delete)")
	addServiceCmd.Flags().BoolVarP(&serviceRedis, "redis", "r", false,
		"Add Redis caching support")
}
