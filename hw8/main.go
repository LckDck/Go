package main

import (
	"fmt"
	"os"

	"github.com/LckDck/go/hw8/envdir"
)

func main() {
	args := os.Args
	// _, err := envdir.ReadDir("~")
	// if err != nil {
	// 	os.Exit(1)
	// }

	myenv := map[string]string{}
	myenv["IGOGO_VAR"] = "igogoshki"
	code := envdir.RunCmd(args, myenv)

	check, exist := os.LookupEnv("IGOGO_VAR")
	if exist {
		fmt.Println("!! exist ")
	}
	fmt.Println("!! " + check)

	os.Exit(code)
}
