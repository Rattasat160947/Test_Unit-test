package entity

import "gorm.io/gorm"

type Dog struct {
	gorm.Model
	Name string `valid:"required~Name is required"`
	Weight float64 `valid:"required~Weight is required,range(1|20)~Weight must be between 1 and 20"`
}