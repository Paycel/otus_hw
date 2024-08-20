package main

import (
	"errors"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

var bar = pb.StartNew(5)

func Copy(fromPath, toPath string, offset, limit int64) error {
	srcFile, err := os.OpenFile(fromPath, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("open source file: %w", err)
	}
	bar.Increment()

	srcStat, err := srcFile.Stat()
	if err != nil {
		return fmt.Errorf("get src file stat: %w", err)
	}
	bar.Increment()

	if srcStat.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	dstFile, err := os.OpenFile(toPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("create dst file: %w", err)
	}
	bar.Increment()

	_, err = srcFile.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("seek: %w", err)
	}
	bar.Increment()
	if limit+offset > srcStat.Size() {
		limit = srcStat.Size() - offset
	} else if limit == 0 {
		limit = srcStat.Size()
	}

	if _, err := io.CopyN(dstFile, srcFile, limit); err != nil {
		return fmt.Errorf("copy to file: %w", err)
	}
	bar.Increment()
	return nil
}
