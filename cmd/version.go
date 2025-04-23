package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// Version information set during build
var (
	Version    = "n/a"
	CommitHash = "n/a"
	BuildDate  = "n/a"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show FiberX version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`FiberX - Production Ready Scaffolding for GoFiber
Version:     %s
Go version:  %s
Git commit:  %s
Build date:  %s
OS/Arch:     %s/%s
`,
			Version,
			runtime.Version(),
			CommitHash,
			BuildDate,
			runtime.GOOS,
			runtime.GOARCH,
		)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
