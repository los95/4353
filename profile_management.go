package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func profile(c *gin.Context) {
	type tempAddress struct {
		username    string  `json: username`
		profileInfo Address `json: profileInfo`
	}

	var userInfo tempAddress

	c.BindJSON(&userInfo)

	if err := UpdateUserInfo(userInfo.username, userInfo.profileInfo); err != nil {
		c.IndentedJSON(http.StatusAccepted, userInfo)
	} else {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
}
