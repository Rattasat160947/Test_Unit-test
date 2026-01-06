package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `valid:"required~Name is required"`
	Price float64 `valid:"required~Price is required,range(1|10000)~Price must between 1 and 10000"`
	SKU string `valid:"matches(^P\\d{4}$)~SKU invalid format"`
}
