package repository

import "errors"

// ErrNotFound é usado quando nenhum documento é encontrado na consulta
var ErrNotFound = errors.New("usuário não encontrado")
