package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func mongoApi(method string, url string, body []byte) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")
	req.Header.Add("api-key", "Vp3ySCFvqFlQgwmxcEnzwFmHP1GuPqJMVOSCMfgM05tn6NwqX6HlewJYsziTgfFE")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}
	fmt.Println(string(resBody))
	return resBody, err
}

func RegistrationQuery(username string, password string) ([]byte, error) {
	registrationCheckerPayload := []byte(fmt.Sprintf(`{
		"collection":"Entry",
		"database":"SoftwareDesign",
		"dataSource":"Cluster0",
		"filter": {"username": "%s"}
	}`, username))

	url := "https://data.mongodb-api.com/app/data-iwttg/endpoint/data/beta/action/findOne"
	res, err := mongoApi("POST", url, registrationCheckerPayload)

	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	fmt.Println(string(res))

	//convert res to json
	var result map[string]interface{}
	var result2 map[string]interface{}
	var result3 map[string]interface{}

	err = json.Unmarshal(res, &result)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	fmt.Println(result)

	if result["document"] != nil {
		return []byte("User already exists"), fmt.Errorf("Registration failed")
	} else {
		fmt.Println("Document is null")
	}

	userInfoPayload := []byte(fmt.Sprintf(`{
		"collection":"UserInfo",
		"database":"SoftwareDesign",
		"dataSource":"Cluster0",
		"document": {
			"username": "%s",
			"Fullname": "1",
			"Address1": "2",
			"Address2": "3",
			"City":     "4",
			"State":    "5",
			"Zipcode":  "6"
		}
	}`, username))

	url = "https://data.mongodb-api.com/app/data-iwttg/endpoint/data/beta/action/insertOne"
	res, err = mongoApi("POST", url, userInfoPayload)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	err = json.Unmarshal(res, &result2)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	fuelHistoryPayload := []byte(`{
		"collection":"History",
		"database":"SoftwareDesign",
		"dataSource":"Cluster0",
		"document": {
			"FuelHistory": []
		}
	}`)

	url = "https://data.mongodb-api.com/app/data-iwttg/endpoint/data/beta/action/insertOne"
	res, err = mongoApi("POST", url, fuelHistoryPayload)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	err = json.Unmarshal(res, &result3)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	registrationPayload := []byte(fmt.Sprintf(`{
		"collection":"Entry",
		"database":"SoftwareDesign",
		"dataSource":"Cluster0",
		"document": {
			"username": "%s",
			"password": "%s",
			"profile": "%s",
			"history": "%s"
		}
	}`, username, password, result2["insertedId"], result3["insertedId"]))

	url = "https://data.mongodb-api.com/app/data-iwttg/endpoint/data/beta/action/insertOne"
	res, err = mongoApi("POST", url, registrationPayload)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	return res, err
}

func LoginQuery(username string, password string) ([]byte, error) {
	loginPayload := []byte(fmt.Sprintf(`{
		"collection":"Entry",
		"database":"SoftwareDesign",
		"dataSource":"Cluster0",
		"filter": {"username": "%s", "password": "%s"}
	}`, username, password))

	url := "https://data.mongodb-api.com/app/data-iwttg/endpoint/data/beta/action/findOne"
	var result map[string]interface{}
	res, err := mongoApi("POST", url, loginPayload)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}
	if result["document"] == nil {
		return []byte(""), fmt.Errorf("Login failed")
	}

	return res, err
}

func ProfileGetterQuery(username string) ([]byte, error) {
	profilePayload := []byte(fmt.Sprintf(`{
		"collection":"UserInfo",
		"database":"SoftwareDesign",
		"dataSource":"Cluster0",
		"filter": {"username": "%s"}
	}`, username))

	url := "https://data.mongodb-api.com/app/data-iwttg/endpoint/data/beta/action/findOne"
	var result map[string]interface{}
	res, err := mongoApi("POST", url, profilePayload)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	return res, err
}

func fuelQuoteInserter(username string, fuelPrice string, fuelQuantity string) ([]byte, error) {
	fuelQuotePayload := []byte(fmt.Sprintf(`{
		"collection":"History",
		"database":"SoftwareDesign",
		"dataSource":"Cluster0",
		"document": {"username": "%s",
			"price": "%s","
			"quantity": "%s"
		}
	}`, username, fuelPrice, fuelQuantity))

	url := "https://data.mongodb-api.com/app/data-iwttg/endpoint/data/beta/action/insertOne"

	res, err := mongoApi("POST", url, fuelQuotePayload)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	return res, err
}

func fuelQuoteGetter(username string) ([]byte, error) {
	fuelQuotePayload := []byte(fmt.Sprintf(`{
		"collection":"History",
		"database":"SoftwareDesign",
		"dataSource":"Cluster0",
		"filter": {"username": "%s"}
	}`, username))

	url := "https://data.mongodb-api.com/app/data-iwttg/endpoint/data/beta/action/findAll"

	var result map[string]interface{}
	res, err := mongoApi("POST", url, fuelQuotePayload)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		fmt.Println(err)
		return []byte(""), err
	}

	return res, err
}
