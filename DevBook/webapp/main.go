package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	router "webapp/src/routers"
	"webapp/src/utils"

	"github.com/gorilla/securecookie"
)

func init() {
	hashKey := securecookie.GenerateRandomKey(16)
	fmt.Println(hashKey)
}

func main() {

	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
