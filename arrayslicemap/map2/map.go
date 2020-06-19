package main

import "fmt"

func main() {
	funcionariosSalarios := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcionariosSalarios["Rafael Filho"] = 1350.0
	delete(funcionariosSalarios, "Inexistente")

	for nome, salario := range funcionariosSalarios {
		fmt.Println(nome, salario)
	}
}
