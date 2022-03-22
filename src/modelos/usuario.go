package modelos

import (
	"errors"
	"strings"
	"time"
)

//Usuario
type Usuario struct {
	ID       uint64    `json: "id, omitempty"`
	Nome     string    `json: "nome, omitempty"`
	Nick     string    `json: "nick, omitempty"`
	Email    string    `json: "email, omitempty"`
	Senha    string    `json: "senha, omitempty"`
	CriadoEm time.Time `json: "nome, omitempty"`
}

func (usuario *Usuario) preparar() error {
	if err := usuario.validar(); err != nil {
		return err
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode está em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode está em branco")
	}

	if usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode está em branco")
	}

	if usuario.Senha == "" {
		return errors.New("A senha é obrigatório e não pode está em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
