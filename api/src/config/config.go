package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

/*
	pacote baixados para pegar as variaveis de ambiente
	go get github.com/joho/godotenv
*/

var (
	StringConnectionSql = ""
	Port                = 0
	SecretKey           []byte
)

func CarregarVariaveis() {
	var err error

	//carrega as variaveis de ambiente, caso dê erro ao carregar ele para o programa com o log.fatal
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	//converte a variavel da porta para inteiro com a função atoi do strconv e caso dê erro usa-se uma outra porta
	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	StringConnectionSql = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("IP_DB"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
