package main

import "fmt"

func inc1(n int) {
	n++
}

// Revisão: um ponteiro é representado por um *
func inc2(n *int) {
	// Revisão: * é usado para desferenciar, ou seja
	// Ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	// Revisão: & usado para obter o endereço da variável
	inc2(&n)
	fmt.Println(n)
}
