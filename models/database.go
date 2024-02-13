package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import PostgreSQL driver
)

func ConnectToDatabase() {
	connStr := "user='tanjung' password=APWa4n3XiTfx host=ep-green-wave-a1c4bn3h.ap-southeast-1.pg.koyeb.app dbname='koyebdb'"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close() // Ensure the database connection is closed when main function exits

	// Create a table
	createTableQuery := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100),
            age INT
        )
    `
	_, err = db.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table 'users' created successfully!")
}
