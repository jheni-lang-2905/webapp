package middlewares

import (
	"api/src/autenticacao"
	"api/src/responses"
	"fmt"
	"log"
	"net/http"
)

/*
	Camada que fica entre a requisição e a resposta usadas para as funções e aninhando as função com o uso dessa função
*/

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.URL, r.Host)
		next(w, r)
	}
}

// Autenticar valida se o usuario que esta fazendo a requisição está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando ..")
		if err := autenticacao.ValidarToken(r); err != nil {
			responses.Erro(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
