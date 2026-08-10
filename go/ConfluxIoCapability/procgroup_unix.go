//go:build !windows

// Unix process-group control for the ConfluxIoCapability host boundary: the child
// leads its own group so the whole tree can be signalled at once. Host-effect glue,
// not part of the verified core.
package ConfluxIoCapability

import (
	"os/exec"
	"syscall"
)

func platformSetpgid(c *exec.Cmd) {
	c.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

func platformKillGroup(pid int) bool {
	pgid, err := syscall.Getpgid(pid)
	if err != nil {
		return false
	}
	return syscall.Kill(-pgid, syscall.SIGKILL) == nil
}
