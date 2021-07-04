package model

import "gorm.io/gorm"

// User ... User Database Model
type User struct {
	gorm.Model
	Email  string  `json:"email" gorm:"unique;not null"`
	Number float64 `json:"wallet"`
}
