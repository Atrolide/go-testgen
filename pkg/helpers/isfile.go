package helpers

import (
	"os"
)

// Checks if a given path exists and is a regular file.
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false // Either file does not exist or some other error occurred
	}
	return !info.IsDir() // Return true only if the path is not a directory
}
