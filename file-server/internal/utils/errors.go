package utils

import "errors"

var (
	ErrFileNotFound        = errors.New("File not found")
	ErrNoFileUploaded      = errors.New("No file uploaded")
	ErrFailedToSaveFile    = errors.New("Failed to save file")
	ErrNoFilenameSpecified = errors.New("No filename specified")
	ErrFailedToOpenFile    = errors.New("Failed to open file")
	ErrFailedToCopyFile    = errors.New("Failed to copy file")
)
