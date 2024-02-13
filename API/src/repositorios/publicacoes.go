package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// REPRESENTA UM REPOSITORIO DE PUBLICACAOES
type Publicacoes struct {
	db *sql.DB
}

// Cria um repositorio de publicacaoes
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Insere uma publicacao no banco de dados
func (repositorio Publicacoes) CriarPublicacao(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
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

func (repositorio Publicacoes) BuscarPorId(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		select p.*, u.nick from 
		publicacoes p inner join usuarios u 
		on u.id = p.autor_id where p.id = ?`,
		publicacaoID)
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

// Buscar publicacaoes dos usuarios que seguem
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {

	linhas, erro := repositorio.db.Query(`
		select distinct p.*, u.nick from publicacoes p inner join usuarios u 
		on u.id = p.autor_id inner join seguidores s on p.autor_id = s.usuario_id
		where u.id = ? or s.seguidor_id = ?  order by 1 desc`, usuarioID, usuarioID)
	if erro != nil {
		return []modelos.Publicacao{}, erro
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

// ATUALIZAR ALTERA OS DADOS DE UMA PUBLICACAO NO BANCO DE DADOS
func (repositorio Publicacoes) Atualizar(publicacaoId uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	fmt.Println(publicacao)
	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoId); erro != nil {
		return erro
	}

	return nil
}

// Exclui uma publicacao do banco de dados
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

// Traz todas as publicacoes de um usuario especifico
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {

	linhas, erro := repositorio.db.Query(`
		select p.*, u.nick from publicacoes p inner join usuarios u
		on u.id = p.autor_id where p.autor_id = ?`, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacaoes []modelos.Publicacao

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
		publicacaoes = append(publicacaoes, publicacao)
	}
	return publicacaoes, nil
}

// Adiciona uma curtida na publicacao
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

// remove uma curtida na publicacao
func (repositorio Publicacoes) Descurtir(publicacaoId uint64) error {
	statement, erro := repositorio.db.Prepare(`
		update publicacoes set curtidas =
	    CASE 
	        WHEN curtidas > 0 THEN curtidas -1
	    	ELSE 0 
	    END 
		where id = ?`)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}
	return nil
}
