package main

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
