package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	utils "github.com/albinekb/dutip/utils"
)

func getExtensions() []string {
	return []string{
		"ts",
		"tsx",
		"md",
		"mdx",
		"json",
		"js",
		"jsx",
		"html",
		"css",
		"yml",
		"yaml",
		"toml",
		"ini",
		"conf",
		"cfg",
	}
}

func getAllAppNames() ([]string, error) {
	command := "mdfind 'kMDItemKind == \"Application\"'"

	output, err := exec.Command("bash", "-c", command).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to get app ID: %w\noutput: %s", err, string(output))
	}

	lines := strings.Split(string(output), "\n")
	// Keep only the last part of each line (*.app)
	appNames := make([]string, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, "/")
		appName := parts[len(parts)-1]
		if !strings.HasSuffix(appName, ".app") {
			continue
		}
		appNames = append(appNames, appName)
	}

	return appNames, nil
}

func isAppId(appName string) (bool, error) {
	// mdfind "kMDItemCFBundleIdentifier == 'com.microsoft.VSCodeInsiders'"
	command := fmt.Sprintf("mdfind 'kMDItemCFBundleIdentifier == \"%s\"'", appName)

	output, err := exec.Command("bash", "-c", command).CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("failed to get app ID: %w\noutput: %s", err, string(output))
	}

	return true, nil
}

func getAppIdFromName(appName string) (string, error) {
	if !strings.HasSuffix(appName, ".app") {
		appId, err := isAppId(appName)
		if err != nil {
			return "", fmt.Errorf("failed to check if app ID exists: %w", err)
		}
		if appId {
			return appName, nil
		}
	}
	// osascript -e 'id of app "IINA.app"'

	command := fmt.Sprintf("osascript -e 'id of app \"%s\"'", appName)

	output, err := exec.Command("bash", "-c", command).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get app ID: %w\noutput: %s", err, string(output))
	}

	appID := strings.TrimSpace(string(output))
	if appID == "" {
		return "", fmt.Errorf("app ID not found for %s", appName)
	}

	return appID, nil
}

func getCurrentAssignedExtensions(appId string) ([]string, error) {
	allExtensions := getExtensions()
	assignedExtensions := []string{}

	for _, ext := range allExtensions {
		command := fmt.Sprintf("duti -x %s", ext)
		output, err := exec.Command("bash", "-c", command).CombinedOutput()
		if err != nil {
			continue // Skip if there's an error, as it might mean the extension is not assigned
		}

		lines := strings.Split(string(output), "\n")
		if len(lines) >= 3 && strings.TrimSpace(lines[2]) == appId {
			assignedExtensions = append(assignedExtensions, ext)
		}
	}

	if len(assignedExtensions) == 0 {
		return nil, fmt.Errorf("no extensions found for %s", appId)
	}

	return assignedExtensions, nil
}

func confirm(message string) bool {
	fmt.Println(message)
	var response string
	fmt.Scanln(&response)
	return strings.ToLower(response) == "y" || strings.ToLower(response) == "yes"
}

func setDefaultApp(appId string, extension string, force bool) (bool, error) {
	command := fmt.Sprintf("duti -s %s .%s all", appId, extension)

	if !force && !confirm(fmt.Sprintf("Are you sure you want to set the default app for .%s to %s? (y/n):\n%s\n", extension, appId, command)) {
		fmt.Println("Operation cancelled.")
		return false, nil
	}

	output, err := exec.Command("bash", "-c", command).CombinedOutput()
	if err != nil {
		return false, fmt.Errorf("failed to set default app: %w\noutput: %s", err, string(output))
	}

	return true, nil
}

func main() {
	// Add a new flag for version
	version := flag.Bool("v", false, "Print version information")
	bump := flag.Bool("bump", false, "Bump the version")

	// allExtensions := getExtensions()

	force := flag.Bool("yes", false, "Force the operation to run without confirmation")
	flag.BoolVar(force, "y", false, "")
	fromInput := flag.String("from", "", "The app name to change from")
	flag.StringVar(fromInput, "f", "", "")
	toInput := flag.String("to", "", "The app name to change to")
	flag.StringVar(toInput, "t", "", " ")

	// Parse the flags
	flag.Parse()
	if *bump {
		utils.BumpCmd()
		return
	}

	// Check if no flags are provided or if -v flag is used
	if flag.NFlag() == 0 || *version {
		fmt.Printf("Version: %s\n", utils.Version)
		return
	}

	utils.EnsureValidEnv()

	// Check if both flags are provided
	if *fromInput == "" || *toInput == "" {
		fmt.Println("Error: Both -from and -to flags are required")
		flag.Usage()
		os.Exit(1)
	}

	// Print the app names (placeholder for actual functionality)
	fmt.Printf("Changing default app from '%s' to '%s'\n", *fromInput, *toInput)
	fromApp, err := getAppIdFromName(*fromInput)
	if err != nil {
		log.Fatalf("Failed to get app ID from name: %v", err)
	}

	toApp, err := getAppIdFromName(*toInput)
	if err != nil {
		log.Fatalf("Failed to get app ID from name: %v", err)
	}

	fmt.Printf("From app ID: %s\nTo app ID: %s\n", fromApp, toApp)

	// Get current assigned extensions
	extensions, err := getCurrentAssignedExtensions(fromApp)
	if err != nil {
		fmt.Printf("Warning: %v\n", err)
		fmt.Println("Would you like to proceed with all supported extensions? (y/n): ")
		var response string
		fmt.Scanln(&response)
		if strings.ToLower(response) != "y" && strings.ToLower(response) != "yes" {
			fmt.Println("Operation cancelled.")
			return
		}
		extensions = getExtensions()
	}

	// Display extensions and ask for confirmation
	fmt.Println("The following extensions will be changed:")
	for _, ext := range extensions {
		fmt.Printf("- .%s\n", ext)
	}

	if !*force {
		fmt.Print("Do you want to proceed? (y/n): ")
		var response string
		fmt.Scanln(&response)

		if strings.ToLower(response) != "y" && strings.ToLower(response) != "yes" {
			fmt.Println("Operation cancelled.")
			return
		}
	}
	// Proceed with changes
	changedCount := 0
	for _, ext := range extensions {
		fmt.Printf("Changing default app for .%s... ", ext)
		changed, err := setDefaultApp(toApp, ext, *force)
		if err != nil {
			fmt.Printf("Failed: %v\n", err)
		} else if changed {
			fmt.Println("Success")
			changedCount++
		}
	}

	fmt.Printf("\nOperation completed. Successfully changed %d out of %d extensions.\n", changedCount, len(extensions))
}
