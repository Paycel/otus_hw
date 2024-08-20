package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	src := "testdata/input.txt"
	dst := "tmp/out.txt"
	t.Run("simple test", func(t *testing.T) {
		defer os.Remove(dst)
		tests := []struct {
			name          string
			offset, limit int64
		}{
			{"full copy", 0, 0},
			{"full copy with offset", 50, 0},
			{"offset 50 limit 100", 50, 100},
			{"offset 100 limit 100", 100, 100},
		}
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				err := Copy(src, dst, tc.offset, tc.limit)
				require.NoError(t, err)
			})
		}
	})

	t.Run("invalid input file", func(t *testing.T) {
		defer os.Remove(dst)
		err := Copy("", dst, 0, 0)
		require.Error(t, err)
	})

	t.Run("invalid arguments", func(t *testing.T) {
		defer os.Remove(dst)
		err := Copy(src, dst, -1, 0)
		require.ErrorIs(t, err, ErrInvalidArguments)

		err = Copy(src, dst, 0, -10)
		require.ErrorIs(t, err, ErrInvalidArguments)

		err = Copy(src, dst, 1_000_000_000, 0)
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize)
	})

	t.Run("limit exceeds file size", func(t *testing.T) {
		defer os.Remove(dst)
		err := Copy(src, dst, 0, 1_000_000_000)
		require.NoError(t, err)
	})
}
