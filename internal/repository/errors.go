package repository

import "errors"

var ErrEntityNotFound = errors.New("Entity not found")
var ErrEntityExists = errors.New("Entity exists")
