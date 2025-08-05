// Storage helpers of the smoke server. This repository is a deliberately vulnerable smoke target for
// appsec-review; see README.md.
package main

import "database/sql"

func openStore() (*sql.DB, error) {
	return sql.Open("sqlite", "app.db")
}
