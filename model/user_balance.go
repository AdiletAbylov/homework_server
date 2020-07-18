package model

import "github.com/jinzhu/gorm"


// UserBalanceModel is an ORM model for user balance table
type UserBalanceModel struct {
	gorm.Model
	UserID   uint `gorm:"AUTO_INCREMENT"`
	Username string
	Balance  float64 `gorm:"DEFAULT:0"`
}
