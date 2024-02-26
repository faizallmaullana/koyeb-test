package models

// authentication

type Authentication struct {
	ID       string `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`

	// foreign keys
	StaffID string `json:"staff_id"`
	Staff   Staff  `json:"staff" gorm:"references:StaffID"`
}

type Token struct {
	ID    string `json:"id" gorm:"primary_key"`
	Token string `json:"token"`
}
