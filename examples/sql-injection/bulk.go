package injection

import (
	"database/sql"
	"fmt"
	"strings"
)

func insertUsers(db *sql.DB, names []string) {
	var values []string
	for _, name := range names {
		values = append(values, fmt.Sprintf("('%s')", name))
	}
	query := fmt.Sprintf(
		"INSERT INTO users (name) VALUES %s",
		strings.Join(values, ","),
	)
	db.Query(query)
}

func insertUsersSafe(db *sql.DB, names []string) {
	// How would you write it? ;)
}
