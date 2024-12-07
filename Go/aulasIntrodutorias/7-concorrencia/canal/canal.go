package main


import (
	"fmt"
	"time"

)

func main(){
	canal := make(chan string)

	go escrever("Olá mundo", canal)


	for {
		mensagem, aberto := <-canal
		if !aberto{
			break
		}
		fmt.Println(mensagem)
	}

}

func escrever(str string, canal chan string){
	for i := 0; i < 5; i++{
		
			canal <- str
		
		time.Sleep(time.Second)
	}

	close(canal)
}