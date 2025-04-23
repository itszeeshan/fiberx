// file: internal/generator/addFeature.go
package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var featurePaths = map[string]string{
	"postgres": "config/database.go", // Check in config for postgres
	"redis":    "cache/redis.go",     // Check in cache for redis
	"jwt":      "auth/jwt.go",        // Check in auth for JWT
	"viper":    "config/config.go",   // Check in config for viper
	// Add other features and their respective paths here
}

func AddFeature(feature string) error {
	if path, exists := featurePaths[feature]; exists {
		if _, err := os.Stat(path); err == nil {
			return fmt.Errorf("feature %s already exists at %s", feature, path)
		}
	} else {
		return fmt.Errorf("unknown feature %s", feature)
	}

	src := filepath.Join("templates", "features", feature)

	config := ProjectConfig{Name: ".", Features: map[string]bool{feature: true}}
	err := ProcessTemplates(src, config)
	if err != nil {
		return fmt.Errorf("error processing templates for feature %s: %w", feature, err)
	}

	// Add feature-specific dependencies to go.mod if they exist in the featureDependencies map
	if deps, ok := featureDependencies[feature]; ok {
		for _, dep := range deps {
			err := appendGoMod(dep)
			if err != nil {
				return fmt.Errorf("failed to update go.mod with %s: %w", dep, err)
			}
		}
	}

	return nil
}

// appendGoMod adds the required dependency to the go.mod file
func appendGoMod(dep string) error {
	cmd := exec.Command("go", "get", dep)
	cmd.Dir = "."
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("go get failed for %s: %v\n%s", dep, err, output)
	}
	return nil
}
