package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Carregar()

	fmt.Println("Runing API!")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}
