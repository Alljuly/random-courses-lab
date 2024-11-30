package main

import (
	"fmt"
	"time"
)


func main(){
	//i := 0


	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j++{
	// 	fmt.Println("Incrementando J",j)
	// 	time.Sleep(time.Second)
	// }


	for {
		fmt.Println("for infinito")
		time.Sleep(time.Second)
	}
}