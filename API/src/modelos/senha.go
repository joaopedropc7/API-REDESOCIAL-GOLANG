package modelos

// Representa o formato da requisicao de alterar senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
