package models

//Senha recebe o json do modelo de senha
type Senha struct {
	NovaSenha  string `json:novaSenha`
	SenhaAtual string `json:senhaAtual`
}
