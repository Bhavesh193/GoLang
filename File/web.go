package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://html5up.net/paradigm-shift"

func main() {
	fmt.Println("Handling Web Request")

	response , err := http.Get(url) ;

	if(err != nil ){
		panic(err)
	}

	

	// fmt.Println("Respone : " , response)
	// fmt.Println("Respone Body : " , response.Body)

	defer response.Body.Close()

	databytes , err := io.ReadAll(response.Body)
	if(err != nil ){
		panic(err)
	}
	fmt.Println("databytes : " ,string(databytes) )
}