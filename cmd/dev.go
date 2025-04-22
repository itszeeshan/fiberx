package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Start development server with hot-reload",
	Run: func(cmd *cobra.Command, args []string) {
		if !isAirInstalled() {
			fmt.Println("🚨 Air is not installed. Installing...")
			installAir()
		}

		if !airConfigExists() {
			createAirConfig()
		}

		startDevServer()
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}

func isAirInstalled() bool {
	_, err := exec.LookPath("air")
	return err == nil
}

func installAir() {
	cmd := exec.Command("go", "install", "github.com/cosmtrek/air@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal("❌ Failed to install Air. Please install manually: go install github.com/cosmtrek/air@latest")
	}
}

func airConfigExists() bool {
	_, err := os.Stat(".air.toml")
	return err == nil
}

func createAirConfig() {
	const airConfig = `# .air.toml
[build]
  cmd = "go build -o ./tmp/main ./cmd/main.go"
  bin = "tmp/main"
  include_ext = ["go", "tpl", "tmpl", "html"]
  exclude_dir = ["assets", "tmp", "vendor", "frontend"]
  include_dir = ["."]
  exclude_regex = ["_test.go"]
  exclude_unchanged = true
  follow_symlink = false
  log = "build-errors.log"
  delay = 1000
  stop_on_error = true
  send_interrupt = false
  kill_delay = 500

[color]
  main = "magenta"
  watcher = "cyan"
  build = "yellow"
  runner = "green"

[log]
  time = false
`

	if err := os.WriteFile(".air.toml", []byte(airConfig), 0644); err != nil {
		log.Fatal("❌ Failed to create .air.toml config file")
	}
	fmt.Println("✅ Created .air.toml configuration file")
}

func startDevServer() {
	fmt.Println("🚀 Starting development server...")
	fmt.Println("🔄 File changes will trigger hot-reload")
	fmt.Println("🛑 Press Ctrl+C to stop")

	cmd := exec.Command("air")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		log.Fatal("❌ Failed to start development server: ", err)
	}
}
