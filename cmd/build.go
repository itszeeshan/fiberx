package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	buildOS      string
	buildArch    string
	buildVersion string
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build production-ready binary",
	Run: func(cmd *cobra.Command, args []string) {
		ensureDistDirectory()
		cleanDependencies()
		buildProductionBinary()
	},
}

func init() {
	buildCmd.Flags().StringVarP(&buildOS, "os", "", runtime.GOOS,
		"Target OS (linux, darwin, windows)")
	buildCmd.Flags().StringVarP(&buildArch, "arch", "", runtime.GOARCH,
		"Target architecture (amd64, arm64)")
	buildCmd.Flags().StringVarP(&buildVersion, "version", "", "0.1.0",
		"Application version")
	buildCmd.Flags().Bool("docker", false, "Build Docker image")
	buildCmd.Flags().Bool("compress", false, "Compress binary with upx")
	buildCmd.Flags().String("output", "", "Custom output name")
	rootCmd.AddCommand(buildCmd)
}

func ensureDistDirectory() {
	distDir := filepath.Join("dist", fmt.Sprintf("%s-%s", buildOS, buildArch))
	if err := os.MkdirAll(distDir, 0755); err != nil {
		log.Fatalf("Failed to create dist directory: %v", err)
	}
}

func cleanDependencies() {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Dependency cleanup failed: %v", err)
	}
}

func buildProductionBinary() {
	ldFlags := []string{
		"-s", // Omit symbol table
		"-w", // Omit DWARF debug info
		fmt.Sprintf("-X main.version=%s", buildVersion),
		fmt.Sprintf("-X main.buildTime=%s", time.Now().Format(time.RFC3339)),
	}

	outputPath := filepath.Join("dist", fmt.Sprintf("%s-%s", buildOS, buildArch),
		getBinaryName())

	cmd := exec.Command("go", "build",
		"-trimpath",
		"-ldflags", strings.Join(ldFlags, " "),
		"-o", outputPath,
		"./cmd/main.go",
	)

	cmd.Env = append(os.Environ(),
		fmt.Sprintf("GOOS=%s", buildOS),
		fmt.Sprintf("GOARCH=%s", buildArch),
		"CGO_ENABLED=0",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("🔨 Building for %s/%s...\n", buildOS, buildArch)
	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ Build failed: %v", err)
	}

	fmt.Printf("✅ Successfully built: %s\n", outputPath)
}

func getBinaryName() string {
	if buildOS == "windows" {
		return "app.exe"
	}
	return "app"
}
