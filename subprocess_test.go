package subprocess

import (
	"testing"
)

func TestRunValidCommand(t *testing.T) {
	response := Run("git", "--help")
	if response.ExitCode != 0 {
		t.Errorf("[FAIL] Expected 'git --help' to return 0 exit status code and instead it returned %d", response.ExitCode)
	}
	if len(response.StdErr) > 0 {
		t.Errorf("[FAIL] Expected 'git --help' to return no standard error output and instead it returned %v", response.StdErr)
	}
	if len(response.StdOut) == 0 {
		t.Errorf("[FAIL] Expected 'git --help' to return standard output but instead it was empty.")
	}
}


func TestRunInValidCommandBadArgument(t *testing.T) {
	response := Run("git", "--bogus")
	if response.ExitCode == 0 {
		t.Errorf("[FAIL] Expected invalid argument to executable to return non-0 exit status code and instead it returned %d", response.ExitCode)
	}
	if len(response.StdErr) == 0 {
		t.Errorf("[FAIL] Expected invalid argument to executable to return standard error output and instead it returned an empty string")
	}
	if len(response.StdOut) > 0 {
		t.Errorf("[FAIL] Expected invalid argument to return no standard output but instead it returned %s.", response.StdOut)
	}
}

func TestRunInvalidCommandMissingExecutable(t *testing.T) {
	response := Run("bogus", "--help")
	if response.ExitCode == 0 {
		t.Errorf("[FAIL] Expected invalid command to return non-0 exit status code and instead it returned %d", response.ExitCode)
	}
	if len(response.StdErr) == 0 {
		t.Errorf("[FAIL] Expected invalid command to return standard error output and instead it returned an empty string")
	}
	if len(response.StdOut) > 0 {
		t.Errorf("[FAIL] Expected invalid command to return no standard output but instead it returned %s.", response.StdOut)
	}
}
