package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseURL = "http://localhost:8080"

type identity struct {
	Name    string `json:"Name`
	NIM     string `json:"NIM`
	Address string `json:"Address`
}

func fetchUsers() ([]person, error) {
	var err error
	var people = &http.People{}
	var data []person

	request, err := http.NewRequest("POST", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := people.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchUser(NIM string) (person, error) {
	var err error
	var people = &http.People{}
	var data person

	var param = url.Values{}
	param.Set("NIM", NIM)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/user", payload)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := people.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	var user1, err = fetchUser("2501983004")
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Printf("Name: %s\t NIM: %s\t Address: %d\n", user1.Name, user1.NIM, user1.Address)
}
