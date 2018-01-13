package subprocess

import (
	"bytes"
	"os/exec"
	"syscall"
)

type Response struct {
	StdOut   string
	StdErr   string
	ExitCode int
}

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

func Pipe() {
	// TODO
}
