package main

type Users struct {
	Id        	string `form:"id" json:"id"`
	FirstName 	string `form:"firstname" json:"firstname"`
	LastName  	string `form:"lastname" json:"lastname"`
	Address		string `form:"address" json:"address"`
}

type Response struct {
	Version string `json:"status"`
	Message string `json:"message"`
	Data    []Users
}