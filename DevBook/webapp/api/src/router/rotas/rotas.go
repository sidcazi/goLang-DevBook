package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas para api
type Rotas struct {
	URI                 string
	Metodo              string
	Funcao              func(http.ResponseWriter, *http.Request)
	RequerAuthenticacao bool
}

// Configurar coloca todas as rotas para o usuario
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasPublicacoes...)

	for _, rota := range rotas {

		if rota.RequerAuthenticacao {
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)

		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	return r

}
