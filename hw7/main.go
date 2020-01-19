package main

import (
	"os"

	"github.com/LckDck/go/hw7/copy"
	flag "github.com/spf13/pflag"
)

var offset *int = flag.Int("offset", 0, "offset before copy")
var limit *int = flag.Int("limit", 0, "limit of copy size")

func main() {
	flag.Parse()
	from := os.Args[1]
	to := os.Args[2]
	copy.Copy(from, to, *limit, *offset)
}
