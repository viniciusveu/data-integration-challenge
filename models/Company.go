package models

import "gorm.io/gorm"
type Company struct {
	gorm.Model
	Name string `json:Name`
	Zip string `json:Zip`
	Site string `json:Site`
}