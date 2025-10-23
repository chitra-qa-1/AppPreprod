package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// SearchUsers demonstrates SQL query construction with fmt.Sprintf (vulnerable)
func SearchUsers(q string) ([]string, error) {
	// Using a file-backed sqlite DB for demo only
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// create sample table
	_, _ = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
	_, _ = db.Exec("INSERT INTO users (name) VALUES ('alice'), ('bob')")

	// VULNERABLE: building SQL with Sprintf and untrusted input
	// Static analysers like Klocwork will flag this as SQL injection risk.
	query := fmt.Sprintf("SELECT name FROM users WHERE name LIKE '%%%s%%';", q)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var names []string
	for rows.Next() {
		var n string
		if err := rows.Scan(&n); err == nil {
			names = append(names, n)
		}
	}
	return names, nil
}
