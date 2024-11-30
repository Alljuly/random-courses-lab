package main

import "fmt"


func diaDaSemana(dia int) string {
	var d string
	switch dia {
	case 1:
		d = "Domingo"
	case 2:
		d = "Segunda"
	case 3: 
		d = "Terça"

	case 4:
		d = "Quarta"

	case 5: 
		d = "Quinta"

	case 6:
		d = "Sexta"

	case 7: 
		d = "Sábado"
	
	default:
		d = "Mensagem invalida"
	}

	return d
}

func diaDaSemana2(dia string) int {
	switch {
	case dia == "Domingo":
		return 1

	case dia == "Segunda":
		return 2

	case dia == "Terça":
		return 3
	
	case dia == "Quarta":
		return 4
	
	case dia == "Quinta":
		return 5
	
	case dia == "Sexta":
		return 6

	case dia == "Sábado":
		return 7

	default:
		return 0
	}
}

func main(){

	diaInt := 5
	diaStr := "Domingo"

	fmt.Println("Hoje é:",diaDaSemana(diaInt))
	fmt.Printf("Hoje é o %d dia da semana",diaDaSemana2(diaStr))
}