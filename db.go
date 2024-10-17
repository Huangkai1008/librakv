package librakv

import (
	"github.com/gofrs/flock"
	"os"
	"path/filepath"
)

const (
	FileLockName = "libra.lock"
)

type Database struct {
	fileLock *flock.Flock
}

// Open creates a new Database instance with dataDir and the specified options.
//
// If the dataDir doesn't exist, it will be created automatically.
func Open(dataDir string) (*Database, error) {
	// If the dataDir does not exist, create it.
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}

	fileLock := flock.New(filepath.Join(dataDir, FileLockName))
	locked, err := fileLock.TryLock()
	if err != nil {
		return nil, err
	}

	if !locked {
		return nil, ErrorDatabaseIsRunning
	}

	return &Database{
		fileLock: fileLock,
	}, nil
}
