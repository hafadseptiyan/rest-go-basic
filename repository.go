package main

import (
	"log"
	"net/http"
	"encoding/json"
)

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
 	var arr_user []Users
 	var response Response
	db := connect()
 	defer db.Close()
	rows, err := db.Query("Select id,first_name,last_name,address from person")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
  		if err := rows.Scan(&users.Id, &users.FirstName, &users.LastName, &users.Address); err != nil {
   			log.Fatal(err.Error())
		} else {
   			arr_user = append(arr_user, users)
  		}
	 }

	response.Version = "v1"
	response.Message = "Success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func insertUsersMultipart(w http.ResponseWriter, r *http.Request) {
  var response Response
  
	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	first_name 	:= r.FormValue("first_name")
	last_name 	:= r.FormValue("last_name")
	address 	:= r.FormValue("address")

	_, err = db.Exec("INSERT INTO person (first_name, last_name,address) values (?,?,?)",
		first_name,
		last_name,
		address,
	)

	if err != nil {
		log.Print(err)
	}

	response.Version = "v1"
	response.Message = "Success Add Data"
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func updateUsersMultipart(w http.ResponseWriter, r *http.Request) {
  var response Response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id 			:= r.FormValue("user_id")
	first_name 	:= r.FormValue("first_name")
	last_name 	:= r.FormValue("last_name")
	address 	:= r.FormValue("address")

	_, err = db.Exec("UPDATE person set first_name = ?, last_name = ? , address = ? where id = ?",
		first_name,
		last_name,
		address,
		id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Version = "v1"
	response.Message = "Success Update Data"
	log.Print("Update data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func deleteUsersMultipart(w http.ResponseWriter, r *http.Request) {
  var response Response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("user_id")

	_, err = db.Exec("DELETE from person where id = ?",
		id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Version = "v1"
	response.Message = "Success Delete Data"
	log.Print("Delete data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}