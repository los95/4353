package main

import (
	"github.com/gin-gonic/gin"
)

func getUsers() []User {
	return users
}

func main() {
	router := gin.Default()
	router.POST("/register", register)
	router.POST("/login", login)
	router.POST("/profile", profile)
	router.GET("/getHistory", history)
	router.Run("localhost:8080")
}
