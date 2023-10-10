package main

import "fmt"

func main() {

	// Creating a map Using make() function
	var employee = map[int]string{
		100: "Kim",
		200: "Jim",
		300: "Robert",
	}

	//now we access the map items by calling index in square brackets
	fmt.Println(employee[100])

}
