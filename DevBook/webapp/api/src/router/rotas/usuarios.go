package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rotas{

	{
		URI:                 "/usuarios",
		Metodo:              http.MethodPost,
		Funcao:              controllers.CriarUsuario,
		RequerAuthenticacao: false,
	},

	{
		URI:                 "/usuarios",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarUsuarios,
		RequerAuthenticacao: true,
	},

	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarUsuario,
		RequerAuthenticacao: true,
	},

	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodPut,
		Funcao:              controllers.AtualizarUsuario,
		RequerAuthenticacao: true,
	},

	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.DeletarUsuario,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/seguir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.SeguirUsuario,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.PararDeSeguirUsuario,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/seguidores",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarSeguidores,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/seguindo",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarSeguindo,
		RequerAuthenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/atualizar-senha",
		Metodo:              http.MethodPost,
		Funcao:              controllers.AtualizarSenha,
		RequerAuthenticacao: true,
	},
}
