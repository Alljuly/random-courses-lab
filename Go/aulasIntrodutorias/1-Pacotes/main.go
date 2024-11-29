package main

import (
	"fmt"
	"github.com/badoux/checkmail"
)

func main(){
	email := checkmail.ValidateFormat("alice@.com")

	fmt.Println(email)
}