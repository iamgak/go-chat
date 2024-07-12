package main

import "errors"

var (
	ErrFileNotFound   = errors.New("File not found")
	ErrCannotLoadFile = errors.New("Unable to load file")
	ErrCannotSaveFile = errors.New("Unable to save file")
)
