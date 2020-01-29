package envdir

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
)

// ReadDir reads the dir
func ReadDir(dir string) (map[string]string, error) {
	m := make(map[string]string)
	env := os.Environ()
	for _, value := range env {
		paar := strings.Split(value, "=")
		if len(paar) != 2 {
			return m, errors.New("Error: wrong environment variable format")
		}
		m[paar[0]] = paar[1]
		//fmt.Println(paar[0] + " = " + paar[1])
	}
	return m, nil
}

// RunCmd runs the cmd
func RunCmd(cmd []string, env map[string]string) int {
	if len(cmd) == 0 {
		return 1
	}

	args := make([]string, len(cmd))
	if len(cmd) > 1 {
		args = cmd[1:]
	}

	command := exec.Command(cmd[0], args...)

	res := make([]string, len(env))
	for k, v := range env {
		res = append(res, k+"="+v)
	}
	command.Env = res

	if err := command.Run(); err != nil {
		log.Fatal(err)
	}
	return 0
}
