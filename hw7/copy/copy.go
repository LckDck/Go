package copy

import (
	"fmt"
	"io"
	"os"
)

// Copy function copys
func Copy(from string, to string, limit int, offset int) error {
	var fileFrom *os.File
	fileFrom, errFrom := os.OpenFile(from, os.O_RDONLY, 0644)
	defer fileFrom.Close()

	var fileTo *os.File
	fileTo, errTo := os.OpenFile(to, os.O_RDWR, 0644)
	defer fileTo.Close()

	fmt.Println(errTo.Error())
	fmt.Println(errFrom.Error())

	// buffer := make([]byte, limit)
	_, err := io.Copy(fileTo, fileFrom)
	return err
}
