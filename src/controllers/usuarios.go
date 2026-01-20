package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requisicoes"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroApi{Erro: err.Error()})
	}

	url := fmt.Sprintf("%s/usuarios", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroApi{Erro: err.Error()})
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}
	responses.JSON(w, response.StatusCode, nil)
}

// parar de seguir usuarios
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroApi{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/parar-de-seguir", config.APIURL, usuarioID)

	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroApi{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroApi{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/seguir", config.APIURL, usuarioID)

	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroApi{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroApi{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(usuario))

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// Chama a api para atualizar a senha do usuario
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	senhas, err := json.Marshal(map[string]string{
		"atual": r.FormValue("atual"),
		"nova":  r.FormValue("nova"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroApi{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d/atualizar-senha", config.APIURL, usuarioID)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(senhas))

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// deletar usuario chama api para deletar usuario
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodDelete, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroApi{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
