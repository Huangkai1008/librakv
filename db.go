package librakv

import "os"

type Database struct {
}

// Open creates a new Database instance with dataDir and the specified options.
//
// If the dataDir doesn't exist, it will be created automatically.
func Open(dataDir string) (*Database, error) {
	// If the dataDir does not exist, create it.
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}

	return &Database{}, nil
}
