package models

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadTable(c *gin.Context) {
	connStr := "user='tanjung' password=APWa4n3XiTfx host=ep-green-wave-a1c4bn3h.ap-southeast-1.pg.koyeb.app dbname='koyebdb'"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close() // Ensure the database connection is closed when main function exits

	// Query to get all tables
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		panic(err)
	}

	defer rows.Close() // Ensure the rows are closed after using them

	// Iterate over the rows to print table names
	fmt.Println("Tables:")
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			panic(err)
		}
		fmt.Println(tableName)
		c.JSON(http.StatusOK, gin.H{"msg": tableName})
	}
}
