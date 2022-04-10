package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

//Vai chamar os metodos para validar e formatar o usuario recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if err := usuario.validar(etapa); err != nil {
		return err
	}

	if err := usuario.formatar(etapa); err != nil {
		return err
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode está em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode está em branco")
	}

	if usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode está em branco")
	}

	if err := checkmail.ValidateFormat(usuario.Email); err != nil {
		return errors.New("O e-mail inserido é inválido!")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatório e não pode está em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaHash, err := seguranca.Hash(usuario.Senha)
		if err != nil {
			return err
		}

		usuario.Senha = string(senhaHash)
	}

	return nil
}
