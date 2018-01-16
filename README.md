# pygz/subprocess 

[![GitHub release](https://img.shields.io/github/release/pygz/subprocess.svg?style=flat-square)](https://github.com/pygz/subprocess/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/pygz/subprocess)
[![Build Status](https://semaphoreci.com/api/v1/pygz/subprocess/branches/master/badge.svg)](https://semaphoreci.com/pygz/subprocess) 
[![Build status](https://ci.appveyor.com/api/projects/status/722snh8mfavt0j93/branch/master?svg=true)](https://ci.appveyor.com/project/chrissimpkins/subprocess/branch/master) 

## About

subprocess is a Go library that returns standard output, standard error, and exit status code data from new spawned processes on Linux, macOS, and Windows platforms.  It was inspired by the Python subprocess standard library module.

The subprocess library API has not reached a stable status yet and backwards incompatible changes may take place at any point before the v1.0.0 release.  The API will be defined as stable and the library will follow SemVer versioning for all backwards incompatible changes as of the v1.0.0 release.

## Install

The subprocess package does not include external dependencies. It is built with the Go standard library.

Install the subprocess library locally for testing and development use with the following command:

```
go get github.com/pygz/subprocess
```

## Usage

subprocess exposes two public functions and a public struct with standard output, standard error, and exit status code response data from executable files that can be called on the command line.  [Full API documentation is available on GoDoc](https://godoc.org/github.com/pygz/subprocess).

### Import `subprocess` into your source files

```go
package main

import (
	"github.com/pygz/subprocess"
)
```

### Public Data Types

#### `subprocess.Response`

The subprocess package defines the `Response` public data type with standard output, standard error, and exit status code fields.  This is populated and returned to the calling code when you run an executable file with the public functions that are available in the subprocess package.

```go
type Response struct {
	StdOut   string
	StdErr   string
	ExitCode int
}
```

### Public Functions

#### `subprocess.Run`

```go
func Run(executable string, args ...string) Response
```

The `Run()` function runs an executable file with optional arguments and returns the standard output, standard error, and exit status code data in a `Response` struct.  Include one or more arguments to the executable as additional function parameters.

##### Example

```go
package main

import (
	"fmt"
	"github.com/pygz/subprocess"
)

func main() {
	response := Run("ls", "-l")
	// print the standard output stream data
	fmt.Printf("%s", response.StdOut)
	// print the standard error stream data
	fmt.Printf("%s", response.StdErr)
	// print the exit status code integer value
	fmt.Printf("%d", response.ExitCode)
}
```

#### `subprocess.RunShell()`

```go
func RunShell(shell string, shellflag string, command ...string) Response
```

The `RunShell()` function runs an executable file with a shell and returns the standard output, standard error, and exit status code data in a `Response` struct.  The default shell for Linux and macOS platforms is `/bin/sh`.  The default shell for Windows is the `cmd.exe` command prompt. The shell can be modified by defining the `shell` function parameter.  A shell flag is included to indicate that the argument that follows is to be executed by the shell.  The default flag on macOS and Linux platforms is `-c`.  On Windows, this is `/C`.  This flag can be modified in the `shellflag` parameter.  Define the command to be executed as one or more parameters at the end of the function call.

##### Example with the default shell on macOS/Linux

```go
package main

import (
	"fmt"
	"github.com/pygz/subprocess"
)

func main() {
	response := RunShell("", "", "ls", "-l")
	// print the standard output stream data
	fmt.Printf("%s", response.StdOut)
	// print the standard error stream data
	fmt.Printf("%s", response.StdErr)
	// print the exit status code integer value
	fmt.Printf("%d", response.ExitCode)
}
```

##### Example with the default shell on Windows

```go
package main

import (
	"fmt"
	"github.com/pygz/subprocess"
)

func main() {
	response := RunShell("", "", "dir", "/AD")
	// print the standard output stream data
	fmt.Printf("%s", response.StdOut)
	// print the standard error stream data
	fmt.Printf("%s", response.StdErr)
	// print the exit status code integer value
	fmt.Printf("%d", response.ExitCode)
}
```

##### Example with redefined shell on macOS/Linux

```go
package main

import (
	"fmt"
	"github.com/pygz/subprocess"
)

func main() {
	response := RunShell("/usr/local/bin/zsh", "", "ls", "-l")
	// print the standard output stream data
	fmt.Printf("%s", response.StdOut)
	// print the standard error stream data
	fmt.Printf("%s", response.StdErr)
	// print the exit status code integer value
	fmt.Printf("%d", response.ExitCode)
}
```

##### Example with redefined shell on Windows

```go
package main

import (
	"fmt"
	"github.com/pygz/subprocess"
)

func main() {
	response := RunShell("bash", "-c", "ls", "-l")
	// print the standard output stream data
	fmt.Printf("%s", response.StdOut)
	// print the standard error stream data
	fmt.Printf("%s", response.StdErr)
	// print the exit status code integer value
	fmt.Printf("%d", response.ExitCode)
}
```

### Contributing

Contributions to the project are welcomed. Please submit changes in a pull request on the Github repository.

### Testing

You can execute source code unit tests and obtain source code coverage data locally by downloading the source repository and executing the following command in the root of the source repository:

```
$ go test -v -cover ./...
```

Go must be installed on your system in order to execute this command.  

### Acknowledgments

The subprocess library was inspired by the Python standard library subprocess module.  Source code for the exit status code retrieval was based on source discussed in the Stack Overflow posts [here](https://stackoverflow.com/a/40770011) and [here](https://stackoverflow.com/a/10385867). A big thanks to Michael (@texhex) and JM (@jublo) for their input and feedback on the Windows platform implementation.

### License

The subprocess library is licensed under the [MIT license](LICENSE).

