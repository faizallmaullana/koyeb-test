package models

import (
	"database/sql"

	_ "github.com/lib/pq" // Import PostgreSQL driver
)

var db *sql.DB // Global variable to store the database connection

// InitializeDatabase initializes the database connection
func ConnectToDatabase() error {
	connStr := "user='tanjung' password=APWa4n3XiTfx host=ep-green-wave-a1c4bn3h.ap-southeast-1.pg.koyeb.app dbname='koyebdb'"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	return nil
}
