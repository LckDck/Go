package main

import (
	"github.com/LckDck/Go/hw5/list"
)

func main() {
	list := list.List{}
	//list.PushFront("один")
	list.PushBack("два")
	list.PushBack("два")
	list.PushBack("два")
	//list.PushFront("три")
	//fmt.Println(list.String())
}
