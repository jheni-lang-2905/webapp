package models

import (
	"errors"
	"strings"
	"time"
)

// Publicacao modelo de publicacao
type Publicacao struct {
	ID        uint64    `json:"id,omitempyt"`
	Titulo    string    `json:"titulo,omitempyt"`
	Conteudo  string    `json:"conteudo,omitempyt"`
	AutorId   uint64    `json:"autorId,omitempyt"`
	AutorNick string    `json:"autorNick,omitempyt"`
	Curtidas  uint64    `json:"curtidas,omitempyt"`
	CriadoEm  time.Time `json:criadoEm",omitempyt"`
}

func (publicacao *Publicacao) Preparar() error {
	if err := publicacao.validar(); err != nil {
		return err
	}

	publicacao.formatar()
	return nil
}

// validar a publicacao recebida
func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("O titulo é obrigatório e não pode estar em branco")
	}

	if publicacao.Conteudo == "" {
		return errors.New("O Conteúdo é obrigatório e não pode estar em branco")
	}

	return nil
}

// formatar a publicacao
func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
