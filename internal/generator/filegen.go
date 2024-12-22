package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/Atrolide/go-testgen/pkg/helpers"
	"github.com/mitchellh/colorstring"
)

// TEMP: Mock method for testing purposes. Replace with real implementation.
// FIXME: Implement a proper file generation solution.

func FileGenByFile(fileFlag string) {
	// Use IsFile to check if the path is a valid file
	if !helpers.IsFile(fileFlag) {
		errorMessage := fmt.Sprintf("\n[bold][red]Error: The '%s' either does not exist or is not a regular file.\n", fileFlag)
		fmt.Println(colorstring.Color(errorMessage))
		return
	}

	// Extract file name and extension from the provided path
	fileName := filepath.Base(fileFlag)
	ext := filepath.Ext(fileName)
	baseName := fileName[:len(fileName)-len(ext)] // Strip the extension

	// Define the test directory
	testDir := "tests"

	// Check if 'tests' directory exists, if not, create it
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		fmt.Printf("Creating directory '%s'...\n", testDir)
		if err := os.Mkdir(testDir, os.ModePerm); err != nil {
			errorMessage := fmt.Sprintf("\n[bold][red]Error: Failed to create directory '%s': %v\n", testDir, err)
			fmt.Println(colorstring.Color(errorMessage))
			return
		}
	}

	// Define the test file name
	testFileName := fmt.Sprintf("%s_test%s", baseName, ext)
	testFilePath := filepath.Join(testDir, testFileName)

	// Check if the test file already exists
	if _, err := os.Stat(testFilePath); !os.IsNotExist(err) {
		// If the file exists, return without overwriting
		errorMessage := fmt.Sprintf("\n[bold][red]Error: The test file '%s' already exists. Aborting generation to avoid overwriting.\n", testFilePath)
		fmt.Println(colorstring.Color(errorMessage))
		return
	}

	// Create the test file
	testFile, err := os.Create(testFilePath)
	if err != nil {
		errorMessage := fmt.Sprintf("\n[bold][red]Error: Failed to create test file '%s': %v\n", testFilePath, err)
		fmt.Println(colorstring.Color(errorMessage))
		return
	}
	defer testFile.Close()

	// Add a comment indicating the file was generated
	comment := "// Generated with testgen @github.com/Atrolide/go-testgen\n"
	if _, err := testFile.WriteString(comment); err != nil {
		errorMessage := fmt.Sprintf("\n[bold][red]Error: Failed to write to test file '%s': %v\n", testFilePath, err)
		fmt.Println(colorstring.Color(errorMessage))
		return
	}

	// Output success message
	successMessage := fmt.Sprintf("[bold][green]Test File Generated for '%s' at '%s'\n", fileFlag, testFilePath)
	fmt.Println(colorstring.Color(successMessage))
}
