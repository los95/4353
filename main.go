package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FuelQuote struct {
	Username string `json: "username"`
	Gallons  int    `json: "gallons"`
}

type Address struct {
	Street  string `json: "street"`
	City    string `json: "city"`
	State   string `json: "state"`
	Zipcode string `json: "zipcode"`
}

type User struct {
	Username  string  `json: "username"`
	FirstName string  `json: "firstName"`
	LastName  string  `json: "lastName"`
	Residence Address `json: "address"`
}

type Credentials struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

var userCredentials = []Credentials{
	{Username: "test user", Password: "test password"},
}

var users = []User{
	{Username: "test user", FirstName: "This", LastName: "Mf", Residence: Address{"123 Cougar Drive", "Houston", "Texas", "77070"}},
}

var fuelQuoteHistory = []FuelQuote{
	{Username: "test user", Gallons: 2},
}

func getUsers() []User {
	return users
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

	registerNewUser(newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

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

func main() {
	router := gin.Default()
	router.POST("/register", register)
	router.POST("/login", login)
	router.POST("/profile", profile)
	router.GET("/getHistory", history)
	router.Run("localhost:8080")
}
