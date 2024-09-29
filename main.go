package main

import (
	"fmt"
)

// variables
var age int = 25
var height float64 = 5.4
var name string = "hello"
var decided bool = false 
const pi = 67.12
var last_name = "world"


// take user input
var user string

func add(a, b int) (result int) {
	// return a + b
	result = a + b
	return
}

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("Denominator must not be zero")
// 	}
// 	return a/b, nil
// }

// func divide(a, b float64) (float64, string) {
// 	if b == 0 {
// 		return 0, "Denominator must not be zero"
// 	}
// 	return a/b, "nil"
// }

// array
// var num [5] int 
var num = [5]int {1, 2, 3, 4, 5}


// structure
type Person struct {
	FirstName string
	LastName string
	Age int
}

// nested structure
type Contact struct {
	Email string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
}

// Go routines -> runs functions concurrently
func sayHello() {
	fmt.Println("hello")
}

func sayHi() {
	fmt.Println("hi")
}



func main() {
	fmt.Println("Welcome to Go")
	// fmt.Println(age)
	// fmt.Println(height)
	// fmt.Println(name)
	// fmt.Println(decided)
	// fmt.Println(pi)
	// fmt.Println(last_name)
	// fmt.Printf("age is %d\n", age)

	// person := 123
	// fmt.Println(person)

	// fmt.Scan(&user)      // scan reas upto the whitespace
	// reader := bufio.NewReader(os.Stdin)
	// name, _ := reader.ReadString('\n')
	// fmt.Println("Welcome", name)

	// function call

	// ans := add(2, 4)
	// fmt.Println(ans)  


	// error handle

	// ans, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error", err)
	// }
	// fmt.Println("Division", ans)

	// ans, _ := divide(10, 0)
	// fmt.Println(ans)

	// fmt.Println(num)
	// fmt.Println(len(num))

	// // slices
	// number := []int {2, 3, 4, 5, 6}
	// number = append(number, 7, 8, 9)    // append adds element to a slice
	// fmt.Println(number)

	// make function creates a slice with length and capacity
	// when decalring slice without any initil value then use make function
	// num := make([], int, 0)
	// number1 := make([]int, 3, 5)


	// for loop
	// for i:=0; i<4; i++ {
	// 	fmt.Println("hello hello")
	// }

	// for loop with range
	// number2 := []int {11, 12, 13, 14, 15}
	// for index, value := range number2 {
	// 	fmt.Printf("Index: %d, value: %d\n", index, value)
	// }


	// strings
	// str := "apple,orange"
	// parts := strings.Split(str, ",")
	// count := strings.Count(str, "apple")
	// fmt.Println(parts)
	// fmt.Println(count)


	// str1 := "Hello"
	// str2 := "World"
	// result := strings.Join([]string {str1, str2}, " ")
	// result1 := strings.Join([]string {str1, "Bikesh", str2}, " ")
	// fmt.Println(result)
	// fmt.Println(result1)
	

	// defer keyword
	// fmt.Println("Start")
	// defer fmt.Println("Middle")   // defer -> used to ensure that a fucntion call is performed later in a program's execution
	// fmt.Println("End")

	// maps
	// a := map[string]string {"brand": "Ford", "model": "2024"}
	// a := make(map[string]string)
	// a["city"] = "Oslo"
	// a["Hello"] = "greetings"
	// a["Hello World"] = "program"
	// delete(a, "Hello")
	// fmt.Println(a)

	// city, exists := a["city"]
	// fmt.Println("City:", city)
	// fmt.Println("Exists:", exists)


	// pointers
	// num3 := 3
	// ptr := &num3
	// fmt.Println(ptr)        // prints address
	// fmt.Println(*ptr)       // prints value stored in that address


	// var bikesh Person
	// bikesh.FirstName = "bikesh"
	// bikesh.LastName = "shrestha"
	// bikesh.Age = 21

	// bikesh := Person {
	// 	FirstName: "Hello",
	// 	LastName: "Bro",
	// 	Age: 21,
	// }

	// fmt.Println("Bikesh:", bikesh)

	// var person2 = new(Person)
	// person2.FirstName = "Hello"
	// person2.LastName = "Gello"
	// person2.Age = 21

	// fmt.Println(person2)   // there will be & sign with output


	// employee1 := Employee
	// var employee1 Employee
	// employee1.Person_Details = Person {
	// 	FirstName: "Hello",
	// 	LastName: "Bro",
	// 	Age: 21,
	// }

	// employee1.Person_Contact.Email = "hello@gmail.com"
	
	// fmt.Println(employee1)
	// fmt.Println(employee1.Person_Contact)

	
	// go routines

	// go sayHello()
	// sayHi()
	// time.Sleep(3000 * time.Millisecond)   // waiting for a moment to allow go routine to finish


}

