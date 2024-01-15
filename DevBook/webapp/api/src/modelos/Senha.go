package modelos

// Senha representa o formato de requisição para alateração de senha
type Senha struct {
	Nova  string `json: "nova"`
	Atual string `json: "atual"`
}
