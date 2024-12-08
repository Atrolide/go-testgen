package generator

import (
	"fmt"
	"github.com/mitchellh/colorstring"
	"time"
)

// TEMP: Mock method for testing purposes. Replace with real implementation.
// FIXME: Implement a proper file generation solution.

func FileGen() {
	fmt.Printf("\nGenerating a file...\n\n")

	startTime := time.Now()
	time.Sleep(time.Duration(0.021 * float64(time.Second)))
	elapsedTime := time.Since(startTime)

	elapsedTimeMs := elapsedTime.Seconds() * 1000
	elapsedTimeFormatted := fmt.Sprintf("%.2f ms", elapsedTimeMs)

	successMessage := fmt.Sprintf("[bold][green]Test File Generated in %s\n", elapsedTimeFormatted)
	fmt.Println(colorstring.Color(successMessage))
}
