package controller

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db          *sql.DB
	dbInitError error
	dbOnce      sync.Once
)

func getDB() (*sql.DB, error) {
	// Use sync.Once to ensure the initialization is done only once
	dbOnce.Do(func() {
		connStr := "user=tanjung password=APWa4n3XiTfx host=ep-green-wave-a1c4bn3h.ap-southeast-1.pg.koyeb.app dbname=koyebdb sslmode=require"
		db, dbInitError = sql.Open("postgres", connStr)
		if dbInitError != nil {
			fmt.Println("Error connecting to database:", dbInitError)
		}
	})
	return db, dbInitError
}
