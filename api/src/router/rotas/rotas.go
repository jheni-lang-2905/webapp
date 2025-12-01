package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rotas é um struct onde ficaram todas as rotas do projeto
type Rotas struct {
	URI          string
	Metodo       string
	Funcao       func(http http.ResponseWriter, r *http.Request)
	Autenticacao bool
}

// ConfigurarRotas configura cada rota
func ConfigurarRotas(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasPublicacoes...) //porque o rotaPublicacoes é um slice de rota e por isso ele precisa dar um append em cada uma

	for _, rota := range rotas {
		if rota.Autenticacao {
			//aninhando funções para configurações das rotas
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			//aninhando funções para configurações das rotas
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

	}
	return r
}
