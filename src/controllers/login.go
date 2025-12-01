package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"webapp/src/responses"
)

func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroApi{Erro: err.Error()})
		return
	}

	response, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer(usuario))

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroApi{Erro: err.Error()})
		return
	}

	token, _ := ioutil.ReadAll(response.Body)

	fmt.Println(response.StatusCode, string(token))
}
