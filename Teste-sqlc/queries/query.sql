-- name: SelectPessoas :many
SELECT * FROM Pessoas;

-- name: CreatePessoa :exec
INSERT INTO Pessoas(codigo, nome, email, idade) VALUES (null, ?, ?, ?) ;

-- name: SelectPessoaByName :one
SELECT * FROM Pessoas WHERE nome = ?;

-- name: SelectPessoaBycodigo :one
SELECT * FROM Pessoas WHERE codigo = ?;

-- name: DeletePessoaByNome :exec
DELETE FROM Pessoas WHERE nome = ?;

-- name: DeletePessoaByCodigo :exec
DELETE FROM Pessoas WHERE codigo = ?;

-- name: UpdatePessoaByCodigo :exec
UPDATE Pessoas SET nome=?, email=?, idade=? WHERE codigo = ?;