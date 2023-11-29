package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Representa um repositorio de Publicações
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicação  no Banco
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {

	statement, erro := repositorio.db.Prepare("INSERT INTO publicacoes (titulo, conteudo, autor_id) values ( ?, ?, ?)")

	if erro != nil {
		return 0, nil
	}

	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)

	if erro != nil {
		return 0, erro
	}
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

// BuscarPorID traz uma única publicação do banco de dados
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
	select p.*, u.nick from 
	publicacoes p inner join usuarios u
	on u.id = p.autor_id where p.id = ?`,
		publicacaoID,
	)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Buscar traz as publicações dos usuários seguidos e também do próprio usuário que fez a requisição
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	select distinct p.*, u.nick from publicacoes p 
	inner join usuarios u on u.id = p.autor_id 
	inner join seguidores s on p.autor_id = s.usuario_id 
	where u.id = ? or s.seguidor_id = ?
	order by 1 desc`,
		usuarioID, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Atualizar altera uma publicação  do usuário logado
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacacao.Titulo, publicacacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}
	return nil
}

// Deletar exclui uma publicação do banco de dados
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from publicacoes where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Busca as publicacões de um usuario especifico
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		select p.*, u.nick from publicacoes p
		join usuarios u on u.id = p.autor_id
		where p.autor_id = ?`,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacacao.ID,
			&publicacacao.Titulo,
			&publicacacao.Conteudo,
			&publicacacao.AutorID,
			&publicacacao.Curtidas,
			&publicacacao.CriadaEm,
			&publicacacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacacao)
	}
	return publicacoes, nil
}

// Curtir publicação implementado de 1 em 1 no banco
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {

	statement, erro := repositorio.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Descurtir publicação subtrai de 1 em 1 no banco
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {

	statement, erro := repositorio.db.Prepare(`
	update publicacoes set curtidas = 
	case 
		when curtidas > 0 then curtidas -1
		else  0
	 end
	where id =?
	`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}
