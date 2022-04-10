package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

//cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statment, err := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}
	defer statment.Close()

	resultado, err := statment.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	ultimoIdInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	linhas, err := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?", nomeOuNick, nomeOuNick,
	)

	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

//Buscar por ID
func (repositorio Usuarios) BuscarPorId(ID uint64) (modelos.Usuario, error) {
	linhas, err := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?", ID,
	)

	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return modelos.Usuario{}, err
		}
	}
	return usuario, nil
}

//Atualizar usuario
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, err := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.ID); err != nil {
		return err
	}

	return nil
}

//Deletar usuarios
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

//Buscar por email e retorna seu id e senha com hash
func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, err := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if err = linha.Scan(&usuario.ID, &usuario.Senha); err != nil {
			return modelos.Usuario{}, err
		}
	}

	return usuario, nil
}
