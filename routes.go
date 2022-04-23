package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserInput struct {
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	Password string  `json:"password" binding:"required"`
}

func handleRequest() {
	router := gin.Default()
	router.POST("/u/sign-up", signUp)
	router.Run("localhost:8080")
}

func signUp(c *gin.Context)  {
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
	if err := ConnectToDatabase().Create(&user).Error; err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}