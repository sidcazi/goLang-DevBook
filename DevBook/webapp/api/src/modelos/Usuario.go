package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Uuário para rede social
type Usuario struct {
	ID       uint64    `json: "id,omitempty"`
	Nome     string    `json: "nome,omitempty"`
	Nick     string    `json: "nome,omitempty"`
	Email    string    `json: "nome,omitempty"`
	Senha    string    `json: "nome,omitempty"`
	CriadoEm time.Time `json: "nome,omitempty"`
}

// Formata os campos e valida
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O campo Nome é Obrigatório, não pode ser em branco!")
	}

	if usuario.Email == "" {
		return errors.New("O campo Email é Obrigatório, não pode ser em branco!")
	}
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é invalido")
	}

	if usuario.Senha == "" {
		return errors.New("O campo Senha é Obrigatório, não pode ser em branco!")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O campo Senha é Obrigatório, não pode ser em branco!")
	}
	return nil

}
func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro

		}
		usuario.Senha = string(senhaComHash)

	}

	return nil

}
