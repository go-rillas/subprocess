## Changelog

### v1.0.0

- first production release
- API is now stable
- SemVer versioning in place for all future backwards incompatible changes

### v0.3.0

- refactored `RunShell` function for support of `cmd.exe` as the default shell on Windows (changed from bash default)
- refactored default shell flag to `/C` on Windows platform = `cmd.exe /C` call

### v0.2.0

- added support for execution of system executables via shells with `RunShell` function
- refactored `Run` function
- expanded source documentation

### v0.1.2

- minor docstring updates in subprocess.go source file

### v0.1.1

- added docstrings

### v0.1.0

- initial release with public `Run` function