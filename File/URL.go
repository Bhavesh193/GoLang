package main

import (
	"fmt" 
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gr54359"

func main() {
	fmt.Println("MYURL" , myurl)

  result , err := url.Parse(myurl) ;

  if(err != nil ){
	panic(err)
  }

 fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("RawQuery:", result.RawQuery)
	
	queryParams := result.Query()
	fmt.Println("Query Parameters:",queryParams )

	for key , value := range queryParams{
		fmt.Println("%s : %s\n" , key, value)
	}
}