// Package subprocess provides support for standard output/error pipe data & exit status codes with new spawned system processes
package subprocess

import (
	"bytes"
	"os/exec"
	"runtime"
	"syscall"
	"strings"
)

// Response is a struct that is returned from the public functions in the subprocess package.  It contains the following
// fields:
//
//     Response.StdOut - (string) standard output stream cast to a string
//     Response.StdErr - (string) standard error stream cast to a string
//     Response.ExitCode - (int) executable exit status code as an integer
type Response struct {
	StdOut   string
	StdErr   string
	ExitCode int
}

/*    ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
 *    ┃                                                                              ┃
 *    ┃                                                                              ┃
 *    ┃                          ______      _     _ _                               ┃
 *    ┃                          | ___ \    | |   | (_)                              ┃
 *    ┃                          | |_/ /   _| |__ | |_  ___                          ┃
 *    ┃                          |  __/ | | | '_ \| | |/ __|                         ┃
 *    ┃                          | |  | |_| | |_) | | | (__                          ┃
 *    ┃                          \_|   \__,_|_.__/|_|_|\___|                         ┃
 *    ┃                                                                              ┃
 *    ┃                                                                              ┃
 *    ┃                                                                              ┃
 *    ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
 */

// Run is a public function that executes a system command and returns the standard output stream,
// standard error stream, and exit status code data in a returned subprocess.Response struct.
// Run takes the following parameters:
//
//  executable (string) - the system executable for the command call
//  args (...string) - comma delimited list of arguments to executable
//
// Example:
//
//     func main() {
//         response := Run("go", "--help")
//         fmt.Printf("%s\n", response.StdOut)
//         fmt.Printf("%s\n", response.StdErr)
//         fmt.Printf("%d\n", response.ExitCode)
//     }
func Run(executable string, args ...string) Response {
	// define function variables
	var res Response
	var outbuf, errbuf bytes.Buffer

	// define the system executable call
	cmd := exec.Command(executable, args...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf
	// execute the system command
	err := cmd.Run()
	// define the returned object fields with the data returned
	res.StdOut = outbuf.String()
	res.StdErr = errbuf.String()
	if err != nil {
		res.ExitCode = getErrorExitCode(err)
	} else {
		res.ExitCode = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	}
	if res.StdErr == "" && res.ExitCode != 0 {
		res.StdErr = err.Error() // return the error raised in standard error stream formatted as a string
	}

	return res
}

func RunShell(shell string, shellflag string, command ...string) Response {
	// define the default shell by platform
	if shell == "" {
		if runtime.GOOS == "windows" {
			shell = `cmd`       // defined as "cmd \C" for Windows
		} else {
			shell = `/bin/sh`   // defined as "/bin/sh -c" for *nix (including macOS)
		}
	}
	// define the default shell flag for execution of system executables
	if shellflag == "" {
		if runtime.GOOS == "windows" {
			shellflag = `\C`
		} else {
			shellflag = `-c`
		}
	}
	// define function variables
	var res Response
	var outbuf, errbuf bytes.Buffer

	// define the system executable call
	shellExecString := strings.Join(command, " ")
	cmd := exec.Command(shell, shellflag, shellExecString)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf
	// execute the system command
	err := cmd.Run()
	// define the returned object fields with the data returned
	res.StdOut = outbuf.String()
	res.StdErr = errbuf.String()
	if err != nil {
		res.ExitCode = getErrorExitCode(err)
	} else {
		res.ExitCode = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	}
	if res.StdErr == "" && res.ExitCode != 0 {
		res.StdErr = err.Error() // return the error raised in standard error stream formatted as a string
	}

	return res
}

//func Pipe() {
//	// TODO
//}

/*    ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
 *    ┃                                                                              ┃
 *    ┃                                                                              ┃
 *    ┃                       ______     _            _                              ┃
 *    ┃                       | ___ \   (_)          | |                             ┃
 *    ┃                       | |_/ / __ ___   ____ _| |_ ___                        ┃
 *    ┃                       |  __/ '__| \ \ / / _` | __/ _ \                       ┃
 *    ┃                       | |  | |  | |\ V / (_| | ||  __/                       ┃
 *    ┃                       \_|  |_|  |_| \_/ \__,_|\__\___|                       ┃
 *    ┃                                                                              ┃
 *    ┃                                                                              ┃
 *    ┃                                                                              ┃
 *    ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
 */

// getErrorExitCode returns an integer value representing the exit code status for non-zero exit code responses from
// the executable called in the public functions in the subprocess package
func getErrorExitCode(err error) int {
	// fail, non-zero exit status conditions
	if exitError, ok := err.(*exec.ExitError); ok {
		return exitError.Sys().(syscall.WaitStatus).ExitStatus()
	}
	// fails that do not define an exec.ExitError (e.g. unable to identify executable on system PATH)
	return 1 // assign a default non-zero fail code value of 1
}
