package main


import (
 	"fmt"
	"teste/enderecos"
)


func main(){
	tipoEndereco := enderecos.TipoEndereco("Avenida do Futuro")

	fmt.Println(tipoEndereco)
}