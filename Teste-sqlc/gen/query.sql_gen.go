// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package db

import (
	"context"
)

const createPessoa = `-- name: CreatePessoa :exec
INSERT INTO Pessoas(codigo, nome, email, idade) VALUES (null, ?, ?, ?)
`

type CreatePessoaParams struct {
	Nome  string `db:"nome" json:"nome"`
	Email string `db:"email" json:"email"`
	Idade int32  `db:"idade" json:"idade"`
}

func (q *Queries) CreatePessoa(ctx context.Context, arg CreatePessoaParams) error {
	_, err := q.db.ExecContext(ctx, createPessoa, arg.Nome, arg.Email, arg.Idade)
	return err
}

const deletePessoaByCodigo = `-- name: DeletePessoaByCodigo :exec
DELETE FROM Pessoas WHERE codigo = ?
`

func (q *Queries) DeletePessoaByCodigo(ctx context.Context, codigo int32) error {
	_, err := q.db.ExecContext(ctx, deletePessoaByCodigo, codigo)
	return err
}

const deletePessoaByNome = `-- name: DeletePessoaByNome :exec
DELETE FROM Pessoas WHERE nome = ?
`

func (q *Queries) DeletePessoaByNome(ctx context.Context, nome string) error {
	_, err := q.db.ExecContext(ctx, deletePessoaByNome, nome)
	return err
}

const selectPessoaByName = `-- name: SelectPessoaByName :one
SELECT codigo, nome, email, idade FROM Pessoas WHERE nome = ?
`

func (q *Queries) SelectPessoaByName(ctx context.Context, nome string) (Pessoa, error) {
	row := q.db.QueryRowContext(ctx, selectPessoaByName, nome)
	var i Pessoa
	err := row.Scan(
		&i.Codigo,
		&i.Nome,
		&i.Email,
		&i.Idade,
	)
	return i, err
}

const selectPessoaBycodigo = `-- name: SelectPessoaBycodigo :one
SELECT codigo, nome, email, idade FROM Pessoas WHERE codigo = ?
`

func (q *Queries) SelectPessoaBycodigo(ctx context.Context, codigo int32) (Pessoa, error) {
	row := q.db.QueryRowContext(ctx, selectPessoaBycodigo, codigo)
	var i Pessoa
	err := row.Scan(
		&i.Codigo,
		&i.Nome,
		&i.Email,
		&i.Idade,
	)
	return i, err
}

const selectPessoas = `-- name: SelectPessoas :many
SELECT codigo, nome, email, idade FROM Pessoas
`

func (q *Queries) SelectPessoas(ctx context.Context) ([]Pessoa, error) {
	rows, err := q.db.QueryContext(ctx, selectPessoas)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pessoa
	for rows.Next() {
		var i Pessoa
		if err := rows.Scan(
			&i.Codigo,
			&i.Nome,
			&i.Email,
			&i.Idade,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePessoaByCodigo = `-- name: UpdatePessoaByCodigo :exec
UPDATE Pessoas SET nome=?, email=?, idade=? WHERE codigo = ?
`

type UpdatePessoaByCodigoParams struct {
	Nome   string         `db:"nome" json:"nome"`
	Email  string `db:"email" json:"email"`
	Idade  int32  `db:"idade" json:"idade"`
	Codigo int32  `db:"codigo" json:"codigo"`
}

func (q *Queries) UpdatePessoaByCodigo(ctx context.Context, arg UpdatePessoaByCodigoParams) error {
	_, err := q.db.ExecContext(ctx, updatePessoaByCodigo,
		arg.Nome,
		arg.Email,
		arg.Idade,
		arg.Codigo,
	)
	return err
}
