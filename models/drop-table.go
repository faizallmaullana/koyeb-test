package models

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DropTable(c *gin.Context) {
	// Get a list of all tables in the database
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
	}

	defer rows.Close() // Ensure the rows are closed after using them

	// Iterate over the rows and drop each table
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		}

		dropTableQuery := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", tableName)
		_, err := db.Exec(dropTableQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err})

		}
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
	}

	c.JSON(http.StatusOK, gin.H{"msg": "All tables dropped successfully"})
}
