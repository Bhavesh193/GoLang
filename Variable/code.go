package main

import "fmt"

// "sort"
// "os"
// "strconv"
// "strings"
// "time"

func main() {
	// var isBoolean bool = true
	// fmt.Println("Variable is : ", isBoolean)
	// fmt.Printf("Variable Type is : %T\n ", isBoolean)

	// var str uint = 567757545;
	// fmt.Println("Str is : " , str) ;
	// fmt.Printf("Type is :%T\n ", str)

	// var num = 45.6
	// fmt.Println("Str is : " , num) ;
	// fmt.Printf("Type is :%T\n ", num)

	// welcome := 5656;
	// fmt.Println("Welcome bro : " , welcome)
	// fmt.Printf("Welcome bro : %T\n " , welcome)

	////////////// Input   //////////////////////////

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter the any number : ") ;

	// // comma ok || err

	// input , _ := reader.ReadString('\n')
	// fmt.Println("Thanks for the number : ", input)
	// fmt.Printf(" Type of Number is %T : ", input)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter the rating for our pizza app between 0 and 5 ")

	// input, _ := reader.ReadString('\n')
	// fmt.Println("Thanks for the rating : ", input)

	// numberRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Added 1 to your Rating: Now you rating is : ", numberRating+1)
	// }

	//////////////// Time  /////////////////////

	// fmt.Println("Time for the series")

	// now := time.Now() ;
	// fmt.Println("Time is : ", now)
	// fmt.Println("Time " , now.Format("02/01/2006 Monday"))

	// fmt.Println("Year:", now.Year())
	// fmt.Println("Month:", now.Month())
	// fmt.Println("Day:", now.Day())
	// fmt.Println("Hour:", now.Hour())
	// fmt.Println("Minute:", now.Minute())
	// fmt.Println("Second:", now.Second())
	// fmt.Println("Nanosecond:", now.Nanosecond())
	// fmt.Println("Weekday:", now.Weekday())

	////////////////// Pointer  ////////////////////////////

	// mynum := 100 ;
	// var ptr = &mynum ;
	// fmt.Println("Value of actual pointer is : " , ptr)
	// fmt.Println("Value of actual pointer is : " , *ptr)

	// mynum = *ptr + 10 ;

	// fmt.Println("New Value is : " , mynum)

	//  mynum := 10
	//     fmt.Println("Before:", mynum) //10

	//     updateValue(&mynum) // pass pointer
	//     fmt.Println("After:", mynum) //110

	////////////  Array /////////////////////////

	// var fruitList [4]string ;
	// fruitList[0] = "Mango"
	// fruitList[3] = "Guava"

	// fmt.Println("FruitList are : ", fruitList)
	// fmt.Println("FruitList length are : ", len(fruitList))

	// var vegList = [4]string{"Potato" , "Peas" , "Tomato"}

	// fmt.Println("FruitList are : ", vegList)
	// fmt.Println("FruitList length are : ", len(vegList))

	// arr := [3]int{10, 20, 30} // array of integers
	// ptr := &arr               // pointer to the whole array

	// fmt.Println("Array:", arr)       // [10 20 30]
	// fmt.Println("Pointer:", ptr)     // memory address of arr  &[10 20 30]
	// fmt.Println("Value via pointer:", *ptr) // [10 20 30]

	// arr := [3]int{10, 20, 30}
	//     fmt.Println("Before:", arr)

	//     updateArray(&arr) // pass pointer to array
	//     fmt.Println("After:", arr) // [999 20 30]

	//////////  Slice ////////////////

	// var vegList = []string{"Potato" , "Peas" , "Tomato"}

	// fmt.Println("FruitList are : ", vegList)
	// fmt.Printf("Type of FruitList is  %T\n", vegList)

	// vegList = append(vegList , "Cabbase" , "Carrot" , "LadyFinger")

	// fmt.Println(vegList)

	// vegList = append(vegList[2:5])
	// fmt.Println(vegList)

	// highScores:= make([]int, 4)
	// highScores [0] = 234
	// highScores [1] = 945
	// highScores [2] = 465
	// highScores [3] = 867

	// highScores = append (highScores, 555, 666, 321)
	// fmt.Println(highScores)

	// var ans = sort.IntsAreSorted(highScores)
	// fmt.Println(ans)

	// var vegList = []string{"Potato" , "Peas" , "Tomato", "Cabbase" , "Carrot" , "LadyFinger"}
	// fmt.Println(vegList)

	// var index int = 2 ;

	// vegList = append(vegList[:index] , vegList[index + 1 :]...)

	// fmt.Println(vegList)

	///////////////  Map    /////////////////////

	// language := make(map[int]string);

	// language[0] = "JS"
	// language[1] = "TS"
	// language[2] = "Java"
	// language[3] = "Py"
	// language[4] = "React"
	// language[5] = "RB"

	// fmt.Println("Before : ", language)

	// delete(language , 4)

	// fmt.Println("After : ", language)

	// for key , value := range language {
	// 	fmt.Printf("Key is %v , Value is %v\n", key , value)
	// }

	///////////////////   Struct  /////////////////////////////////

	// type User struct {
	// 	Name       string
	// 	Email      string
	// 	isLoggedIn bool
	// 	Age        int
	// }

	// // Struct
	// UserDetails := User{
	// 	"Bhavesh" , 
	// 	"bhavesh@gmail.com" ,
	// 	 false ,
	// 	 23 ,
	// }

	// fmt.Printf("UserDetails are : %+v\n" , UserDetails)
	// fmt.Printf("Name is : %v and email is : %v " , UserDetails.Name , UserDetails.Email)



	///////////////// LOOPS  ///////////////////////////////

	// age := 15 
	// if age > 18 {
	// 	fmt.Println("you are eligible to drive")
	// } else {
	// fmt.Println("you are not eligible to drive")

	// }


//     num := 674

//   if num % 2 == 0 {
// 	fmt.Println("It is an even number")
//   }else{
//     fmt.Println("It is an ODD number")
//   }


/////////////////////  Methods ///////////////////////////////////////






}


// func updateValue(num *int) {
//     *num = *num + 100
// }

// func updateArray(arr *[3]int) {
//     arr[0] = 999 // modifies original array
// }
