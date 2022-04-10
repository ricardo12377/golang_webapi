package seguranca

import "golang.org/x/crypto/bcrypt"

//Hash recebe uma string e coloca um hash nela
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

//Verificar senha compara uma senha com hash e compara se elas são iguais
func VerificarSenha(senhaHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaString))
}
