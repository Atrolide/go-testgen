# Testgen

`Testgen` is a tool designed to generate sample test files for multiple programming languages.

![Go version](https://img.shields.io/badge/Go-1.23-blue)

> [!NOTE]  
> The tool is currently under development.

## Go Documentation
For detailed package documentation, visit the `Testgen` page on [pkg.go.dev](https://pkg.go.dev) by clicking on the badge below:

[![Go Reference](https://pkg.go.dev/badge/github.com/Atrolide/go-testgen@master.svg)](https://pkg.go.dev/github.com/Atrolide/go-testgen@master)


## Building the Project

To build the project, you can use the provided [`Makefile`](Makefile):

- **`make build`**  
  Build binaries for Linux and Windows.

- **`make zip`**  
  Zip Linux and Windows binaries into archives.

- **`make all`**  
  Build and zip the binaries with one command.

> [!IMPORTANT]  
> Requires installation of `GNU MAKE`.<br><br>
> Check [GNU MAKE Installation Instructions](INSTALL.md#gnu-make)

![Make version](https://img.shields.io/badge/Make-4.3-red)

## Utility Scripts

The project also uses a [`Rakefile`](Rakefile) for various other tasks:

### Available Commands

- **`rake fmt`**  
  Format all Go files in the project using `go fmt`.

- **`rake test`**  
  Run all tests located in the `./tests` directory using `go test`.

- **`rake gosec`**  
  Perform a security scan on all Go files to identify potential vulnerabilities using `gosec`.

- **`rake`**  
  Run all of the above tasks: formatting, testing, and security scanning.

> [!IMPORTANT]  
> Requires installation of `ruby`, `gem`, and `rake`.<br><br>
> Check [Ruby Installation Instructions](INSTALL.md#ruby--gem--rake)

![Ruby version](https://img.shields.io/badge/Ruby-3.0-red)
![Gem version](https://img.shields.io/badge/Gem-3.3.5-red)
![Rake version](https://img.shields.io/badge/Rake-13.2.1-red)

> [!IMPORTANT]  
> Requires installation of `go fmt`, `go test`, and `go sec`.<br><br>
> Check [Go Tools Installation Instructions](INSTALL.md#gofmt--gotest--gosec)

## License
This project is licensed under the [MIT License](https://opensource.org/license/mit). For more details, refer to the [LICENSE](LICENSE) file.

![License: MIT](https://img.shields.io/badge/License-MIT-purple)