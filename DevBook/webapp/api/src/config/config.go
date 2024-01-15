package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (

	//conecat com o mysql
	StringConexaoBanco = ""

	//Porta listem api
	Porta = 0

	//É a chave par assinar o token
	SecretKey []byte
)

// Inicializa variaveis de ambiente
func Carregar() {

	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Porta, erro = strconv.Atoi(os.Getenv("API_PORTA"))
	if erro != nil {
		Porta = 9000

	}
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
	fmt.Println(StringConexaoBanco)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
