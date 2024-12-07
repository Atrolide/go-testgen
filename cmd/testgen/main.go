package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Config structure to hold the configuration details
type Config struct {
	Suffix   string
	Languages map[string]LangConfig
}

// LangConfig holds the extension for each language
type LangConfig struct {
	Extension string
	Description string
}

func main() {
	// Get the full path to the executable
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error finding executable path:", err)
		return
	}

	// Get the directory where the executable is located
	exeDir := filepath.Dir(exePath)

	// Define the path to the config file (config.ini is in the same directory as the executable)
	configFile := filepath.Join(exeDir, "config.ini")

	// Load configuration from the config file
	config, err := loadConfig(configFile)
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	// Parse the command-line arguments
	lang := ""
	filename := ""
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "--language" && i+1 < len(os.Args) {
			lang = os.Args[i+1]
		}
		if os.Args[i] == "--filename" && i+1 < len(os.Args) {
			filename = os.Args[i+1]
		}
	}

	// Check if language and filename are provided
	if lang == "" || filename == "" {
		fmt.Println("Usage: testgen --language <language> --filename <filename>")
		return
	}

	// Normalize the language to lowercase
	lang = strings.ToLower(lang)

	// Check if the language is supported in the config
	langConfig, ok := config.Languages[lang]
	if !ok {
		fmt.Printf("Error: Language '%s' is not supported\n", lang)
		return
	}

	// Print the message for generating the test file
	fmt.Printf("Generating %s%s%s\n", filename, config.Suffix, langConfig.Extension)
}

// Load configuration from the config.ini file
func loadConfig(filename string) (*Config, error) {
	// Open the config.ini file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	defer file.Close()

	config := &Config{
		Languages: make(map[string]LangConfig),
	}

	// Read the config file line by line
	scanner := bufio.NewScanner(file)
	var currentLang string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Skip empty lines or comments
		if len(line) == 0 || line[0] == ';' {
			continue
		}

		// Parse sections and key-value pairs
		if line[0] == '[' && line[len(line)-1] == ']' {
			// New section (language)
			currentLang = strings.ToLower(line[1 : len(line)-1]) // Convert section name to lowercase
		} else if strings.Contains(line, "=") && currentLang != "" {
			// Parse key-value pairs for the current section
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				if key == "suffix" {
					config.Suffix = value
				} else if key == "extension" {
					// Save extension under the current language
					langConfig := config.Languages[currentLang]
					langConfig.Extension = value
					config.Languages[currentLang] = langConfig
				} else if key == "description" {
					// Save description under the current language
					langConfig := config.Languages[currentLang]
					langConfig.Description = value
					config.Languages[currentLang] = langConfig
				}
			}
		}
	}

	// Return any error encountered during the scan
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	return config, nil
}
