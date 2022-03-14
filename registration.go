package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func userExists(newUsername string) bool {
	for _, user := range userCredentials {
		if newUsername == user.Username {
			return true
		}
	}
	return false
}

func registerNewUser(newUser Credentials) {
	userCredentials = append(userCredentials, newUser)
	users = append(users, User{newUser.Username, "", "", Address{}})
}

func register(c *gin.Context) {
	var newUser Credentials

	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	if userExists(newUser.Username) {
		c.IndentedJSON(http.StatusBadRequest, errors.New("user already exists"))
	} else {
		registerNewUser(newUser)
		c.IndentedJSON(http.StatusCreated, newUser)
	}
}
