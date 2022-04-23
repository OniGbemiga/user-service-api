package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserInput struct {
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
	Password string  `json:"password" binding:"required"`
}

var user User

func handleRequest() {
	router := gin.Default()
	router.POST("/u/sign-up", signUp)
	router.POST("/u/sign-in", signIn)
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

	//err := connectToRedis().Set("verification_token","sgd8d0", time.Minute * 2).Err()
	//
	//if err != nil {
	//	errors.New("unable to save to redis")
	//	fmt.Println(err)
	//	return
	//}
	//
	//getToken, err := connectToRedis().Get("verification_token").Result()
	//
	//if err != nil {
	//	errors.New("unable to get verification token")
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Printf("Verification Token: %v \n", getToken)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func signIn(c *gin.Context)  {
	var useriinput UserInput
	if err := c.ShouldBindJSON(&useriinput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inputs := User{
		Email: useriinput.Email,
		Phone: useriinput.Phone,
		Password: useriinput.Password,
	}

	fmt.Printf("Email %v\n", inputs.Email)
	if inputs.Email != nil {
		if err := ConnectToDatabase().Where("email = ? AND password = ?",inputs.Email,inputs.Password).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Email or Password!"})
			return
		}
	} else {
		if err := ConnectToDatabase().Where("phone = ? AND password = ?",inputs.Phone,inputs.Password).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Email or Password!"})
			return
		}
	}
	fmt.Println(user)

	c.JSON(http.StatusOK, gin.H{"data": &user})
}