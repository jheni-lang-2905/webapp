package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/*
	pacotes baixados para o main interarir com o router
	go get github.com/gorilla/mux
*/

/*
usado para criar o base64 das variveis de ambiente
func init() {
	chave := make([]byte, 64)

	_, err := rand.Read(chave)
	if err != nil {
		log.Fatal(err)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)

}*/

func main() {
	fmt.Println("Iniciando primeiro projeto no golang")
	config.CarregarVariaveis()

	r := router.GerarRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
