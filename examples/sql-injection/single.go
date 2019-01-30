package injection

import (
	"database/sql"
	"fmt"
)

func findUser(db *sql.DB, name string) {
	query := fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", name) // HL
	db.Query(query)
}

func findUserSafe(db *sql.DB, name string) {
	query := "SELECT * FROM users WHERE name = $1" // HL
	db.Query(query, name)
}
