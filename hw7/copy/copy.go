package copy

import (
	"io"
	"os"
	"time"

	"github.com/cheggaaa/pb/v3"
)

// Copy function copies
func Copy(from string, to string, limit int, offset int) error {
	// create and start new bar
	bar := pb.StartNew(limit)

	var fileFrom *os.File
	fileFrom, _ = os.OpenFile(from, os.O_RDONLY, 0644)
	defer fileFrom.Close()

	var fileTo *os.File
	fileTo, _ = os.OpenFile(to, os.O_RDWR, 0644)
	defer fileTo.Close()

	fileFrom.Seek(int64(offset), io.SeekStart)
	_, err := io.CopyN(fileTo, fileFrom, int64(limit))

	for i := 0; i < limit; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()
	return err
}
