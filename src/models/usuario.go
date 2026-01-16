package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

// representa uma pessoa usando a rede social
type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacao"`
}

// buscar usuario completo faz 4 requisições na api para montar usuario
func BuscarUsuarioCompleto(UsuarioId uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacao := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, UsuarioId, r)
	go BuscarSeguidores(canalSeguidores, UsuarioId, r)
	go BuscarSeguindo(canalSeguindo, UsuarioId, r)
	go BuscarPublicacoes(canalPublicacao, UsuarioId, r)

	var (
		usuario    Usuario
		seguidores []Usuario
		seguindo   []Usuario
		publicacao []Publicacao
	)

	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.ID == 0 {
				return Usuario{}, errors.New("Erro ao buscar usuario")
			}
			usuario = usuarioCarregado
		case seguidoresCarregados := <-canalSeguidores:
			if seguidoresCarregados == nil {
				return Usuario{}, errors.New("Erro ao buscar os seguidores")
			}
			seguidores = seguidoresCarregados
		case seguindoCarregados := <-canalSeguindo:
			if seguindoCarregados == nil {
				return Usuario{}, errors.New("Erros ao buscar quem o usuario esta seguindo")
			}
			seguindo = seguindoCarregados
		case publicacoesCarregadas := <-canalPublicacao:
			if publicacoesCarregadas == nil {
				return Usuario{}, errors.New("Erro ao buscar publicacoes")
			}
			publicacao = publicacoesCarregadas
		}
	}

	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacao
	return usuario, nil
}

// Buscar dados base para montar usuario
func BuscarDadosDoUsuario(canal chan<- Usuario, UsuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, UsuarioId, r)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- Usuario{}
		return
	}

	defer response.Body.Close()

	var usuario Usuario
	if err = json.NewDecoder(response.Body).Decode(&usuario); err != nil {
		canal <- Usuario{}
		return
	}

	canal <- usuario
}

// Buscar seguidores chama api para buscar seguidores do usuario
func BuscarSeguidores(canal chan<- []Usuario, UsuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.APIURL, UsuarioId, r)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- nil
		return
	}

	defer response.Body.Close()
	var seguidores []Usuario

	if err = json.NewDecoder(response.Body).Decode(&seguidores); err != nil {
		canal <- nil
		return
	}

	if seguidores == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguidores
}

func BuscarSeguindo(canal chan<- []Usuario, UsuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.APIURL, UsuarioId, r)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- nil
		return
	}

	defer response.Body.Close()
	var seguindo []Usuario

	if err = json.NewDecoder(response.Body).Decode(&seguindo); err != nil {
		canal <- nil
		return
	}

	if seguindo == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- seguindo
}

// chama api para buscar publicacoes de usuario
func BuscarPublicacoes(canal chan<- []Publicacao, UsuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.APIURL, UsuarioId)

	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- nil
		return
	}

	defer response.Body.Close()

	var publicacoes []Publicacao
	if err = json.NewDecoder(response.Body).Decode(&publicacoes); err != nil {
		canal <- nil
		return
	}

	if publicacoes == nil {
		canal <- make([]Publicacao, 0)
		return
	}
	canal <- publicacoes
}
