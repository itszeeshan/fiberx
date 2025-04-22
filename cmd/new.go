package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var features []string

var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Scaffold a new GoFiber project",
	Args:  cobra.ExactArgs(1),
	Run:   createProject,
}

func init() {
	newCmd.Flags().StringSliceVarP(&features, "with", "w", []string{}, "Optional features (e.g. postgres, redis)")
	rootCmd.AddCommand(newCmd)
}

func createProject(cmd *cobra.Command, args []string) {
	name := args[0]
	dirs := []string{
		"cmd",
		"handlers",
		"middlewares",
		"services",
		"config",
	}

	fmt.Println("🔧 Creating project:", name)

	if err := os.Mkdir(name, 0755); err != nil {
		log.Fatalf("Failed to create project folder: %v", err)
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(name, dir), 0755); err != nil {
			log.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}

	createMain(name)
	createGoMod(name)
	createGitignore(name)
	createReadme(name)

	fmt.Println("✅ Project scaffolded successfully!")
}

func createMain(name string) {
	const mainTpl = `package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// TODO: Register routes

	log.Fatal(app.Listen(":3000"))
}
`
	writeFile(filepath.Join(name, "cmd", "main.go"), mainTpl)
}

func createGoMod(name string) {
	modTpl := fmt.Sprintf("module %s\n\ngo 1.21\n\nrequire github.com/gofiber/fiber/v2 v2.50.0\n", name)
	writeFile(filepath.Join(name, "go.mod"), modTpl)
}

func createGitignore(name string) {
	writeFile(filepath.Join(name, ".gitignore"), "bin/\n.env\n*.log\n")
}

func createReadme(name string) {
	content := fmt.Sprintf("# %s\n\nGenerated with FiberX ⚡", strings.Title(name))
	writeFile(filepath.Join(name, "README.md"), content)
}

func writeFile(path, content string) {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		log.Fatalf("Failed to write file %s: %v", path, err)
	}
}
