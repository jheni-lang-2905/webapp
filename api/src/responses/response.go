package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma resposta em json na requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if err := json.NewEncoder(w).Encode(dados); err != nil {
			log.Fatal(err)
		}
	}
}

// retorna um erro em format de json
func Erro(w http.ResponseWriter, statuscode int, erro error) {
	JSON(w, statuscode, struct {
		Erro string `json:erro`
	}{
		Erro: erro.Error(),
	})
}
