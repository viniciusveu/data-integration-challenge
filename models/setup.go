package models

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB  *gorm.DB

// Conecta ao banco de dados api_challenge e configura o gorm
func ConnectDatabase() {

	dsn := "root:root@tcp(172.17.0.2:3306)/api_challenge?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	} 

	database.AutoMigrate(&Company{})

	DB = database

	fmt.Println("Connected to the database")
}