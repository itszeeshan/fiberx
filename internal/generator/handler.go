package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const handlerTemplate = `package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type {{.Name}}Handler struct {
	// Add service dependencies here
}

func New{{.Name}}Handler() *{{.Name}}Handler {
	return &{{.Name}}Handler{}
}

func (h *{{.Name}}Handler) RegisterRoutes(router fiber.Router) {
	{{- range .Methods}}
	router.{{.}}("/{{$.RouteName}}", h.handle{{.}})
	{{- end}}
}
{{range .Methods}}

func (h *{{$.Name}}Handler) handle{{.}}(c *fiber.Ctx) error {
	// Implement {{.}} {{$.Name}} logic
	return c.JSON(fiber.Map{
		"message": "{{.}} {{$.Name}} handler",
	})
}
{{end}}
`

type HandlerConfig struct {
	Name      string
	RouteName string
	Methods   []string
}

func GenerateHandler(name string, methods []string) error {
	cfg := HandlerConfig{
		Name:      strings.Title(name),
		RouteName: strings.ToLower(name),
		Methods:   methods,
	}

	// Create handlers directory if not exists
	if err := os.MkdirAll("handlers", 0755); err != nil {
		return fmt.Errorf("failed to create handlers directory: %w", err)
	}

	// Create handler file
	path := filepath.Join("handlers", fmt.Sprintf("%s_handler.go", strings.ToLower(name)))
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create handler file: %w", err)
	}
	defer file.Close()

	// Parse and execute template
	tmpl := template.Must(template.New("handler").Parse(handlerTemplate))
	if err := tmpl.Execute(file, cfg); err != nil {
		return fmt.Errorf("template execution failed: %w", err)
	}

	return nil
}
