package utils

type ErrorCode string

const (
	ErrorCodeCustom          ErrorCode = "T000"
	ErrorCodeInvalidDatabase ErrorCode = "T001"
	ErrorCodeQueryError      ErrorCode = "T002"
	ErrorCodeSerialize       ErrorCode = "T003"
	ErrorCodeDeserialize     ErrorCode = "T004"
	ErrorCodeInvalidArgs     ErrorCode = "T005"
)

type CLIError struct {
	Code    ErrorCode
	Message string
	Detail  string
	Err     error
}
