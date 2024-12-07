package main


import "fmt"


func main(){
	tarefas := make(chan int, 100)
	resultado := make(chan int, 100)

	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)

	for i := 0; i < 45; i++{
		tarefas <- i
	}
	close(tarefas)
	
	for i := 0; i < 45; i++{
		resultado :=  <-resultado
		fmt.Println(resultado)
	}
	close(resultado)

}


func worker(tarefa <-chan int, resultado chan<- int){
	for numero := range tarefa {
		resultado <- fibonacci(numero)
	}
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n - 2)

	
}

