package models

import (
	"fmt"
)

var tableCreationQueries = []string{
	`
			CREATE TABLE IF NOT EXISTS users (
					id SERIAL PRIMARY KEY,
					username VARCHAR(255),
					password VARCHAR(255)
			)
	`,
	`
			CREATE TABLE IF NOT EXISTS suggestion (
					id SERIAL PRIMARY KEY,
					suggester VARCHAR(255),
					email VARCHAR(255),
					suggest VARCHAR(255)
			)
	`,
	`
			CREATE TABLE IF NOT EXISTS menu_speciality (
				id SERIAL PRIMARY KEY,
				menu VARCHAR(100),
				caption VARCHAR(255),
				price VARCHAR(5)
			)
	`,
	`
			CREATE TABLE IF NOT EXISTS menu_eatery (
				id SERIAL PRIMARY KEY,
				menu VARCHAR(100),
				price VARCHAR(5)
			)
	`,
	`
			CREATE TABLE IF NOT EXISTS menu_coffee (
				id SERIAL PRIMARY KEY,
				menu VARCHAR(100),
				hot VARCHAR(5),
				ice VARCHAR(5)
			)
	`,
	`
			CREATE TABLE IF NOT EXISTS menu_non_coffee (
				id SERIAL PRIMARY KEY,
				menu VARCHAR(100),
				hot VARCHAR(5),
				ice VARCHAR(5)
			)
	`,
}

func CreateTable() error {
	for _, query := range tableCreationQueries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	fmt.Println("Tables created successfully!")
	return nil
}
