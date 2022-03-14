package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func editProfile(username string, address Address) error {
	for i, user := range users {
		if user.Username == username {
			users[i].Residence = address
			return nil
		}
	}

	return errors.New("unable to find user")
}

func profile(c *gin.Context) {
	type tempAddress struct {
		username    string  `json: username`
		profileInfo Address `json: profileInfo`
	}

	var userInfo tempAddress

	c.BindJSON(&userInfo)

	if err := editProfile(userInfo.username, userInfo.profileInfo); err != nil {
		c.IndentedJSON(http.StatusAccepted, userInfo)
	} else {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
}
