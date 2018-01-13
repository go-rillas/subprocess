// Package subprocess provides support for standard output/error pipe data & exit status codes with new spawned system processes
package subprocess

import (
	"bytes"
	"os/exec"
	"syscall"
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

// Run is a public function that executes a system command and returns the standard output stream,
// standard error stream, and exit status code data in a returned subprocess.Response struct.
// Run takes the following parameters:
//
//  executable (string) - the system executable for the command call
//  args (...string) - comma delimited list of arguments to executable
//
// Example:
//
// func main() {
//     response := Run("go", "--help")
//     fmt.Printf("%s\n", response.StdOut)
//     fmt.Printf("%s\n", response.StdErr)
//     fmt.Printf("%d\n", response.ExitCode)
// }
func Run(executable string, args ...string) Response {
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
		// fail, non-zero exit status conditions
		if exitError, ok := err.(*exec.ExitError); ok {
			res.ExitCode = exitError.Sys().(syscall.WaitStatus).ExitStatus()
		} else {
			// fails that do not define an exec.ExitError (e.g. unable to identify executable on system PATH)
			res.ExitCode = 1 // assign a default non-zero fail code value of 1
			if res.StdErr == "" {
				res.StdErr = err.Error() // return the error raised in standard error stream formatted as a string
			}
		}
	} else {
		// success, zero exit status condition
		res.ExitCode = cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	}

	return res
}

//func Pipe() {
//	// TODO
//}
