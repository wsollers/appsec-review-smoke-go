// Operator tools of the smoke server. This repository is a deliberately vulnerable smoke target for
// appsec-review; see README.md.
package main

import "os/exec"

// runTool runs a maintenance helper named by the operator.
func runTool(name string) error {
	return exec.Command("/bin/sh", "-c", "/usr/libexec/smoke/"+name).Run()
}
