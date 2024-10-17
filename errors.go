package librakv

import "errors"

var (
	ErrorDatabaseIsRunning = errors.New("database is used by another process")
)
