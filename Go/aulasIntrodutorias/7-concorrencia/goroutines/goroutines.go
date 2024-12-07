package main



import (
	"fmt"
	"time"
)


func main(){

	go write("Go routine")
	write("Hello World")
}

func write(str string) {
	for{
		fmt.Println(str)
		time.Sleep(time.Second)
	}
}