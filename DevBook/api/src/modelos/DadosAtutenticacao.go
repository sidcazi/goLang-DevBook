package modelos

// Dados da autenticação contem o token e o id do usuario
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
