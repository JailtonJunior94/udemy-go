package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[43040606867] = "Jailton Junior"
	aprovados[98978787787] = "Stefany Teixeira"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[43040606867])
	delete(aprovados, 98978787787)

	fmt.Println(aprovados[98978787787])
}
