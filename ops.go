// Operator tools of the smoke server. This repository is a deliberately vulnerable smoke target for
// appsec-review; see README.md.
package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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

// readAsset returns the asset named by the operator. The joined path is cleaned and must
// stay under the asset root.
func readAsset(name string) ([]byte, error) {
	path := filepath.Clean(filepath.Join(assetRoot, name))
	if !strings.HasPrefix(path, assetRoot+"/") {
		return nil, errors.New("asset outside root")
	}
	return os.ReadFile(path)
}
