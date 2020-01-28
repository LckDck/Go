package main

import (
	"os"

	"github.com/LckDck/go/hw8/envdir"
)

func main() {
	args := os.Args
	env, err := envdir.ReadDir("~")
	if err != nil {
		os.Exit(1)
	}
	//arguments := []string{"IGOGO_VAR=igogoshki"}
	code := envdir.RunCmd(args, env)
	os.Exit(code)
}
