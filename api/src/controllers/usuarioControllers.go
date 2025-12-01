package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.Usuario
	err = json.Unmarshal(corpoRequest, &user)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Preparar("cadastro"); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()
	repository := repository.NewRepositoryUsuarios(db)

	user.Id, err = repository.Criar(user)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)

}

func BuscarTodosUsuarios(w http.ResponseWriter, r *http.Request) {
	paramsRota := strings.ToLower(r.URL.Query().Get("usuario"))
	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUsuarios(db)
	usuarios, err := repository.BuscarTodos(paramsRota)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, usuarios)
}

func BuscarUsuarioPorId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
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
	usuario, err := repository.BuscarPorId(userId)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, usuario)
}

func AtualizarUSuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	corpoRequisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	usuariosId, err := autenticacao.ExtraiUsuariosId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userId != usuariosId {
		responses.Erro(w, http.StatusForbidden, errors.New("você não tem acesso para essa ação"))
	}

	var user models.Usuario
	if err := json.Unmarshal(corpoRequisicao, &user); err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	if err := user.Preparar("editar"); err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()
	repository := repository.NewRepositoryUsuarios(db)
	if err := repository.Atualizar(userId, user); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadGateway, err)
		return
	}

	usuariosId, err := autenticacao.ExtraiUsuariosId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userId != usuariosId {
		responses.Erro(w, http.StatusForbidden, errors.New("você não tem acesso para essa ação"))
	}

	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUsuarios(db)
	if err := repository.DeletarUsuario(userId); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	userId, err := autenticacao.ExtraiUsuariosId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	params := mux.Vars(r)

	seguidorId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userId == seguidorId {
		responses.Erro(w, http.StatusUnauthorized, errors.New("Não é possível seguir vc mesmo"))
		return
	}

	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUsuarios(db)

	if err := repository.SeguirUsuario(userId, seguidorId); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func PararSeguir(w http.ResponseWriter, r *http.Request) {
	userId, err := autenticacao.ExtraiUsuariosId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	params := mux.Vars(r)

	seguidorId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	if userId == seguidorId {
		responses.Erro(w, http.StatusUnauthorized, errors.New("Não é possível deixar de seguir vc mesmo"))
		return
	}

	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUsuarios(db)

	if err := repository.PararSeguir(seguidorId, userId); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	defer db.Close()
	repo := repository.NewRepositoryUsuarios(db)

	seguidores, err := repo.BuscarSeguidores(id)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, seguidores)
}

func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	defer db.Close()
	repo := repository.NewRepositoryUsuarios(db)
	users, err := repo.BuscarSeguindo(id)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	usuarioIdToken, err := autenticacao.ExtraiUsuariosId(r)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	params := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadGateway, err)
		return
	}

	if usuarioId != usuarioIdToken {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar a senha de outro usuários"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusBadGateway, err)
		return
	}
	var senha models.Senha
	if err := json.Unmarshal(body, &senha); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	defer db.Close()
	repository := repository.NewRepositoryUsuarios(db)
	senhaBanco, err := repository.BuscarSenha(usuarioId)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := security.VerificarSenha(senha.SenhaAtual, senhaBanco); err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	senhaHash, err := security.Hash(senha.NovaSenha)
	fmt.Printf("%s senha", string(senhaHash))
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err := repository.AtualizarSenha(usuarioId, string(senhaHash)); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)

}
