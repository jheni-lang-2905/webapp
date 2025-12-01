package models

import (
	"api/src/security"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

/*
	pacote que fará validação do e-mail
	go get github.com/badoux/checkmail
*/

// Usuario struct do tipo usuario
type Usuario struct {
	Id       uint64    `json."id,omitempty"`
	Nome     string    `json."nome,omitempty"`
	Nick     string    `json."nick,omitempty"`
	Email    string    `json."email,omitempty"`
	Senha    string    `json."senha,omitempty"`
	CriadoEm time.Time `json."CriadoEm,omitempty"`
}

// Preparar vai chamar os metodos para formatar e valida se esta vazio
func (usuario *Usuario) Preparar(etapa string) error {
	if err := usuario.validarVazio(etapa); err != nil {
		return err
	}
	if err := usuario.formatarCampos(etapa); err != nil {
		return err
	}
	return nil
}

func (usuario *Usuario) validarVazio(etapa string) error {
	if usuario.Nome == "" || usuario.Nick == "" || usuario.Email == "" {
		return errors.New(fmt.Sprintf("Erro pois objeto dos campos nome: %s, Nick: %s, E-mail: %s estão em branco", usuario.Nome, usuario.Nick, usuario.Email))
	}

	if err := checkmail.ValidateFormat(usuario.Email); err != nil {
		return errors.New("Email esta em forma inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New(fmt.Sprintf("Erro pois o campo de senha é obrigatório"))
	}
	return nil
}

func (Usuario *Usuario) formatarCampos(etapa string) error {
	Usuario.Nome = strings.TrimSpace(Usuario.Nome)
	Usuario.Nick = strings.TrimSpace(Usuario.Nick)
	Usuario.Email = strings.TrimSpace(Usuario.Email)
	if etapa == "cadastro" {
		senhaHash, err := security.Hash(Usuario.Senha)
		if err != nil {
			return err
		}
		Usuario.Senha = string(senhaHash)
	}
	return nil
}
