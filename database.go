package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type User struct {
	gorm.Model
	Username *string `gorm:"unique" json:"username"`
	Email    *string `gorm:"unique" json:"email"`
	Phone    *string `gorm:"unique" json:"phone"`
	Password string  `gorm:"null" json:"password"`
}

func ConnectToDatabase() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/user-system-test?charset=utf8mb4&parseTime=True&loc=Local"
	dbs, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection to database failed")
	} else {
		fmt.Println("Connection Established")
	}

	errs := dbs.AutoMigrate(&User{})
	if errs != nil {
		errors.New("unable to migrate User database")
	}

	return dbs

}