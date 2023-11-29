package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Retorna uma resposta em JSON para requisiçao
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}

	}
}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
