package entity

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string  `valid:"required~Name is required"`
	Price    float64 `valid:"required~Price is required,range(1|100000)~Price must be between 1.00 and 100000.00"`
	SKU      string  `valid:"required~SKU is required,matches(^[S][K][U]\\d{5}$)~SKU is invalid"`
	Quantity int     `valid:"required~Quantity is required,range(1|500)~Quantity must be between 1 and 500"`
}
