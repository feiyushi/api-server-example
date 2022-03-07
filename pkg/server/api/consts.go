package api

type ErrorCode string

const (
	BadRequest          ErrorCode = "BadRequest"
	InternalServerError ErrorCode = "InternalServerError"
	NotFound            ErrorCode = "NotFound"
)
