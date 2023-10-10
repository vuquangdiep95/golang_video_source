package main

import (
    "log"
    "os"
)

func main() {
	//Create a file
    file, error := os.Create("test.txt")
    if error != nil {
        log.Fatal(error)
    }
	
	error = os.Rename("test.txt", "temp.txt")
    if error != nil {
        log.Fatal(error)
    }
}