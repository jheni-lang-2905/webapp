package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/responses"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroApi{Erro: err.Error()})
	}

	url := fmt.Sprintf("%s/usuarios", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroApi{Erro: err.Error()})
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.TratarStatusCodeDeErro(w, response)
		return
	}
	responses.JSON(w, response.StatusCode, nil)
}
