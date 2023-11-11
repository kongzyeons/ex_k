package services

import "errors"

var (
	ErrRepoChilGetQuery   = errors.New("err repo chil get query")
	ErrRepoIconGetQuery   = errors.New("err repo icon get query")
	ErrRepoParentGetQuery = errors.New("err repo parent get query")
)

// Check
var (
	ErrNotFound = errors.New("err not found")
)
