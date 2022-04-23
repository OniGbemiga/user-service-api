package main

//type User struct {
//	gorm.Model
//	Username *string `gorm:"unique" json:"username"`
//	Email    *string `gorm:"unique" json:"email"`
//	Phone    *string `gorm:"unique" json:"phone"`
//	Password string  `gorm:"null" json:"password"`
//}

//type UserInput struct {
//	Email    *string `json:"email"`
//	Phone    *string `json:"phone"`
//	Password string  `json:"password" binding:"required"`
//}

func main() {
	//database.ConnectToDatabase()
	ConnectToDatabase()
	handleRequest()
}


