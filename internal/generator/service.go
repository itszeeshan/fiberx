package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

type ServiceConfig struct {
	Name      string
	DBType    string
	Methods   []string
	WithRedis bool
}

const serviceTemplate = `package services

import (
	{{- if .DBType}}
	"gorm.io/gorm"
	{{- end}}
	{{- if .WithRedis}}
	"context"
	"github.com/redis/go-redis/v9"
	{{- end}}
)

{{if .DBType}}
type {{.Name | title}} struct {
	// Add your model fields here
}
{{end}}

type {{.Name | title}}Service struct {
	{{- if .DBType}}
	db *gorm.DB
	{{- end}}
	{{- if .WithRedis}}
	redis *redis.Client
	{{- end}}
}

func New{{.Name | title}}Service(
	{{- if .DBType}}db *gorm.DB{{if .WithRedis}}, {{end}}{{end}}
	{{- if .WithRedis}}redis *redis.Client{{end}}
) *{{.Name | title}}Service {
	return &{{.Name | title}}Service{
		{{- if .DBType}}
		db: db,
		{{- end}}
		{{- if .WithRedis}}
		redis: redis,
		{{- end}}
	}
}

{{if .DBType}}
{{if hasMethod .Methods "crud"}}
// CRUD Operations
func (s *{{.Name | title}}Service) Create(item interface{}) error {
	return s.db.Create(item).Error
}

func (s *{{.Name | title}}Service) GetByID(id uint, out interface{}) error {
	return s.db.First(out, id).Error
}

func (s *{{.Name | title}}Service) Update(id uint, updates interface{}) error {
	return s.db.Model(updates).Where("id = ?", id).Updates(updates).Error
}

func (s *{{.Name | title}}Service) Delete(id uint) error {
	return s.db.Delete(&{{.Name | title}}{}, id).Error
}
{{end}}

{{range .Methods}}
{{if eq . "create"}}
func (s *{{$.Name | title}}Service) Create(item interface{}) error {
	return s.db.Create(item).Error
}
{{end}}

{{if eq . "read"}}
func (s *{{$.Name | title}}Service) GetByID(id uint, out interface{}) error {
	return s.db.First(out, id).Error
}
{{end}}

{{if eq . "update"}}
func (s *{{$.Name | title}}Service) Update(id uint, updates interface{}) error {
	return s.db.Model(updates).Where("id = ?", id).Updates(updates).Error
}
{{end}}

{{if eq . "delete"}}
func (s *{{$.Name | title}}Service) Delete(id uint) error {
	return s.db.Delete(&{{$.Name | title}}{}, id).Error
}
{{end}}
{{end}}
{{else}}
// Business logic methods
func (s *{{.Name | title}}Service) SampleMethod() error {
	// Implement your business logic
	return nil
}
{{end}}

{{if .WithRedis}}
// Redis Operations
var ctx = context.Background()

func (s *{{.Name | title}}Service) CacheGet(key string) (string, error) {
	return s.redis.Get(ctx, key).Result()
}

func (s *{{.Name | title}}Service) CacheSet(key string, value interface{}) error {
	return s.redis.Set(ctx, key, value, 0).Err()
}
{{end}}`

func GenerateService(config ServiceConfig) error {
	funcMap := template.FuncMap{
		"title": title,
		"hasMethod": func(methods []string, target string) bool {
			return hasMethod(methods, target)
		},
	}

	if err := os.MkdirAll("services", 0755); err != nil {
		return fmt.Errorf("failed to create services directory: %w", err)
	}

	tmpl := template.Must(template.New("service").Funcs(funcMap).Parse(serviceTemplate))

	path := filepath.Join("services", strings.ToLower(config.Name)+"_service.go")
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create service file: %w", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, config); err != nil {
		return fmt.Errorf("template execution failed: %w", err)
	}

	return nil
}

func hasMethod(methods []string, target string) bool {
	for _, method := range methods {
		if strings.EqualFold(method, target) {
			return true
		}
	}
	return false
}

func title(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
