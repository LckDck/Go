package main

import (
	"github.com/LckDck/go/hw7/copy"
	flag "github.com/spf13/pflag"
)

var offset *int64 = flag.Int64("offset", 0, "offset before copy")
var limit *int64 = flag.Int64("limit", 0, "limit of copy size")

func main() {
	flag.Parse()
	copy.Copy("from.txt", "to.txt", *limit, *offset)
}
