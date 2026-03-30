// Operator tools of the smoke server. This repository is a deliberately vulnerable smoke target for
// appsec-review; see README.md.
package main

import (
	"errors"
	"log"
	"os"
	"runtime"
)

const assetRoot = "/var/lib/smoke/assets"

// runTool performs one of the two maintenance tools; any other name is refused.
func runTool(name string) error {
	switch name {
	case "status":
		log.Printf("status ok")
		return nil
	case "version":
		log.Printf("smoke %s", runtime.Version())
		return nil
	}
	return errors.New("unknown tool")
}

// readAsset returns the asset named by the operator.
func readAsset(name string) ([]byte, error) {
	return os.ReadFile(assetRoot + "/" + name)
}
