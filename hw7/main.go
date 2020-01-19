package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/LckDck/go/hw7/copy"
)

func main() {
	tasks := []func() error{	
	copy.Copy("from.txt", "to.txt", 0, 0)
}
