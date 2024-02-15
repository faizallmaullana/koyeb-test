package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetAllMenu(c *gin.Context) {
	// Ensure the database is initialized before proceeding
	db, err := getDB()
	if err != nil {
		// Handle error (e.g., return an error response to the client)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close() // Close the database connection when done

	// Query for eatery menu
	eateryRows, err := db.Query("SELECT * FROM menu_eatery")
	if err != nil {
		fmt.Println("Error executing eatery query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute eatery query"})
		return
	}
	defer eateryRows.Close()

	// Query for coffee menu
	coffeeRows, err := db.Query("SELECT * FROM menu_coffee")
	if err != nil {
		fmt.Println("Error executing cooffee query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute coffee query"})
		return
	}
	defer coffeeRows.Close()

	// Query for eatery non coffee
	nonCoffeeRows, err := db.Query("SELECT * FROM menu_non_coffee")
	if err != nil {
		fmt.Println("Error executing non-coffee query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute non coffee query"})
		return
	}
	defer nonCoffeeRows.Close()

	// Query for special menu
	specialityRows, err := db.Query("SELECT * FROM menu_speciality")
	if err != nil {
		fmt.Println("Error executing speciality query:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute speciality query"})
		return
	}
	defer specialityRows.Close()

	var eateryMenus []gin.H // slice to store eatery menu items
	var coffeeMenus []gin.H
	var nonCoffeeMenus []gin.H
	var specialityMenus []gin.H

	// Iterate through the eatery menu result set
	for eateryRows.Next() {
		var id int
		var menu string
		var price string

		err := eateryRows.Scan(&id, &menu, &price)
		if err != nil {
			fmt.Println("Error scanning eatery row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan eatery row"})
			return
		}

		// Append each eatery menu item to the slice
		eateryMenus = append(eateryMenus, gin.H{
			"id":    id,
			"menu":  menu,
			"price": price,
		})
	}

	// Iterate through the coffee menu result set
	for coffeeRows.Next() {
		var id int
		var menu string
		var hot string
		var ice string

		err := coffeeRows.Scan(&id, &menu, &hot, &ice)
		if err != nil {
			fmt.Println("Error scanning coffee row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan coffee row"})
			return
		}

		// Append each coffee menu item to the slice
		coffeeMenus = append(coffeeMenus, gin.H{
			"id":   id,
			"menu": menu,
			"hot":  hot,
			"ice":  ice,
		})
	}

	// Iterate through the non coffee menu result set
	for nonCoffeeRows.Next() {
		var id int
		var menu string
		var hot string
		var ice string

		err := nonCoffeeRows.Scan(&id, &menu, &hot, &ice)
		if err != nil {
			fmt.Println("Error scanning non coffee row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan non coffee row"})
			return
		}

		// Append each coffee menu item to the slice
		nonCoffeeMenus = append(nonCoffeeMenus, gin.H{
			"id":   id,
			"menu": menu,
			"hot":  hot,
			"ice":  ice,
		})
	}

	// Iterate through the eatery menu result set
	for specialityRows.Next() {
		var id int
		var menu string
		var price string
		var caption string

		err := specialityRows.Scan(&id, &menu, &caption, &price)
		if err != nil {
			fmt.Println("Error scanning eatery row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan speciallity menu row"})
			return
		}

		// Append each eatery menu item to the slice
		specialityMenus = append(specialityMenus, gin.H{
			"id":      id,
			"menu":    menu,
			"caption": caption,
			"price":   price,
		})
	}

	// Check for errors during iteration
	if err := eateryRows.Err(); err != nil {
		fmt.Println("Error iterating eatery rows:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating eatery rows"})
		return
	}
	if err := coffeeRows.Err(); err != nil {
		fmt.Println("Error iterating coffee rows:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating coffee rows"})
		return
	}
	if err := nonCoffeeRows.Err(); err != nil {
		fmt.Println("Error iterating noncoffee rows:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating noncoffee rows"})
		return
	}
	if err := specialityRows.Err(); err != nil {
		fmt.Println("Error iterating coffee rows:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating coffee rows"})
		return
	}

	// Return the JSON response with all menu items for eatery and coffee
	c.JSON(http.StatusOK, gin.H{
		"eatery":     eateryMenus,
		"coffee":     coffeeMenus,
		"non-coffee": nonCoffeeMenus,
		"speciality": specialityMenus,
	})
}
