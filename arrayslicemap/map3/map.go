package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"Jailton Junior": 14500.00,
		},
		"S": {
			"Stefany Teixeira": 4500.00,
		},
	}

	delete(funcsPorLetra, "G")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
