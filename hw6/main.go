package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/LckDck/go/hw6/parallel"
)

func main() {

	tasks := []func() error{
		func() error {
			time.Sleep(2 * time.Second)
			fmt.Println("2")
			return errors.New("ошибка 2")
		},
		func() error {
			time.Sleep(3 * time.Second)
			fmt.Println("3")
			return errors.New("ошибка 3")
		},
		func() error {
			time.Sleep(1 * time.Second)
			fmt.Println("1")
			return nil
		},
		func() error {
			time.Sleep(4 * time.Second)
			fmt.Println("4")
			return errors.New("ошибка 4")
		},

		func() error {
			time.Sleep(5 * time.Second)
			fmt.Println("5")
			return errors.New("ошибка 5")
		}}

	parallel.Parallel(tasks, 2, 3)
}
