package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectToDatabase() {
	pwd, _ := os.Getwd()
	fmt.Println("Current working directory:", pwd)
	database, err := gorm.Open("sqlite3", "zpd.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(
		&Authentication{},
		&Token{},
		&Staff{},
		&Jurusan{},
		&Kelas{},
		&Siswa{},
		&PivotKelasSiswa{},
		&MataPelajaran{},
		&MateriAjar{},
		&CapaianSiswa{},
	)

	DB = database
}
