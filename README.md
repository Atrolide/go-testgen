# Testgen

`Testgen` is a tool designed to generate sample test files for multiple programming languages.

> [!NOTE]  
> The tool is currently under development.

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

## License

This project is licensed under the [MIT License](https://opensource.org/license/mit). For more details, refer to the [LICENSE](LICENSE) file.

[![Go Reference](https://pkg.go.dev/badge/github.com/Atrolide/go-testgen.svg)](https://pkg.go.dev/github.com/Atrolide/go-testgen)
