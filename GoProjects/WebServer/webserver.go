package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler( w http.ResponseWriter , r *http.Request) {

	if r.URL.Path != "/hello"{
		http.Error(w , "404 Not Found", http.StatusAccepted)
		return
	}

	http.ServeFile(w , r , "hello.html")
}

func formHandler(w http.ResponseWriter , r *http.Request){

	err := r.ParseForm() ;
	if err != nil {
		fmt.Fprintf(w , "ParseForm()", err)
		return
	}

	fmt.Println("Post Request Successful")
	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Println("Name: " , name)
	fmt.Println("Email: " , email)

	fmt.Fprintf( w, "Name:%s\n" , name)
	fmt.Fprintf(w , "Email:%s\n " , email)
}

func main() {
	fileServer := http.FileServer(http.Dir("./"))
	http.Handle("/" , fileServer)
	http.HandleFunc("/hello" , helloHandler)
	http.HandleFunc("/form" , formHandler)

	fmt.Println("Server started at PORT 8000")
	err := http.ListenAndServe(":8000" , nil)

	if err != nil {
		log.Fatal(err)
	}
}