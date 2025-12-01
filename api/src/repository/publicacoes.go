package repository

import (
	"api/src/models"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NewRepositoryPublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repository Publicacoes) CriarPublicacao(Publicacoes models.Publicacao) (uint64, error) {
	statement, err := repository.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	resul, err := statement.Exec(Publicacoes.Titulo, Publicacoes.Conteudo, Publicacoes.AutorId)
	if err != nil {
		return 0, err
	}

	ultimoId, err := resul.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(ultimoId), nil
}

func (repository Publicacoes) BuscarPublicacaoPorId(idPublicacao uint64) (models.Publicacao, error) {
	linhas, err := repository.db.Query(`
		select p.*, u.nick from publicacoes p inner join usuarios u on u.id = p.autor_id where p.id = ?
	`, idPublicacao)
	if err != nil {
		return models.Publicacao{}, err
	}
	defer linhas.Close()
	var publicacao models.Publicacao

	if linhas.Next() {
		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorId,
			&publicacao.AutorNick,
		); err != nil {
			return models.Publicacao{}, err
		}
	}
	return publicacao, nil
}

func (repository Publicacoes) BuscarPublicacao(userId uint64) ([]models.Publicacao, error) {

	linhas, err := repository.db.Query(`select distinct p.*, u.nick from publicacoes p inner join usuarios u on
	 u.id = p.autor_id = s.usuario_id where i.id = ? or s.seguidor_id = ? order by 1 desc`, userId, userId)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()
	var publicacoes []models.Publicacao
	for linhas.Next() {
		var publicacao models.Publicacao
		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorId,
			&publicacao.AutorNick,
		); err != nil {
			return nil, err
		}

		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

func (repository Publicacoes) AtualizarPublicacao(publicacaoId uint64, publicacao models.Publicacao) error {
	statement, err := repository.db.Prepare(`update publicacoes set titulo = ?, conteudo = ? where id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoId); err != nil {
		return err
	}
	return nil
}

func (repository Publicacoes) DeletarPublicacao(publicacaoId uint64) error {
	statement, err := repository.db.Prepare(`delete from publicacoes where id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacaoId); err != nil {
		return err
	}
	return nil
}

func (repository Publicacoes) BuscarPorUsuario(usuarioId uint64) ([]models.Publicacao, error) {
	linhas, err := repository.db.Query(`
	select p.*, u.nick from publicacoes p
	join usuarios u on u.id = p.autor_id
	where p.autor_id = ?`, usuarioId)

	if err != nil {
		return nil, err
	}
	defer linhas.Close()
	var publicacoes []models.Publicacao
	for linhas.Next() {
		var publicacao models.Publicacao
		if err = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorId,
			&publicacao.AutorNick,
		); err != nil {
			return nil, err
		}

		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

func (repository Publicacoes) CurtirPublicacao(publicacaoId uint64) error {
	statement, err := repository.db.Prepare(`update publicacoes set curtidas = curtidas + 1 where id = ?`)
	if err != nil {
		return err
	}

	defer statement.Close()
	if _, err = statement.Exec(publicacaoId); err != nil {
		return err
	}
	return nil
}

func (repository Publicacoes) DescurtirPublicacao(publicacaoId uint64) error {
	statement, err := repository.db.Prepare(`
	update publicacoes set curtidas =
	CASE WHEN curtidas > 0 THEN curtidas - 1
	ELSE curtidas END
	where id = ?`)

	if err != nil {
		return err
	}

	defer statement.Close()
	if _, err = statement.Exec(publicacaoId); err != nil {
		return err
	}
	return nil
}
