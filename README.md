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

Install the library locally for testing and use in development with the following command:

```
go get github.com/pygz/subprocess
```

## Usage

subprocess exposes two public functions and a public struct with standard output, standard error, and exit status code response data from system executable calls.  [Full API documentation is available on GoDoc](https://godoc.org/github.com/pygz/subprocess).

**NOTE**: The subprocess library does not currently support automated shell escaping of strings that are used for execution of system commands.  _These commands have the potential to do significant harm_.  Please understand these concepts and how to avoid problems before you use this library, particularly if you intend to open your application to command definitions at runtime by untrusted sources.  The [Python `shlex.quote` documentation](https://docs.python.org/3.6/library/shlex.html#shlex.quote) is a good place to start.

### Import `subprocess` into your source files

```go
package main

import (
	"github.com/pygz/subprocess"
)
```

### Public Data Types

#### `subprocess.Response`

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

The `Run()` function executes a system executable call with optional arguments and returns the standard output, standard error, and exit status code data in a `Response` struct.  Include one or more arguments to the executable as additional function parameters.

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

The `RunShell()` function executes a system executable with a default or function parameter defined shell and returns the standard output, standard error, and exit status code data in a `Response` struct.  The default shell for Linux and macOS platforms is `/bin/sh`.  The default shell for Windows is `bash` and Windows 10+ with bash installed is a mandatory dependency for use of the default settings in this function on the Windows platform.  The shell can be modified by defining the `shell` function parameter.  By default, all platforms use the `-c` flag to the shell executable as an indicator that subsequent arguments define a command that is to be executed by the shell.  This flag can be modified in the `shellflag` parameter.  Include one or more arguments for the command that is to be executed as additional function parameters.

##### Example with default shell

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

##### Example with redefined shell

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

### Contributing

Contributions to the project are welcomed. Please submit changes as pull requests on the Github repository.

### Testing

You can execute source code unit tests and obtain source code coverage data locally by downloading the source repository and executing the following command in the root of the source repository:

```
$ go test -v -cover ./...
```

Go must be installed on your system in order to execute this command.  

### Acknowledgments

The subprocess library was inspired by the Python standard library subprocess module.  Source code for the exit status code retrieval was based on source discussed in the Stack Overflow posts [here](https://stackoverflow.com/a/40770011) and [here](https://stackoverflow.com/a/10385867).

### License

The subprocess library is licensed under the [MIT license](LICENSE).

