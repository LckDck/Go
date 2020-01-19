package copy

import (
	"io"
	"os"
)

// Copy function copys
func Copy(from string, to string, limit int64, offset int64) error {
	var fileFrom *os.File
	fileFrom, _ = os.OpenFile(from, os.O_RDONLY, 0644)
	defer fileFrom.Close()

	var fileTo *os.File
	fileTo, _ = os.OpenFile(to, os.O_RDWR, 0644)
	defer fileTo.Close()

	fileFrom.Seek(offset, io.SeekStart)
	_, err := io.CopyN(fileTo, fileFrom, limit)
	return err
}
