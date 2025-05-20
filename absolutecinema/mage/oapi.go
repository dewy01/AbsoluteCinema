//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	OpenAPIPath    = "../internal/openapi"
	HandlersOutput = "../internal/handlers"
)

func Generate() error {
	schemaPath := filepath.Join(OpenAPIPath, "schema")
	fmt.Println("Looking for OpenAPI specs in:", schemaPath)

	files, err := os.ReadDir(schemaPath)
	if err != nil {
		return err
	}

	var specs []string
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		lower := strings.ToLower(f.Name())
		if strings.HasSuffix(lower, ".yaml") || strings.HasSuffix(lower, ".yml") || strings.HasSuffix(lower, ".json") {
			specs = append(specs, filepath.Join(schemaPath, f.Name()))
		}
	}

	if len(specs) == 0 {
		return fmt.Errorf("no OpenAPI spec files found in %s", schemaPath)
	}

	for _, spec := range specs {
		base := filepath.Base(spec)
		name := base[:len(base)-len(filepath.Ext(base))]
		outFile := filepath.Join(HandlersOutput, name+".gen.go")

		fmt.Printf("Generating code for %s -> %s\n", spec, outFile)

		cmd := exec.Command("oapi-codegen",
			"-config", filepath.Join(OpenAPIPath, "oapi-codegen.yaml"),
			"-o", outFile,
			spec,
		)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to generate for %s: %w", spec, err)
		}
	}

	return nil
}
