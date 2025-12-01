package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var usuario models.Usuario
	if err := json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		responses.Erro(w, http.StatusBadGateway, err)
		return
	}

	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewRepositoryUsuarios(db)
	usuarioLogado, err := repository.BuscarPorEmail(usuario.Email)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if err := security.VerificarSenha(usuario.Senha, usuarioLogado.Senha); err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, err := autenticacao.CriarToken(usuarioLogado.Id)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	w.Write([]byte(fmt.Sprintf(token)))
}
