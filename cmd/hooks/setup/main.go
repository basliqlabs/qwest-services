package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	// Get the git hooks directory
	hookDir := filepath.Join(".git", "hooks")
	scriptDir := filepath.Join("scripts", "git-hooks")

	// Create hooks directory if it doesn't exist
	if err := os.MkdirAll(hookDir, 0755); err != nil {
		fmt.Printf("Error creating hooks directory: %v\n", err)
		os.Exit(1)
	}

	// Copy pre-commit hook
	preCommitSrc := filepath.Join(scriptDir, "pre-commit")
	preCommitDst := filepath.Join(hookDir, "pre-commit")
	if runtime.GOOS == "windows" {
		preCommitDst += ".exe"
	}

	// Read the source file
	content, err := os.ReadFile(preCommitSrc)
	if err != nil {
		fmt.Printf("Error reading pre-commit source: %v\n", err)
		os.Exit(1)
	}

	// Write to destination
	if err := os.WriteFile(preCommitDst, content, 0755); err != nil {
		fmt.Printf("Error writing pre-commit hook: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Git hooks installed successfully!")
}
