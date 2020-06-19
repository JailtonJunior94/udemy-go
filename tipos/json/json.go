package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json to struct
	var p2 produto
	jsonString := `{"id":1,"nome":"Notebook","preco":1899.9,"tags":["Promoção","Eletrônico"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}
