package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rotas{
	{
		URI:          "/usuarios",
		Metodo:       http.MethodPost,
		Funcao:       controllers.CriarUsuario,
		Autenticacao: false,
	},
	{
		URI:          "/usuarios",
		Metodo:       http.MethodGet,
		Funcao:       controllers.BuscarTodosUsuarios,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{id}",
		Metodo:       http.MethodGet,
		Funcao:       controllers.BuscarUsuarioPorId,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{id}",
		Metodo:       http.MethodPut,
		Funcao:       controllers.AtualizarUSuario,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{id}",
		Metodo:       http.MethodDelete,
		Funcao:       controllers.DeletarUsuario,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{id}/seguir",
		Metodo:       http.MethodPost,
		Funcao:       controllers.SeguirUsuario,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{id}/nao-seguir",
		Metodo:       http.MethodPost,
		Funcao:       controllers.PararSeguir,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{id}/seguidores",
		Metodo:       http.MethodGet,
		Funcao:       controllers.BuscarSeguidores,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{id}/seguindo",
		Metodo:       http.MethodGet,
		Funcao:       controllers.BuscarSeguindo,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{id}/atualizar-senha",
		Metodo:       http.MethodPost,
		Funcao:       controllers.AtualizarSenha,
		Autenticacao: true,
	},
}
