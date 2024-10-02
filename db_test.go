package librakv_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Huangkai1008/librakv"
)

func TestOpen(t *testing.T) {
	t.Run("should succeed without error", func(t *testing.T) {
		t.Run("with default option", func(t *testing.T) {
			tmpDir := t.TempDir()

			dataDir := filepath.Join(tmpDir, "librakv")
			db, err := librakv.Open(dataDir)

			require.NoError(t, err)
			require.NotNil(t, db)
		})
	})
}
