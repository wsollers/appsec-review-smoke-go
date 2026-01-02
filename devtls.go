// Development-only HTTP client. This repository is a deliberately vulnerable smoke target for
// appsec-review; see README.md.
package main

import (
	"crypto/tls"
	"net/http"
)

// insecureClient returns a client for the local development proxy, which presents a self-signed
// certificate. Nothing calls it.
func insecureClient() *http.Client {
	return &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
}
