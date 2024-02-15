package controller

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	db          *sql.DB
	dbInitError error
)

func getDB() (*sql.DB, error) {
	connStr := "user=tanjung password=APWa4n3XiTfx host=ep-green-wave-a1c4bn3h.ap-southeast-1.pg.koyeb.app dbname=koyebdb sslmode=require"
	db, dbInitError = sql.Open("postgres", connStr)
	if dbInitError != nil {
		fmt.Println("Error connecting to database:", dbInitError)
		return nil, dbInitError
	}
	return db, dbInitError
}
