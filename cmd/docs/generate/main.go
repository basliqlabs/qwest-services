package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("swag", "init",
		"--generalInfo", "main.go",
		"--output", "docs/api",
		"--parseInternal",
		"--parseDepth", "2",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to generate API documentation", err)
		os.Exit(1)
	}
}
