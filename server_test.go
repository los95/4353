package main

import (
	"errors"
	"fmt"
	"testing"
)

func findHistory(username string) (int, error) {
	for _, user := range fuelQuoteHistory {
		if user.Username == username {
			return user.Gallons, nil
		}
	}

	return 0, errors.New("no history found")
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

func TestGetUsers(t *testing.T) {
	actualResult := getUsers()
	expectedResult := []User{
		{Username: "test user", FirstName: "This", LastName: "Mf", Residence: Address{"123 Cougar Drive", "Houston", "Texas", "77070"}},
	}

	for i, user := range actualResult {
		if expectedResult[i] != user {
			t.Errorf("Arrays do not match")
		}
	}
}

func TestRegistration(t *testing.T) {
	newUser := Credentials{
		Username: "testUserName",
		Password: "testPassword",
	}

	registerNewUser(newUser)

	if len(users) != 2 && len(userCredentials) != 2 {
		t.Errorf("User Length and/or Credential Lengths do not match")
	}
}

func TestRegistrationDuplicateUser(t *testing.T) {
	newUser := Credentials{
		Username: "test user",
		Password: "testPassword",
	}

	if !userExists(newUser.Username) {
		t.Errorf("User already exists but still registered")
	}
}

func TestLoginWrongUser(t *testing.T) {
	_, err := getUser(Credentials{"wrong username", "some password"})

	fmt.Println(err)

	if err == errors.New("User not found") {
		t.Errorf("Still logged in ")
	}
}

func TestLoginWrongPassword(t *testing.T) {
	_, err := getUser(Credentials{"test user", "some password"})

	if err == errors.New("incorrect password") {
		t.Errorf("Still logged in ")
	}
}

func TestLoginValid(t *testing.T) {
	result, _ := getUser(Credentials{"test user", "test password"})
	expectedResult := User{"test user", "This", "Mf", Address{"123 Cougar Drive", "Houston", "Texas", "77070"}}

	if *result != expectedResult {
		t.Errorf("Did not log in")
	}
}

func TestProfileEditingValid(t *testing.T) {
	result := editProfile("test user", Address{"", "", "", ""})

	if result != nil {
		t.Errorf("Did not edit profile")
	}
}

func TestProfileEditingInvalid(t *testing.T) {
	result := editProfile("wrong username", Address{"", "", "", ""})

	if result == errors.New("unable to find user") {
		t.Errorf("Somehow found user")
	}
}
