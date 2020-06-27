package main

import "fmt"

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Executou!")
	ch <- 4
	ch <- 5
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	fmt.Println(<-ch)
}
