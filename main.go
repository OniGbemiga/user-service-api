package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	database "user-service"
)

type User struct {
	gorm.Model
	Username *string `gorm:"unique" json:"username"`
	Email    *string `gorm:"unique" json:"email"`
	Phone    *string `gorm:"unique" json:"phone"`
	Password string  `gorm:"null" json:"password"`
}

type UserInput struct {
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	Password string  `json:"password" binding:"required"`
}

func main() {
	database.ConnectToDatabase()
	handleRequest()
}

func handleRequest() {
	router := gin.Default()
	router.POST("/u/sign-up", func(c *gin.Context) {
		var useriinput UserInput
		if err := c.ShouldBindJSON(&useriinput); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := User{
			Email:    useriinput.Email,
			Phone:    useriinput.Phone,
			Password: useriinput.Password,
		}
		if err := database.ConnectToDatabase().Create(&user).Error; err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, gin.H{"data": user})

	})
	router.Run("localhost:8080")
}
