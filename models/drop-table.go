package models

import (
	"database/sql"
	"fmt"
)

func DropTable() {
	connStr := "user='tanjung' password=APWa4n3XiTfx host=ep-green-wave-a1c4bn3h.ap-southeast-1.pg.koyeb.app dbname='koyebdb'"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close() // Ensure the database connection is closed when main function exits

	// Get a list of all tables in the database
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		panic(err)
	}

	defer rows.Close() // Ensure the rows are closed after using them

	// Iterate over the rows and drop each table
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			panic(err)
		}

		dropTableQuery := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tableName)
		_, err := db.Exec(dropTableQuery)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Table '%s' dropped successfully!\n", tableName)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		panic(err)
	}

	fmt.Println("All tables dropped successfully!")
}
