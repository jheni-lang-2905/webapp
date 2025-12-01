package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErroApi struct {
	Erro string `json:"erro"`
}

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Contenty-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(dados); err != nil {
		log.Fatal(err)
	}
}

func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroApi
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
