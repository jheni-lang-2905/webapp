package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

// NewRepositoryUsuarios cria um repositorio de usuarios
func NewRepositoryUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um dado no banco de dados
func (repository Usuarios) Criar(Usuarios models.Usuario) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?);")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resul, err := statement.Exec(Usuarios.Nome, Usuarios.Nick, Usuarios.Email, Usuarios.Senha)
	if err != nil {
		return 0, err
	}

	ultimoID, err := resul.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(ultimoID), nil
}

// BuscarTodos retorna os usuarios que atendem ao filtro da rota
func (repository Usuarios) BuscarTodos(paramsRota string) ([]models.Usuario, error) {
	nomeNick := fmt.Sprintf("%%%s%%", paramsRota) // %string%

	linhas, err := repository.db.Query("select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeNick, nomeNick)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario
		if err = linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarPorId retorna um usuario por id
func (repository Usuarios) BuscarPorId(id uint64) (models.Usuario, error) {
	linhas, err := repository.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		id,
	)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if err = linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return models.Usuario{}, err
		}
		return usuario, nil
	}
	return usuario, nil
}

func (repository Usuarios) Atualizar(id uint64, user models.Usuario) error {
	statement, err := repository.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Nome, user.Nick, user.Email, id); err != nil {
		return err
	}
	return nil
}

// DeletarUsuario exclui as infos de um usuario
func (repository Usuarios) DeletarUsuario(id uint64) error {
	statement, err := repository.db.Prepare(
		"delete from usuarios where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		return err
	}
	return nil
}

func (repository Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linhas, err := repository.db.Query("select id, senha from usuarios where email = ?", email)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linhas.Close()

	var user models.Usuario

	for linhas.Next() {
		if err = linhas.Scan(&user.Id, &user.Senha); err != nil {
			return models.Usuario{}, err
		}
	}
	return user, nil
}

// SeguirUsuario permite que usuario siga outro
func (repository Usuarios) SeguirUsuario(userId, seguidorId uint64) error {
	//ignore ignora caso a linha na tabela ja exista com os mesmo valores que estamos tentando inserir
	statement, err := repository.db.Prepare("INSERT IGNORE INTO seguidores (usuario_id, seguidor_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userId, seguidorId); err != nil {
		return err
	}
	return nil
}

// PararSeguir permite que deixe de seguir um usuario
func (repository Usuarios) PararSeguir(seguidor_id, usuario_id uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM seguidores WHERE usuario_id = ? and seguidor_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(usuario_id, seguidor_id); err != nil {
		return err
	}
	return nil
}

// BuscarSeguidores traz todos os seguidores de um usuário
func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?`,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}

// BuscarSeguindo traz todos os usuários que um determinado usuário está seguindo
func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]models.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ?`,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.Id,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarSenha traz a senha de um usuário pelo ID
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where id = ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

// AtualizarSenha altera a senha de um usuário
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}

	return nil
}
