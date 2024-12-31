package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Generate OpenAPI spec
	cmd := exec.Command("swag", "init",
		"--generalInfo", "main.go",
		"--output", "docs/api",
		"--parseInternal",
		"--parseDepth", "2",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to generate API documentation")
		os.Exit(1)
	}

	// Convert to OpenAPI 3.0 format
	swaggerFile := filepath.Join("docs", "api", "swagger.json")
	openapiFile := filepath.Join("docs", "api", "openapi.yaml")

	cmd = exec.Command("npx", "swagger2openapi",
		swaggerFile,
		"--outfile", openapiFile,
		"--yaml",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to convert API documentation")
		os.Exit(1)
	}
}
