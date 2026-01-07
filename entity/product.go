package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title  string `valid:"required~Title must be greater than 3,stringlength(3|101)~Title must be greater than 3"`
	Price float64 `valid:"required~Price must be greater than 0,range(50|5000)~Price must be greater than 0"`
	Code string `valid:"required~Code is required"`
	// Code string `valid:"required~Code is required,range(^[BK][0-9][6]$)~Code is in BK and 6 Number"`
}

func (s *Product) Validate() error {
	_, err := govalidator.ValidateStruct(s)
	return err
}

