package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //driver do para conexão com banco de dados
)

/*
	Para conectar com o banco necessário baixar um driver para conexão com banco
	go get github.com/go-sql-driver/mysql
*/

func ConectarComBanco() (*sql.DB, error) {
	//abre conexão com banco de dados
	db, err := sql.Open("mysql", config.StringConnectionSql)
	//trata o erro
	if err != nil {
		return nil, err
	}
	//faz um ping e caso dê erro ele fecha a conexão com banco de dados e retorna erro
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	//senão ele retorna a conexão com banco de dados aberta
	return db, nil
}
