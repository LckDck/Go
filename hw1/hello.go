package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("")
	fmt.Println("Time: " + time.String() + "\n")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error! "+err.Error())
		os.Exit(3)
	}
}
