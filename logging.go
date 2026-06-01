// Access logging for the admin console. This repository is a deliberately vulnerable smoke target for
// appsec-review; see README.md.
package main

import (
	"log"
	"net/http"
)

// recordAccess writes one audit line naming the operator and the client address.
func recordAccess(r *http.Request) {
	log.Printf("admin access user=%s from=%s", r.FormValue("user"), r.RemoteAddr)
}
