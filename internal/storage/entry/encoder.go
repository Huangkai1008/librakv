package entry

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sync"
)

//nolint:gochecknoglobals // Global variable for the encoder buffer pool
var bufPool = sync.Pool{
	New: func() any { return make([]byte, HeaderSize) },
}

// NewEncoder creates a streaming Entry encoder.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: bufio.NewWriter(w)}
}

// Encoder wraps an underlying io.Writer and
// allows you to stream Entry encodings on it.
type Encoder struct {
	w *bufio.Writer
}

// Encode takes any Entry and streams it to the underlying writer.
// It returns the total size of encoded bytes and the first error
// encountered while encoding, if any.
//
// A successful Encode returns total size > 0, err == nil; otherwise, it returns
// total size = 0, err != nil.
func (e *Encoder) Encode(entry *Entry) (int64, error) {
	buf, ok := bufPool.Get().([]byte)
	if !ok {
		return 0, errors.New("invalid buffer type")
	}

	defer bufPool.Put(&buf)

	// Write CRC.
	offset := 0
	binary.BigEndian.PutUint32(buf[offset:offset+CRCSize], entry.CRC)
	offset += CRCSize

	// Write Type.
	buf[offset] = byte(entry.Type)
	offset += TypeSize

	// Write Timestamp.
	binary.BigEndian.PutUint64(buf[offset:offset+TimeStampSize], entry.Timestamp)
	offset += TimeStampSize

	// Write Key Length.
	binary.BigEndian.PutUint32(buf[offset:offset+KeySize], uint32(len(entry.Key)))
	offset += KeySize

	// Write Value Length.
	binary.BigEndian.PutUint32(buf[offset:], uint32(len(entry.Value)))

	if _, err := e.w.Write(buf); err != nil {
		return 0, fmt.Errorf("write entry header: %w", err)
	}

	if _, err := e.w.Write(entry.Key); err != nil {
		return 0, fmt.Errorf("write entry key: %w", err)
	}

	if _, err := e.w.Write(entry.Value); err != nil {
		return 0, fmt.Errorf("write entry value: %w", err)
	}

	if err := e.w.Flush(); err != nil {
		return 0, fmt.Errorf("flush buffered data: %w", err)
	}

	return int64(HeaderSize + len(entry.Key) + len(entry.Value)), nil
}
