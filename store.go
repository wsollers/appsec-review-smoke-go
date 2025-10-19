// Storage helpers of the smoke server. This repository is a deliberately vulnerable smoke target for
// appsec-review; see README.md.
package main

import (
	"database/sql"
	"errors"
	"strconv"
)

func openStore() (*sql.DB, error) {
	return sql.Open("sqlite", "app.db")
}

// recentUsers lists the n newest users; n must lie in [0, 100).
func recentUsers(db *sql.DB, n int) error {
	if n >= 0 && n < 100 {
		_, err := db.Exec("SELECT id FROM users ORDER BY created DESC LIMIT " + strconv.Itoa(n))
		return err
	}
	return errors.New("limit out of range")
}

// userByID loads one user; the id must parse as an integer.
func userByID(db *sql.DB, id string) error {
	n, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = db.Exec("SELECT name FROM users WHERE id = " + strconv.Itoa(n))
	return err
}
