package main


import "fmt"


func fibonacci(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n - 2)

	
}

func main(){


	fmt.Printf("%d", fibonacci(6))
}