package main

import (
	"fmt"

	"github.com/LckDck/Go/hw5/list"
)

func main() {
	list := list.List{}
	list.PushFront("первыйнах")
	fmt.Println(list.Len())
	fmt.Println(list.First())
	fmt.Println(list.Len())
}
