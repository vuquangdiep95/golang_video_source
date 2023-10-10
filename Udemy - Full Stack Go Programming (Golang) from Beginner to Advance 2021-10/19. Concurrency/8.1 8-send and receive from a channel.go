package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Start Main Function ........")

	ch := make(chan int)

	go send(ch)

	go receive(ch)

	time.Sleep(time.Microsecond * 100)

	fmt.Println("End of Main Function ........")

}

func send(ch chan int) {

	ch <- 50
}

func receive(ch chan int) {

	fmt.Println(200 + <-ch)
}
