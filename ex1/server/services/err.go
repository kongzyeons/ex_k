package services

import "errors"

// Repo
var (
	ErrRepoBeerCreate      = errors.New("err repo beer ceate")
	ErrRepoBeerGet         = errors.New("err repo beer get")
	ErrRepoBeerGetPaginate = errors.New("err repo beer get paginate")
	ErrRepoBeerDelete      = errors.New("err repo beer delete")
	ErrrepoBeerUpdate      = errors.New("err repo beer update")

	ErrRepoLogCreate = errors.New("err repo log create")
)

// Check
var (
	ErrNotFound  = errors.New("err not found")
	ErrCheckName = errors.New("err check name")
)

// Request
var (
	ErrCreateBeerRequest     = errors.New("err create beer request")
	ErrGetPaginatBeerRequest = errors.New("err get paginate beer request")
)

// os
var (
	ErrcCreateFolder = errors.New("err create folder")
	ErrOepnFile      = errors.New("err open file")
	ErrSaveFile      = errors.New("err save file")
	ErrRemoveFile    = errors.New("err remove file")
	ErrRenameFile    = errors.New("err rename file")
)
