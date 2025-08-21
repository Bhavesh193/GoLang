package main 

import "fmt"

func main() {
	fmt.Println("Function :- ")
    // result,_ := addNum(5,8)
	// fmt.Println("Result is :- " , result)
	// result := menuSum(3,6,2,8,3) ;
	// fmt.Println("Result is  : " , result)

	////////////////////////////////////////////////////////

		UserData := User{"Bhavesh" , "bh@gmail.com" , false , 10}
		UserData.getData()


}

func addNum (num1 int , num2 int) (int,string ){
	var result = num1 + num2 ;
	return result , "Hii"
}

func menuSum (values ...int) int {
	total := 0 

	for _ , value := range values {
		total += value
	}

	return total
}

type User struct {
	Name string 
	Email string 
	status bool 
	Age int
}

func (u User) getData() {
	fmt.Println("User Name is : " , u.Name)
}