package generator

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var templateFS embed.FS

type ProjectConfig struct {
	Name     string
	Features map[string]bool
}

func ScaffoldProject(config ProjectConfig) error {

	if err := os.Mkdir(config.Name, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	if err := processTemplates("templates/base", config); err != nil {
		return fmt.Errorf("failed to process base templates: %w", err)
	}

	for feature := range config.Features {
		featurePath := filepath.Join("templates/features", feature)
		if err := processTemplates(featurePath, config); err != nil {
			return fmt.Errorf("failed to process %s templates: %w", feature, err)
		}
	}

	if err := initGoModule(config); err != nil {
		return err
	}

	return nil
}

func processTemplates(templateDir string, cfg ProjectConfig) error {
	return fs.WalkDir(templateFS, templateDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}

		content, err := fs.ReadFile(templateFS, path)
		if err != nil {
			return fmt.Errorf("failed to read template %s: %w", path, err)
		}

		tmpl, err := template.New(path).Parse(string(content))
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", path, err)
		}

		relPath, _ := filepath.Rel(templateDir, path)
		outputPath := filepath.Join(cfg.Name, strings.TrimSuffix(relPath, ".tmpl"))

		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			return fmt.Errorf("failed to create directories for %s: %w", outputPath, err)
		}

		file, err := os.Create(outputPath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", outputPath, err)
		}
		defer file.Close()

		if err := tmpl.Execute(file, cfg); err != nil {
			return fmt.Errorf("failed to execute template %s: %w", path, err)
		}

		return nil
	})
}

func initGoModule(cfg ProjectConfig) error {
	modContent := fmt.Sprintf("module %s\n\ngo 1.22\n", cfg.Name)
	if err := os.WriteFile(filepath.Join(cfg.Name, "go.mod"), []byte(modContent), 0644); err != nil {
		return err
	}

	var deps []string
	deps = append(deps, "github.com/gofiber/fiber/v2@latest")

	for feature := range cfg.Features {
		if pkgs, ok := featureDependencies[feature]; ok {
			deps = append(deps, pkgs...)
		}
	}

	cmd := exec.Command("go", append([]string{"get", "-u"}, deps...)...)
	cmd.Dir = cfg.Name
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("go get failed: %v\n%s", err, output)
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = cfg.Name
	return cmd.Run()
}
