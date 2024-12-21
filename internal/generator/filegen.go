package generator

import (
	"fmt"
	"github.com/Atrolide/go-testgen/pkg/helpers"
	"github.com/mitchellh/colorstring"
	"time"
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

	fmt.Printf("\nGenerating a test for the file: %s...\n\n", fileFlag)

	startTime := time.Now()
	time.Sleep(time.Duration(0.021 * float64(time.Second))) // Simulate file generation delay
	elapsedTime := time.Since(startTime)

	elapsedTimeMs := elapsedTime.Seconds() * 1000
	elapsedTimeFormatted := fmt.Sprintf("%.2f ms", elapsedTimeMs)

	successMessage := fmt.Sprintf("[bold][green]Test File Generated for '%s' in %s\n", fileFlag, elapsedTimeFormatted)
	fmt.Println(colorstring.Color(successMessage))
}
