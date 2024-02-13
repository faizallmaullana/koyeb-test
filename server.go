package main

import (
	"fmt"
	"os"

	"github.com/faizallmaullana/test-koyeb/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Import PostgreSQL driver
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	models.ConnectToDatabase()

	r.GET("api/v1/kopisore/ReadTables", models.ReadTable)

	r.Run(fmt.Sprintf(":%s", port))
}

// func main() {
// 	connStr := "user='tanjung' password=APWa4n3XiTfx host=ep-green-wave-a1c4bn3h.ap-southeast-1.pg.koyeb.app dbname='koyebdb'"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer db.Close() // Ensure the database connection is closed when main function exits

// 	// Query to get all tables
// 	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close() // Ensure the rows are closed after using them

// 	// Iterate over the rows to print table names
// 	fmt.Println("Tables:")
// 	for rows.Next() {
// 		var tableName string
// 		if err := rows.Scan(&tableName); err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(tableName)
// 	}

// 	// Check for errors during row iteration
// 	if err := rows.Err(); err != nil {
// 		panic(err)
// 	}
// }
