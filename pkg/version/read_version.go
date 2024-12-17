package version

import (
	"fmt"
	"github.com/mitchellh/colorstring"
	"runtime"
)

var (
	Version  string
	Platform = runtime.GOOS + "_" + runtime.GOARCH
	Build    string
)

func ReadVersion() {
	// Print the version details
	fmt.Println(colorstring.Color("[bold][yellow]Testgen version:") + " " + Version)
	fmt.Println(colorstring.Color("[bold][yellow]Platform:       ") + " " + Platform)
	fmt.Println(colorstring.Color("[bold][yellow]Build:          ") + " " + Build)
}
