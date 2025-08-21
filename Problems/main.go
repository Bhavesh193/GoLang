package main

import "fmt"

func main() {
	// var num int = 42
	var num int 
	fmt.Println("Enter the number: ")
	fmt.Scanf("%v ", &num)
	result := CheckEvenOrOdd(num)
	fmt.Println(result)
}

func CheckEvenOrOdd(num int) string {
	
	if num%2 == 0 {
		return "It is Even num "
	} else {
		return "It is odd"
	}
}