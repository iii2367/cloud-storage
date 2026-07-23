package repository

import (
	"errors"
)

var (
	ErrRepository       = errors.New("repository error")
	ErrTreeNodeNotFound = errors.New("tree node not found")
)
