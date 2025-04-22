package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fiberx",
	Short: "FiberX - CLI for speeding up GoFiber development without magic",
	Long: `FiberX is a CLI tool that generates idiomatic Go code 
for Fiber projects without enforcing rigid structures or frameworks.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("FiberX CLI ⚡️ — use 'fiberx help' to get started")
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add components to your project",
}

func init() {
	rootCmd.AddCommand(addCmd)
}

// Execute is the main entry point for the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
