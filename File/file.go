package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Handling files in Golang");
	content := "This is th Golang tutorial and currenly I am working on the File handling task in Golang" ;

	file , err := os.Create("./sampleData.txt")
	checkErr(err) ;
	
	length, err := io.WriteString(file , content)
	checkErr(err) ;

	fmt.Println("Length is: " , length) 

	defer file.Close()

	readFile("./sampleData.txt")
	

}


func checkErr(err error){
	if err != nil {
		panic(err)
	}
}

func readFile(filename string){
	databyte, _ := ioutil.ReadFile(filename)
		fmt.Println("Data is : " , string (databyte) )

}