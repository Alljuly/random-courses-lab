package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)


type Gato struct {
	Idade uint  `json:"idade"`
	Cor string  `json:"cor"`
	Nome string `json:"nome"`
} 

func main(){

	//Marshal
	gato := Gato{4, "frajola preto e branco", "Marlon"}

	g, erro := json.Marshal(gato)
	if erro != nil {
		log.Fatal(erro)
	}


	fmt.Println(gato)
	fmt.Println(bytes.NewBuffer(g))


	//Unmarshall

	gatoJson := `{"idade" : 10, "nome" : "Branco", "cor" : "amarelo"}`


	
	var gatoUnMarshall Gato

	if erro := json.Unmarshal([]byte (gatoJson), &gatoUnMarshall); erro != nil{
		log.Fatal(erro)
	}
	
	
	fmt.Println(gatoJson)
	fmt.Println(gatoUnMarshall)
}