package service

import ("errors")

var (
	ErrNodeNotFound = errors.New("node not found")
	ErrNotFile 		= errors.New("node is not a file")
)
