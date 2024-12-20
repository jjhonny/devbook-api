package securanca

import "golang.org/x/crypto/bcrypt"

// Hash recebe uma string e retorna a hash dela
func Hash(senha string) ([]byte, error) {

	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha compara uma senha e um hash e retorna se elas são iguais ou não
func VerificarSenha(senhaComHash string, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senhaString))
}
