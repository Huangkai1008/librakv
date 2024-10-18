package librakv_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Huangkai1008/librakv"
)

func TestOpen(t *testing.T) {
	t.Run("should succeed", func(t *testing.T) {
		t.Run("with default options", func(t *testing.T) {
			tmpDir := t.TempDir()
			dataDir := filepath.Join(tmpDir, "librakv")

			db, err := librakv.Open(dataDir)

			require.NoError(t, err)
			require.NotNil(t, db)
		})

		t.Run("with options", func(t *testing.T) {
			tmpDir := t.TempDir()
			dataDir := filepath.Join(tmpDir, "librakv")

			db, err := librakv.Open(dataDir, librakv.WithLogFileThresholdSize(1<<20))

			require.NoError(t, err)
			require.NotNil(t, db)
		})
	})

	t.Run("should error when wrong options", func(t *testing.T) {
		tmpDir := t.TempDir()
		dataDir := filepath.Join(tmpDir, "librakv")

		var tests = []struct {
			name string
			opts []librakv.Option
		}{
			{"zero log file threshold size", []librakv.Option{librakv.WithLogFileThresholdSize(0)}},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				db, err := librakv.Open(dataDir, tt.opts...)

				require.Nil(t, db)
				require.Error(t, err)
			})
		}
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
