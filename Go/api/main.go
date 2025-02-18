package main

import (
	"api/src/config"
	"api/src/routers"
	"log"
	"net/http"
	"fmt"
)

func main(){

	config.Prepare()

	r := router.Generate()

	fmt.Print(config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",config.Port), r))

}