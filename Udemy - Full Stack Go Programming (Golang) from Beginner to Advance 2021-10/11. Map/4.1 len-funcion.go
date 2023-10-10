package main

import "fmt"

func main() {

	var employee = make(map[string]int)

	employee["Kin"] = 10
	employee["Jim"] = 20
	employee["Robert"] = 30

	employeeList := make(map[string]int)

	fmt.Println(len(employee))

	fmt.Println(len(employeeList))

}
