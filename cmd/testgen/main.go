package main

import (
	"fmt"
	"github.com/Atrolide/go-testgen/internal/commands" // Importing commands from internal/commands
)

func main() {
	if err := commands.Execute(); err != nil {
		// Print the error message when a command fails
		fmt.Println(err)
	}
}
