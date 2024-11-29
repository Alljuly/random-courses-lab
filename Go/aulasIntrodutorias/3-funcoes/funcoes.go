package main

import (
	"fmt"
)


func sum(a int, b int) int {
	return a + b
}

func main(){

	x := 10
	y := 25

	res := sum(x,y)

	var f = func(txt string) {
		fmt.Println(txt)

	}

	f("Texto da função")

	fmt.Println(res)
}