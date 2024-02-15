package jsonData

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

type Coffee struct {
	Menu string `json:"menu"`
	Hot  string `json:"hot"`
	Ice  string `json:"ice"`
}

type Special struct {
	Menu    string `json:"menu"`
	Price   string `json:"price"`
	Caption string `json:"caption"`
}

type Eatery struct {
	Menu  string `json:"menu"`
	Price string `json:"price"`
}

func ReadJson(c *gin.Context) {
	connStr := "user=tanjung password=APWa4n3XiTfx host=ep-green-wave-a1c4bn3h.ap-southeast-1.pg.koyeb.app dbname=koyebdb sslmode=require"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close() // Close the database connection when done

	// Buka file JSON Cofee base
	jsonFile, err := os.Open("./jsonData/coffee-base.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// Baca isi file JSON================
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// dekoding
	var coffees []Coffee
	err = json.Unmarshal(byteValue, &coffees)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for _, coffee := range coffees {
		_, err := db.Exec("INSERT INTO menu_coffee (menu, hot, ice) VALUES ($1, $2, $3)", coffee.Menu, coffee.Hot, coffee.Ice)
		if err != nil {
			fmt.Println("Error inserting into database:", err)
			return
		}
		fmt.Println("Read Done")
	}

	// Buka file JSON Cofee base
	jsonFile, err = os.Open("./jsonData/non-coffee.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// Baca isi file JSON=============
	byteValue, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// dekoding
	// var coffees []Coffee
	err = json.Unmarshal(byteValue, &coffees)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for _, coffee := range coffees {
		_, err := db.Exec("INSERT INTO menu_non_coffee (menu, hot, ice) VALUES ($1, $2, $3)", coffee.Menu, coffee.Hot, coffee.Ice)
		if err != nil {
			fmt.Println("Error inserting into database:", err)
			return
		}
		fmt.Println("Read Done")
	}

	// Buka file JSON Cofee base
	jsonFile, err = os.Open("./jsonData/speciality-es-kopi.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// Baca isi file JSON
	byteValue, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// dekoding
	var specialitys []Special
	err = json.Unmarshal(byteValue, &specialitys)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for _, special := range specialitys {
		_, err := db.Exec("INSERT INTO menu_speciality (menu, caption, price) VALUES ($1, $2, $3)", special.Menu, special.Caption, special.Price)
		if err != nil {
			fmt.Println("Error inserting into database:", err)
			return
		}
		fmt.Println("Read Done")
	}

	// Buka file JSON Cofee base
	jsonFile, err = os.Open("./jsonData/eatery.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// Baca isi file JSON
	byteValue, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// dekoding
	var eaterys []Eatery
	err = json.Unmarshal(byteValue, &eaterys)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for _, eatery := range eaterys {
		_, err := db.Exec("INSERT INTO menu_eatery (menu, price) VALUES ($1, $2)", eatery.Menu, eatery.Price)
		if err != nil {
			fmt.Println("Error inserting into database:", err)
			return
		}
		fmt.Println("Read Done")
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg": "all data created successfully",
	})
}
