package main
 

import (
	"fmt" 
	"time"
	"math/rand"
)


func main(){
	exemplo := multiplexar(escrever("oii"), escrever("tudo bem?"))

	for i:= 0; i<10;i++{
		fmt.Println(<-exemplo)
	}

}

func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string{
	canalSaida := make(chan string)

	go func() {
		for {
			
			select {
			case mensagem := <-canalEntrada1:
				canalSaida <-mensagem
			case mensagem := <-canalEntrada2:
				canalSaida <-mensagem
		}
		}
	}()

	return canalSaida
}

func escrever(texto string) <-chan string{

	canal := make(chan string)


	go func(){
		for{
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}

