package models

type Users struct {
	ID       string `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
}
