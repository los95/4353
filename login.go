package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUser(credentials Credentials) (*User, error) {
	for i, user := range userCredentials {
		if user.Username == credentials.Username {
			if user.Password == credentials.Password {
				return &users[i], nil
			} else {
				return nil, errors.New("incorrect password")
			}
		}
	}

	return nil, errors.New("User not found")
}

func login(c *gin.Context) {
	var credentials Credentials

	if err := c.BindJSON(&credentials); err != nil {
		return
	}

	user, err := getUser(credentials)

	if err != nil {
		c.IndentedJSON(http.StatusAccepted, user)
	} else {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
}
