package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rotas{

	{
		URI:                 "/publicacoes",
		Metodo:              http.MethodPost,
		Funcao:              controllers.CriarPublicacao,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/publicacoes",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarPublicacoes,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoID}",
		Metodo:              http.MethodPost,
		Funcao:              controllers.BuscarPublicacao,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoID}",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarPublicacao,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoID}",
		Metodo:              http.MethodPut,
		Funcao:              controllers.AtualizarPublicacao,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoID}",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.DeletarPublicacao,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuariosID}/publicacoes", //usuário especifico
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarPublicacoesPorUsuario,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoID}/curtir", //usuário especifico
		Metodo:              http.MethodPost,
		Funcao:              controllers.CurtirPublicacao,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoID}/descurtir", //usuário especifico
		Metodo:              http.MethodPost,
		Funcao:              controllers.DescurtirPublicacao,
		RequerAuthenticacao: true,
	},
}
