package vcat

import "errors"

var (
	ErrTranscriptNotFound  error = errors.New("no trancript found for given url")
	ErrCaptionsNotFound    error = errors.New("no captions found for url")
	ErrUnsupportedLanguage error = errors.New("given language is not supported for url")
)
