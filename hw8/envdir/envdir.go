package envdir

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// ReadDir reads the dir
func ReadDir(dir string) (map[string]string, error) {
	// file, err := os.Open(dir)
	// if err != nil {
	// 	return nil, errors.New("Error: unknown path")
	// }
	m := make(map[string]string)
	env := os.Environ()
	for _, value := range env {
		paar := strings.Split(value, "=")
		if len(paar) != 2 {
			return m, errors.New("Error: wrong environment variable format")
		}
		m[paar[0]] = paar[1]
		fmt.Println(paar[0] + " = " + paar[1])
	}
	return m, nil
}

// RunCmd rund the cmd
func RunCmd(cmd []string, env map[string]string) int {
	return 0
}
