package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/itszeeshan/fiberx/internal/generator"
	"github.com/spf13/cobra"
)

var withFeatures string

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new Fiber project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		features := strings.Split(withFeatures, ",")

		// Validate features
		validFeatures := map[string]bool{
			"postgres": true,
			"viper":    true,
			"redis":    true,
			"jwt":      true,
		}

		featureMap := make(map[string]bool)
		for _, f := range features {
			f = strings.TrimSpace(f)
			if f == "" {
				continue
			}
			if !validFeatures[f] {
				log.Fatalf("Invalid feature: %s. Valid options: postgres,viper,redis,jwt", f)
			}
			featureMap[f] = true
		}

		config := generator.ProjectConfig{
			Name:     projectName,
			Features: featureMap,
		}

		if err := generator.ScaffoldProject(config); err != nil {
			log.Fatalf("Error creating project: %v", err)
		}

		fmt.Printf("✅ Successfully created project: %s\n", projectName)
		fmt.Printf("Features enabled: %v\n", features)
	},
}

func init() {
	newCmd.Flags().StringVarP(&withFeatures, "with", "w", "",
		"Comma-separated list of features to enable (postgres,viper,redis,jwt)")
	rootCmd.AddCommand(newCmd)
}
