package requests

import (
	"bytes"
	"client/clientModel"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func UserRegistration(info clientModel.UserEntry) bool {
	//info := clientModel.UserEntry{
	//	Username: "clientTest",
	//	Password: "password",
	//}
	fmt.Println("Does this run?")
	//jsonValue, _ := json.Marshal(info)
	//response, err := http.Post("http://localhost:8000/register", "application/json", bytes.NewBuffer(jsonValue))
	response, err := RegistrationQuery(info.Username, info.Password)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		//data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(response))
		return true
	}
}

func UserLogin(info clientModel.UserEntry) (clientModel.UserEntryInfo, clientModel.ResponseResult) {
	//info := clientModel.UserEntry{
	//	Username: "clientTest",
	//	Password: "password",
	//}
	fmt.Println("Does this run too?")
	//jsonValue, _ := json.Marshal(info)
	response, err := LoginQuery(info.Username, info.Password)
	if err != nil {
		var error clientModel.ResponseResult
		error.Error = err.Error()
		fmt.Println(err)
		fmt.Println(response)
		fmt.Println("login failed my dude")
		//json.NewDecoder(response).Decode(&error)
		return clientModel.UserEntryInfo{}, error
	} else {
		var userEntry clientModel.UserEntryInfo
		//json.NewDecoder(response.Body).Decode(&userEntry)
		//data, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))
		fmt.Println(userEntry)
		userEntry.Username = info.Username
		return userEntry, clientModel.ResponseResult{}
	}
}

func UserProfileSetter(token string, userInfo clientModel.ProfileInfo) {
	jsonValue, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println(err)
	}
	//var bearer = "Bearer " + token
	request, _ := http.NewRequest("PUT", "http://localhost:8000/profileSetter", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("Authorization", token)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

func UserProfileGetter(username string) clientModel.ProfileInfo {
	// request, _ := http.NewRequest("GET", "http://localhost:8000/profileInfo", nil)
	// request.Header.Add("Authorization", token)

	// client := &http.Client{}
	// response, err := client.Do(request)
	response, err := ProfileGetterQuery(username)
	fmt.Println(response)
	if err != nil {
		fmt.Println(err)
		return clientModel.ProfileInfo{}
	} else {
		result := map[string]interface{}{}
		json.Unmarshal(response, &result)
		fmt.Println(result)
		var returnInfo clientModel.ProfileInfo
		//json.NewDecoder(response.Body).Decode(&returnInfo)
		ActualResult := result["document"].(map[string]interface{})
		returnInfo.Fullname = ActualResult["Fullname"].(string)
		fmt.Println(returnInfo.Fullname)
		returnInfo.Address = [5]string{ActualResult["Address1"].(string), ActualResult["Address2"].(string), ActualResult["City"].(string), "Tx", ActualResult["Zipcode"].(string)}
		fmt.Println(returnInfo)
		return returnInfo
	}
}

func FuelQuoteForm(deliveryInfo clientModel.DeliveryData) FullDeliveryData {
	jsonValue, err := json.Marshal(deliveryInfo)
	if err != nil {
		fmt.Println(err)
		return FullDeliveryData{}
	}

	response, err := fuelQuoteInserter(username, deliveryInfo.SuggestedPrice, deliveryInfo.Amount)
	if err != nil {
		fmt.Println(err)
		return FullDeliveryData{}
	} else {
		result := map[string]interface{}{}
		json.Unmarshal(response, &result)
		fmt.Println(result)
		var returnInfo FullDeliveryData
		ActualResult := result["document"].(map[string]interface{})
		returnInfo.DeliveryData = deliveryInfo
		returnInfo.FuelPrice = ActualResult["FuelPrice"].(float64)
		returnInfo.FuelType = ActualResult["FuelType"].(string)
		returnInfo.FuelVolume = ActualResult["FuelVolume"].(float64)
		returnInfo.TotalPrice = ActualResult["TotalPrice"].(float64)
		return returnInfo
	}
}

func FuelQuoteInfo(username string) clientModel.FullDeliveryData {
	response, err := fuelQuoteGetter(username)
	if err != nil {
		fmt.Println(err)
		return clientModel.FullDeliveryData{}
	} else {
		result := map[string]interface{}{}
		json.Unmarshal(response, &result)
		fmt.Println(result)
		var returnInfo clientModel.FullDeliveryData
		ActualResult := result["document"].(map[string]interface{})
		returnInfo.DeliveryData.SuggestedPrice = ActualResult["SuggestedPrice"].(float64)
		returnInfo.DeliveryData.Amount = ActualResult["Amount"].(float64)
		returnInfo.FuelPrice = ActualResult["FuelPrice"].(float64)
		returnInfo.FuelType = ActualResult["FuelType"].(string)
		returnInfo.FuelVolume = ActualResult["FuelVolume"].(float64)
		returnInfo.TotalPrice = ActualResult["TotalPrice"].(float64)
		return returnInfo
}

func StatesQuery() clientModel.States {
	request, _ := http.NewRequest("GET", "http://localhost:8000/getStates", nil)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return clientModel.States{}
	} else {
		var stateInfo clientModel.States
		json.NewDecoder(response.Body).Decode(&stateInfo)
		fmt.Println(stateInfo)
		return stateInfo
	}
}
