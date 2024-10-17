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

	t.Run("should error when another process is running", func(t *testing.T) {
		tmpDir := t.TempDir()
		dataDir := filepath.Join(tmpDir, "librakv")
		db1, err := librakv.Open(dataDir)

		require.NoError(t, err)
		require.NotNil(t, db1)

		db2, err := librakv.Open(dataDir)

		require.Error(t, err)
		require.Nil(t, db2)
	})
}
