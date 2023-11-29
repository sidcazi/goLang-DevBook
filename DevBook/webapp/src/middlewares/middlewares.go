package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Autenticar Verifica  a existencia do cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//cookies.Ler
		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		proximaFuncao(w, r)
	}
}
