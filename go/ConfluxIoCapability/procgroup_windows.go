//go:build windows

// Windows has no POSIX process groups: platformSetpgid is a no-op and
// platformKillGroup returns false so callers fall back to killing the process
// directly. Host-effect glue, not part of the verified core.
package ConfluxIoCapability

import "os/exec"

func platformSetpgid(c *exec.Cmd) {}

func platformKillGroup(pid int) bool { return false }
