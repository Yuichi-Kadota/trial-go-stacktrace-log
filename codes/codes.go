package codes

type Code string

const (
	OK Code = "OK"

	Database Code = "database_error"
	Unknown  Code = "unknown"
)
