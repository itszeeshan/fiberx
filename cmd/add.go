package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/itszeeshan/fiberx/internal/constants"
	"github.com/itszeeshan/fiberx/internal/generator"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [feature]",
	Short: "Add a feature to your existing Fiber project (e.g. postgres, viper)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		feature := strings.ToLower(args[0])
		validFeatures := constants.ValidFeatures

		if !validFeatures[feature] {
			log.Fatalf("Invalid feature: %s. Valid options: postgres, viper, redis, jwt", feature)
		}

		// Ensure we're in a Go project
		if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
			log.Fatal("No go.mod found in current directory. Are you inside a Fiber project?")
		}

		err := generator.AddFeature(feature)
		if err != nil {
			log.Fatalf("Failed to add feature: %v", err)
		}

		fmt.Printf("✅ Successfully added feature: %s\n", feature)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
