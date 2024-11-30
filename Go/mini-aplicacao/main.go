package main


import (
	"linha-de-comando/app"
	"os"
	"log"

)


func main(){
	application := app.Generate()

	erro := application.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}