package entry_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Huangkai1008/librakv/internal/storage/entry"
)

func TestEncoder_Encode(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		var buf bytes.Buffer
		encoder := entry.NewEncoder(&buf)
		e, _ := entry.NewEntry(
			[]byte("key"),
			[]byte("value"),
		)

		n, err := encoder.Encode(e)

		require.NoError(t, err)
		assert.EqualValues(t, 29, n)
		assert.Len(t, buf.Bytes(), 29)
	})
}
