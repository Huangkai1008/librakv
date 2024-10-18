package librakv

import (
	"errors"
)

const (
	DefaultLogFileThresholdSize = 64 << 20 // 64MB
)

var (
	ErrNilOption            = errors.New("nil option provided")
	ErrLogFileThresholdSize = errors.New("log file threshold size cannot be 0")
)

// Options of the database.
type Options struct {
	// logFileThresholdSize is the threshold size of each log file.
	// Active log file will be closed if reach the threshold.
	//
	// The default value is DefaultLogFileThresholdSize.
	logFileThresholdSize uint32
}

func DefaultOptions() *Options {
	return &Options{
		logFileThresholdSize: DefaultLogFileThresholdSize,
	}
}

type Option func(*Options) error

func WithLogFileThresholdSize(size uint32) Option {
	return func(o *Options) error {
		o.logFileThresholdSize = size
		return nil
	}
}

// Apply applies the given options to Options.
func (o *Options) Apply(opts ...Option) error {
	for _, opt := range opts {
		if opt == nil {
			return ErrNilOption
		}
		if err := opt(o); err != nil {
			return err
		}
	}

	if o.logFileThresholdSize <= 0 {
		return ErrLogFileThresholdSize
	}

	return nil
}
