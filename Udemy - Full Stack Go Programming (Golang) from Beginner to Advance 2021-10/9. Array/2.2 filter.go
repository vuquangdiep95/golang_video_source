package main

import "fmt"

func main() {

	names := [5]string{"Kim", "Jim", "Bill", "Robert", "David"}

	fmt.Printf("Name:%v\n", names)

	fmt.Println(names[:])

	fmt.Println(names[:3])

	fmt.Println(names[2:])

	fmt.Println(names[1:4])

}
