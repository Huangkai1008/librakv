package librakv

import "errors"

var (
	ErrDatabaseIsRunning = errors.New("database is used by another process")
)
