package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/blang/semver/v4"
)

// Version holds the current version number
const VERSION_STRING = "0.1.10"

var Version = VERSION_STRING

func bumpVersion(from string, level string) string {
	version, err := semver.Parse(from)
	if err != nil {
		log.Fatalf("Failed to parse version: %v", err)
	}

	if err != nil {
		log.Fatalf("Failed to bump version: %v", err)
	}

	err = version.IncrementPatch()
	if err != nil {
		log.Fatalf("Failed to bump version: %v", err)
	}

	return version.String()
}

func filename() string {
	dir := ProjectRoot()
	dir = filepath.Join(dir, "utils")

	path := filepath.Join(dir, "version.go")
	// Check if the file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("File %s does not exist", path)
	}

	return path
}

func BumpCmd() {
	// bump the version, update the version.go file
	nextVersion := bumpVersion(VERSION_STRING, "patch")
	fmt.Println(nextVersion)
	file := filename()
	fmt.Println(file)

	target := fmt.Sprintf("const VERSION_STRING = \"%s\"", VERSION_STRING)
	content := strings.Replace(target, VERSION_STRING, nextVersion, 1)

	currentContent, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	if strings.Contains(string(currentContent), nextVersion) {
		log.Fatalf("Version is already up to date")
	}

	if !strings.Contains(string(currentContent), target) {
		log.Fatalf("Version is not found in the file")
	}

	updatedContent := strings.Replace(string(currentContent), target, content, 1)

	if updatedContent == string(currentContent) {
		log.Fatalf("Failed to update version")
	}

	err = os.WriteFile(file, []byte(updatedContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}
