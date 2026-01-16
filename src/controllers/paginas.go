package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requisicoes"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	utils.ExecutarTemplate(w, "login.html", nil)
}

func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroApi{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}

	var publicacoes []models.Publicacao
	if err := json.NewDecoder(response.Body).Decode(&publicacoes); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroApi{Erro: err.Error()})
		return
	}
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutarTemplate(w, "cadastro.html", struct {
		Publicacoes []models.Publicacao
		UsuarioID   uint64
	}{
		Publicacoes: publicacoes,
		UsuarioID:   usuarioID,
	})
}

// CarregarPaginaPrincipal carrega a página principal com as publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}

	var publicacoes []models.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroApi{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutarTemplate(w, "home.html", struct {
		Publicacoes []models.Publicacao
		UsuarioID   uint64
	}{
		Publicacoes: publicacoes,
		UsuarioID:   usuarioID,
	})
}

// função que carrega pagina de edição de publicacao
func CarregarPaginaDeEdicaoDePublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroApi{Erro: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d", config.APIURL, publicacaoId)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroApi{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}

	var publicacao models.Publicacao
	if err = json.NewDecoder(response.Body).Decode(&publicacao); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroApi{Erro: err.Error()})
		return
	}

	utils.ExecutarTemplate(w, "atualizar-publicacao.html", publicacao)
}

// Carrega pagina com os usuarios que atendem o filtro passado
func CarregarPaginaDeUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.APIURL, nomeOuNick)

	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroApi{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}

	var usuarios []models.Usuario
	if err = json.NewDecoder(response.Body).Decode(&usuarios); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroApi{Erro: err.Error()})
		return
	}

	utils.ExecutarTemplate(w, "usuarios.html", usuarios)
}

// Carrega pagaina do perfil do usuario
func CarregarPerfilDoUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioID, err := strconv.ParseUint(params["usuarioId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroApi{Erro: err.Error()})
		return
	}

	usuario, err := models.BuscarUsuarioCompleto(usuarioID, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroApi{Erro: err.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioLogadoId, _ := strconv.ParseUint(cookie["id"], 10, 64)
	utils.ExecutarTemplate(w, "usuario.html", struct {
		Usuario         models.Usuario
		UsuarioLogadoId uint64
	}{
		Usuario:         usuario,
		UsuarioLogadoId: usuarioLogadoId,
	})
}
