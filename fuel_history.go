package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHistory(username string) (int, error) {
	for _, user := range fuelQuoteHistory {
		if user.Username == username {
			return user.Gallons, nil
		}
	}

	return 0, errors.New("no history found")
}

func history(c *gin.Context) {
	username := c.Param("username")

	if history, err := getHistory(username); err != nil {
		c.IndentedJSON(http.StatusAccepted, history)
	} else {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
}
