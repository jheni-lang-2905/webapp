package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// GerarRouter retornar um router com as rotas configuradas
func GerarRouter() *mux.Router {
	r := mux.NewRouter()
	return rotas.ConfigurarRotas(r)
}
