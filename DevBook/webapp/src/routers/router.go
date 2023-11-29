package router

import (
	"webapp/src/routers/rotas"

	"github.com/gorilla/mux"
)

// Gerar reotnra todas as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)

}
