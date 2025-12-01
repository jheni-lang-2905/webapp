package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	userId, err := autenticacao.ExtraiUsuariosId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publicacoes models.Publicacao
	if err = json.Unmarshal(body, &publicacoes); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	publicacoes.AutorId = userId

	if err := publicacoes.Preparar(); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryPublicacoes(db)
	publicacoesID, err := repository.CriarPublicacao(publicacoes)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, publicacoesID)

}
func BuscarPublicacaoPorId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idPublicacao, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewRepositoryPublicacoes(db)
	publicacao, err := repo.BuscarPublicacaoPorId(idPublicacao)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, publicacao)

}
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	userId, err := autenticacao.ExtraiUsuariosId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repository := repository.NewRepositoryPublicacoes(db)
	publicacao, err := repository.BuscarPublicacao(userId)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusAccepted, publicacao)

}
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	userid, err := autenticacao.ExtraiUsuariosId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	params := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(params["PublicacaoId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryPublicacoes(db)
	publicacaoSalvaNoBanco, err := repository.BuscarPublicacaoPorId(publicacaoId)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publicacaoSalvaNoBanco.AutorId != userid {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar uma publicacao que não seja a sua"))
		return
	}

	corpoRequsicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusBadGateway, err)
		return
	}

	var publicacao models.Publicacao
	if err = json.Unmarshal(corpoRequsicao, &publicacao); err != nil {
		responses.Erro(w, http.StatusBadGateway, err)
		return
	}

	if err = publicacao.Preparar(); err != nil {
		responses.Erro(w, http.StatusBadGateway, err)
		return
	}

	if err = repository.AtualizarPublicacao(publicacaoId, publicacao); err != nil {
		responses.Erro(w, http.StatusBadGateway, err)
		return
	}

	responses.JSON(w, http.StatusAccepted, nil)

}
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	userid, err := autenticacao.ExtraiUsuariosId(r)
	if err != nil {
		responses.Erro(w, http.StatusUnauthorized, err)
		return
	}
	params := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(params["PublicacaoId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryPublicacoes(db)
	publicacaoSalvaNoBanco, err := repository.BuscarPublicacaoPorId(publicacaoId)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if publicacaoSalvaNoBanco.AutorId != userid {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar uma publicacao que não seja a sua"))
		return
	}

	if err = repository.DeletarPublicacao(publicacaoId); err != nil {
		responses.Erro(w, http.StatusBadGateway, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func BuscarPublicacoesPorUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(params["usuarioId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryPublicacoes(db)
	publicacoes, err := repository.BuscarPorUsuario(usuarioId)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, publicacoes)
}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryPublicacoes(db)

	if err := repository.CurtirPublicacao(publicacaoId); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(params["publicacaoId"], 10, 64)
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.ConectarComBanco()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryPublicacoes(db)

	if err := repository.DescurtirPublicacao(publicacaoId); err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
