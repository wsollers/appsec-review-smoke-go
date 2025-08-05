// Admin console of the smoke server. This repository is a deliberately vulnerable smoke target for
// appsec-review; see README.md.
package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/admin", adminHandler)
}

// adminHandler serves /admin: a small operator console driven by the op parameter.
func adminHandler(w http.ResponseWriter, r *http.Request) {
	op := r.FormValue("op")
	db, err := openStore()
	if err != nil {
		http.Error(w, "store unavailable", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	status := http.StatusOK
	switch op {
	case "ping":
		fmt.Fprintln(w, "pong")
	default:
		status = http.StatusNotFound
	}
	if status != http.StatusOK {
		w.WriteHeader(status)
	}
}
