package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	colorRed   = "\033[0;31m"
	colorGreen = "\033[0;32m"
	colorNC    = "\033[0m"
)

func runCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return out.String(), err
}

func getChangedGoFiles() ([]string, error) {
	out, err := runCommand("git", "diff", "--cached", "--name-only", "--diff-filter=ACMR")
	if err != nil {
		return nil, err
	}

	var goFiles []string
	for _, file := range strings.Split(out, "\n") {
		if strings.TrimSpace(file) != "" && strings.HasSuffix(file, ".go") {
			goFiles = append(goFiles, file)
		}
	}
	return goFiles, nil
}

func getUniqueDirs(files []string) []string {
	dirsMap := make(map[string]bool)
	for _, file := range files {
		dirsMap[filepath.Dir(file)] = true
	}

	var dirs []string
	for dir := range dirsMap {
		dirs = append(dirs, dir)
	}
	return dirs
}

func main() {
	fmt.Println("Running pre-commit checks...")

	// Get changed Go files
	changedFiles, err := getChangedGoFiles()
	if err != nil {
		fmt.Printf("%sError getting changed files: %v%s\n", colorRed, err, colorNC)
		os.Exit(1)
	}

	if len(changedFiles) == 0 {
		fmt.Println("No Go files changed")
		os.Exit(0)
	}

	// Get unique directories
	changedDirs := getUniqueDirs(changedFiles)

	// Check 1: Go Format
	fmt.Println("Running go fmt...")
	for _, file := range changedFiles {
		_, err := runCommand("go", "fmt", file)
		if err != nil {
			fmt.Printf("%s❌ go fmt failed%s\n", colorRed, colorNC)
			os.Exit(1)
		}
	}
	fmt.Printf("%s✓ go fmt passed%s\n", colorGreen, colorNC)

	// Check 2: Go Vet
	fmt.Println("Running go vet...")
	for _, dir := range changedDirs {
		_, err := runCommand("go", "vet", "./"+dir+"/...")
		if err != nil {
			fmt.Printf("%s❌ go vet failed%s\n", colorRed, colorNC)
			os.Exit(1)
		}
	}
	fmt.Printf("%s✓ go vet passed%s\n", colorGreen, colorNC)

	// Check 3: Go Test
	fmt.Println("Running tests for changed directories...")
	for _, dir := range changedDirs {
		fmt.Printf("Testing directory: %s\n", dir)
		_, err := runCommand("go", "test", "./"+dir+"/...")
		if err != nil {
			fmt.Printf("%s❌ tests failed in %s%s\n", colorRed, dir, colorNC)
			os.Exit(1)
		}
	}
	fmt.Printf("%s✓ all tests passed%s\n", colorGreen, colorNC)

	// Check 4: Unwanted files
	out, err := runCommand("git", "diff", "--cached", "--name-only")
	if err != nil {
		fmt.Printf("%sError checking for unwanted files: %v%s\n", colorRed, err, colorNC)
		os.Exit(1)
	}

	unwantedPatterns := []string{".exe$", ".test$", ".out$", ".log$", ".env$"}
	for _, file := range strings.Split(out, "\n") {
		for _, pattern := range unwantedPatterns {
			if strings.HasSuffix(file, pattern) {
				fmt.Printf("%s❌ Attempting to commit forbidden file: %s%s\n", colorRed, file, colorNC)
				os.Exit(1)
			}
		}
	}
	fmt.Printf("%s✓ no unwanted files found%s\n", colorGreen, colorNC)

	fmt.Printf("%sAll pre-commit checks passed!%s\n", colorGreen, colorNC)
}
