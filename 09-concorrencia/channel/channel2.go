// package main

// import (
// 	"fmt"
// 	"time"
// )

// // Channel (canal) - É a forma de comunicação entre goroutines
// // Channel é um tipo

// func doisTresQuatroVezes(base int, c chan int) {
// 	time.Sleep(time.Second)
// 	c <- 2 * base // enviando dados para o canal

// 	time.Sleep(time.Second)
// 	c <- 3 * base

// 	time.Sleep(3 * time.Second)
// 	c <- 4 * base
// }

// func main() {
// 	c := make(chan int)
// 	go doisTresQuatroVezes(2, c)

// 	a, b := <-c, <-c // Recebendo dados do canal
// 	fmt.Println(a, b)
// 	fmt.Println(<-c)
// }
