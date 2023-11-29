package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Representa a resposta de erro da API
type ErrorAPI struct {
	Erro string `json: "erro"`
}

// Retorna uma responsa em formato Json para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

// Essa funcao faz o tratmento do erro do stauts code 400 ou maior
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErrorAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
