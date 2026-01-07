package models

import "time"

//publicacao representa uma publicacao feita por um usuario
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"Titulo,omitempty"`
	Conteudo  string    `json:"Conteudo,omitempty"`
	AutorID   uint64    `json:"AutorID,omitempty"`
	AutorNick string    `json:"AutorNick,omitempty"`
	Curtidas  uint64    `json:"Curtidas"`
	CriadaEm  time.Time `json:"CriadaEm,omitempty"`
}
