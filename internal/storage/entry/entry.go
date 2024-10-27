package entry

import (
	"errors"
	"hash/crc32"
	"time"
)

const (
	CRCSize       = 4
	TypeSize      = 1
	TimeStampSize = 8
	KeySize       = 4
	ValueSize     = 4
	HeaderSize    = CRCSize + TypeSize + TimeStampSize + KeySize + ValueSize
)

var (
	ErrWrongMachineTime = errors.New("invalid machine time")
	ErrKeyTooLarge      = errors.New("key is too large")
	ErrValueTooLarge    = errors.New("value is too large")
)

// Type of entry.
type Type uint8

type Key []byte

type Value []byte

const (
	// Normal represents a regular entry.
	Normal Type = iota
	// Deleted represents a deleted entry.
	Deleted
)

// NewEntry creates a new entry.
func NewEntry(key Key, value Value) (*Entry, error) {
	timestamp := time.Now().UnixNano()
	if timestamp < 0 {
		return nil, ErrWrongMachineTime
	}

	e := &Entry{
		Type:      Normal,
		Key:       key,
		Value:     value,
		Timestamp: uint64(timestamp),
		CRC:       crc32.ChecksumIEEE(value),
	}
	return e, nil
}

// Entry represents a log entry with key-value pair and metadata.
//
// The entry header format looks like this:
//
//	┌─────────┬─────────┬───────────────┬──────────────┬────────────────┐
//	│ crc(4B) │ type(1B)│ timestamp(8B) │ key_size(4B) │ value_size(4B) │
//	└─────────┴─────────┴───────────────┴──────────────┴────────────────┘
//
// Notice: Timestamp is the unix time in nanoseconds.
//
// In original paper, the crc field don't join the cyclic redundancy check,
// the entry crc header format look like this:
//
//	┌─────────┬───────────────┬──────────────┬────────────────┐
//	│ type(1B)│ timestamp(8B) │ key_size(4B) │ value_size(4B) │
//	└─────────┴───────────────┴──────────────┴────────────────┘
//
// For simplicity, directly use the Value to generate CRC, see more details
// in NewEntry.
type Entry struct {
	// Type of the entry.
	Type      Type
	Key       Key
	Value     Value
	Timestamp uint64
	CRC       uint32
}
