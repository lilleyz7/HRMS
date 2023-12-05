package types

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Salary    int    `json:"salary"`
	SickDays  int    `gorm:"default:0"`
	PaidDays  int    `gorm:"default:0"`
	ClockedIn bool   `gorm:"default:false"`
}
