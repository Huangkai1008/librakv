package librakv

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofrs/flock"
)

const (
	FileLockName = "libra.lock"
)

type Key []byte

type Value []byte

type Database struct {
	options  *Options
	fileLock *flock.Flock
}

// Open creates a new Database instance with dataDir and the specified options.
//
// If the dataDir doesn't exist, it will be created automatically.
func Open(dataDir string, opts ...Option) (*Database, error) {
	options := DefaultOptions()
	if err := options.Apply(opts...); err != nil {
		return nil, err
	}

	// If the dataDir does not exist, create it.
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("create datadir (path %s): %w", dataDir, err)
	}

	fileLock := flock.New(filepath.Join(dataDir, FileLockName))
	locked, err := fileLock.TryLock()
	if err != nil {
		return nil, fmt.Errorf("get file lock: %w", err)
	}

	if !locked {
		return nil, ErrDatabaseIsRunning
	}

	return &Database{
		options:  options,
		fileLock: fileLock,
	}, nil
}
