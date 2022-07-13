package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
)

type Student struct {
	Id    string
	Name  string
	Grade int8
}

func main() {
	var users, err = fetchUsers()
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	for _, each := range users {
		fmt.Printf("ID %s\t Name:%s \t Grade:%d \t \n", each.Id, each.Name, each.Grade)
	}

	var idFlag = flag.String("id", "1", "userId")

	var user, err1 = fetchUser(string(*idFlag))

	if err1 != nil {
		fmt.Println("Error")
		return
	}

	fmt.Printf("ID : %s \t Name:%s \t Grade:%d\n", user.Id, user.Name, user.Grade)

}

func fetchUsers() ([]Student, error) {
	var baserUrl = "http://localhost:8000"
	var err error
	var client = &http.Client{}
	var data []Student

	req, err := http.NewRequest("GET", baserUrl+"/users", nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)

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

func fetchUser(Id string) (Student, error) {
	var err error
	var client = &http.Client{}
	var data Student

	var param = url.Values{}

	param.Set("id", Id)
	// var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("GET", "http://localhost:8000"+"/user", nil)
	if err != nil {
		return data, err
	}

	request.URL.RawQuery = param.Encode()

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)

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
