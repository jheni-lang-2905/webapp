package security

import (
	"golang.org/x/crypto/bcrypt"
)

/*
	pacote baixado para colocar um hash na senha do usuario
	go get golang.org/x/crypto/bcrypt
*/

// Hash recebe  uma string e coloca um hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha compara uma senha e um hash e retorna se são iguais
func VerificarSenha(senhaString, senhaComHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
