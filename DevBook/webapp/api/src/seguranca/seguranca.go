package seguranca

import "golang.org/x/crypto/bcrypt"

// Hash recebe a string para colocar o hash
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// Compara a senha com o Hash
func VerificarSenha(senha string, senhaComHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senha), []byte(senhaComHash))

}
