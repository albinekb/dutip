package utils

import (
	"errors"
	"log"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Filename is the __filename equivalent
func getFilename() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}

// Dirname is the __dirname equivalent
func getDirname() (string, error) {
	filename, err := getFilename()
	if err != nil {
		return "", err
	}
	return filepath.Dir(filename), nil
}

func ProjectRoot() string {
	dirname, err := getDirname()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(dirname, "..")
}

func ensureDutiInstalled() {
	command := exec.Command("duti", "-v")
	output, err := command.CombinedOutput()
	if err != nil {
		log.Fatalf("duti is not installed. Please install duti and try again.\n%s", string(output))
	}

}

// EnsureValidEnv ensures that the environment is valid for the tool to run
func EnsureValidEnv() {
	if runtime.GOOS != "darwin" {
		log.Fatal("This tool only works on macOS")
	}

	ensureDutiInstalled()
}
