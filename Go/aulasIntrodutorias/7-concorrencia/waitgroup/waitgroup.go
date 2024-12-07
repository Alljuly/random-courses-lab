package main


import (
	"fmt"
	"time"
	"sync"
)


func main(){

	var wait sync.WaitGroup

	wait.Add(2)

	go func(){
		write("Go routine")
		wait.Done()
	}()

	go func(){
		write("Hello World")
		wait.Done()
	}()

	wait.Wait()
}

func write(str string) {
	for i:= 0; i <5; i++{
		fmt.Println(str)
		time.Sleep(time.Second)
	}
}