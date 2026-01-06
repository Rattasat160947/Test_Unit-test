package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Username string  `valid:"stringlength(5|20)~Username must be between 5 and 20 characters"`
	Email string `valid:"email~Invalid email format"`
	Gender string 	`valid:"in(Male|Female)~Gender must be Male or Female"`
	Age uint `valid:"range(18|99)~Age must be between 18 and 99"`
}