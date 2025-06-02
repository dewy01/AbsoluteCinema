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
	HandlersOutput = "../internal/openapi/gen"
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
		packageName := name + "gen"
		outDir := filepath.Join(HandlersOutput, packageName)

		if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create output dir %s: %w", outDir, err)
		}

		outFile := filepath.Join(outDir, name+".gen.go")

		fmt.Printf("Generating code for %s -> %s (package: %s)\n", spec, outFile, packageName)

		cmd := exec.Command("oapi-codegen",
			"-config", filepath.Join(OpenAPIPath, "oapi-codegen.yaml"),
			"-package", packageName,
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
