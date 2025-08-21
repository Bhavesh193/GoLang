package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main () {
fmt.Println("Welcome to the GOlang server")
fmt.Println("Get Data")
PerformGetRequest()
fmt.Println("Post Data")
PerformPostRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	response , err := http.Get(myurl)


	if err != nil {
		panic(err)
	}

	defer response.Body.Close() ;


	fmt.Println("Status : ", response.Status)
	fmt.Println("Status Code : ", response.StatusCode)

	content , _ := io.ReadAll(response.Body)
	fmt.Println("content: " , string(content))
} 

func PerformPostRequest() {
	const url ="http://localhost:3000/post"

	requestBody := strings.NewReader(`
	{
	 " coursename" : "Let's go with golang" ,
	 "price" : 100 ,
	 "platform" : "Youtube" 
	}
	`)

	response, err := http.Post(url , "application/json" , requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content , _ := io.ReadAll(response.Body)

	fmt.Println("response.Body: ", response.Body)
	fmt.Println("Content: ", content)
	fmt.Println("Content: ", string(content))
}
