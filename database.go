package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func connect(payload *strings.Reader, requestType string) string {
	url := fmt.Sprintf("https://data.mongodb-api.com/app/data-iwttg/endpoint/data/beta/action/%s", requestType)
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		//fmt.Println(err)
		return err.Error()
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")
	req.Header.Add("api-key", API_KEY)

	res, err := client.Do(req)
	if err != nil {
		//fmt.Println(err)
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		//fmt.Println(err)
		return err.Error()
	}
	return string(body)
}

func FindUser(username string) bool {
	payload := strings.NewReader(fmt.Sprintf(`{
        "collection":"Entry",
        "database":"SoftwareDesign",
        "dataSource":"Cluster0",
        "projection": {"username": "%s"},
    }`, username))

	res := connect(payload, "findOne")

	if res == `{"error":"Cannot find document"` {
		return true
	} else {
		return false
	}
}

func ValidatePassword(username string, password string) (bool, string) {
	payload := strings.NewReader(fmt.Sprintf(`{
        "collection":"Entry",
        "database":"SoftwareDesign",
        "dataSource":"Cluster0",
        "projection": {"username": "%s",
						"password": "%s"},
    }`, username, password))

	res := connect(payload, "findOne")

	if res == `{"error":"Cannot find document"` {
		return false, nil
	} else {
		return true, res
	}
}

func FindHistory(historyId string) string {
	payload := strings.NewReader(fmt.Sprintf(`{
        "collection":"History",
        "database":"SoftwareDesign",
        "dataSource":"Cluster0",
        "projection": %s)`, historyId))

	return connect(payload, "findOne")
}

func FindUserInfo(userId string) string {
	payload := strings.NewReader(fmt.Sprintf(`{
        "collection":"UserInfo",
        "database":"SoftwareDesign",
        "dataSource":"Cluster0",
        "projection": %s)`, userId))

	return connect(payload, "findOne")
}

func UpdateUserInfo(userId string, firstName string, lastName string, address string, state string) string {
	payload := strings.NewReader(fmt.Sprintf(`{
        "collection":"UserInfo",
        "database":"SoftwareDesign",
        "dataSource":"Cluster0",
        "projection": %s,
		"update" : {
			"$set": {
				"firstName" : %s,
				"lastName" : %s,
				"address" : %s,
				"state" : %s
			}
		} `, userId, firstName, lastName, address, state))

	return connect(payload, "updateOne")
}

func updateHistory(historyId string, oldHistory string, newEntry string) string {
	payload := strings.NewReader(fmt.Sprintf(`{
        "collection":"UserInfo",
        "database":"SoftwareDesign",
        "dataSource":"Cluster0",
        "projection": %s,
		"update" : {
			"$set": {
				"history" : %s %s
			}
		} `, historyId, oldHistory, newEntry))

	return connect(payload, "updateOne")
}
