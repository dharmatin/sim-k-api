package model

import "github.com/jinzhu/gorm"

type Seller struct {
	gorm.Model
	name    string
	address string
}
