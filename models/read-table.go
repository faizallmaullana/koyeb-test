package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadTable(c *gin.Context) {
	// Query to get all tables
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query"})
		return
	}
	defer rows.Close() // Ensure the rows are closed after using them

	// Iterate over the rows to print table names
	tables := make([]string, 0)
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row"})
			return
		}
		tables = append(tables, tableName)
	}

	c.JSON(http.StatusOK, gin.H{"tables": tables})
}
