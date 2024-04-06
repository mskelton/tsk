package utils

type CLIError struct {
	Message string
	Detail  string
	Err     error
}
