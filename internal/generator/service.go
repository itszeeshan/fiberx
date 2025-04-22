package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const serviceTemplate = `package services

import (
	{{if .WithDB}}"gorm.io/gorm"{{end}}
)

type {{.Name}}Service struct {
	{{if .WithDB}}db *gorm.DB{{end}}
}

func New{{.Name}}Service({{if .WithDB}}db *gorm.DB{{end}}) *{{.Name}}Service {
	return &{{.Name}}Service{
		{{if .WithDB}}db: db,{{end}}
	}
}

// Add your business logic methods below
`

type ServiceConfig struct {
	Name   string
	WithDB bool
}

func GenerateService(serviceName string) error {
	cfg := ServiceConfig{
		Name:   strings.Title(serviceName),
		WithDB: true, // Can be configurable via flags
	}

	// Create services directory if not exists
	if err := os.MkdirAll("services", 0755); err != nil {
		return fmt.Errorf("failed to create services directory: %w", err)
	}

	// Create service file
	path := filepath.Join("services", strings.ToLower(serviceName)+"_service.go")
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create service file: %w", err)
	}
	defer file.Close()

	// Parse and execute template
	tmpl := template.Must(template.New("service").Parse(serviceTemplate))
	if err := tmpl.Execute(file, cfg); err != nil {
		return fmt.Errorf("template execution failed: %w", err)
	}

	return nil
}
