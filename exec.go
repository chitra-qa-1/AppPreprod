package main

import (
	"os/exec"
	"strings"
)

// RunCommand shows execution via the shell with untrusted data
func RunCommand(cmd string) (string, error) {
	// VULNERABLE: using /bin/sh -c with user-supplied content allows command injection
	// Klocwork and similar tools flag patterns that call exec.Command with shell concatenation.
	c := exec.Command("/bin/sh", "-c", cmd)
	out, err := c.CombinedOutput()
	return string(out), err
}
