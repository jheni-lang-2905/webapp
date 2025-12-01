package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rotas{
	{
		URI:          "/publicacoes",
		Metodo:       http.MethodPost,
		Funcao:       controllers.CriarPublicacao,
		Autenticacao: true,
	},
	{
		URI:          "/publicacoes/{id}",
		Metodo:       http.MethodGet,
		Funcao:       controllers.BuscarPublicacaoPorId,
		Autenticacao: true,
	},
	{
		URI:          "/publicacoes/",
		Metodo:       http.MethodGet,
		Funcao:       controllers.BuscarPublicacao,
		Autenticacao: true,
	},
	{
		URI:          "/publicacoes/{id}",
		Metodo:       http.MethodPut,
		Funcao:       controllers.AtualizarPublicacao,
		Autenticacao: true,
	},
	{
		URI:          "/publicacoes/{id}",
		Metodo:       http.MethodDelete,
		Funcao:       controllers.DeletarPublicacao,
		Autenticacao: true,
	},
	{
		URI:          "/usuarios/{usuarioId}/publicacoes",
		Metodo:       http.MethodGet,
		Funcao:       controllers.BuscarPublicacoesPorUsuario,
		Autenticacao: true,
	},
	{
		URI:          "/publicacoes/{publicacaoId}/curtir",
		Metodo:       http.MethodPost,
		Funcao:       controllers.CurtirPublicacao,
		Autenticacao: true,
	},
	{
		URI:          "/publicacoes/{publicacaoId}/descurtir",
		Metodo:       http.MethodPost,
		Funcao:       controllers.DescurtirPublicacao,
		Autenticacao: true,
	},
}
