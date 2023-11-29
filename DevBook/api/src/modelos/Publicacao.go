package modelos

import (
	"errors"
	"strings"
	"time"
)

// Representa uma publicação feita por um usuario
type Publicacao struct {
	ID        uint64    `json: "id, omiempty"`
	Titulo    string    `json: "titulo, omiempty"`
	Conteudo  string    `json: "conteudo, omiempty"`
	AutorID   uint64    `json: "autorID, omiempty"`
	AutorNick string    `json: "autorNick, omiempty"`
	Curtidas  uint64    `json: "curtidas"`
	CriadaEm  time.Time `json: "criadaEm, omiempty"`
}

// Prepara valida e formata a publicação
func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}
	publicacao.formatar()
	return nil

}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("O título não pode estar me branco")
	}
	if publicacao.Conteudo == "" {
		return errors.New("O conteúdo não pode estar me branco")
	}
	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)

}
