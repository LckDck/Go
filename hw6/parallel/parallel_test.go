package parallel

import (
	"errors"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	tasks := make([]func() error, 5)

	tasks = append(tasks, func() error {
		time.Sleep(1 * time.Second)
		return nil
	})
	tasks = append(tasks, func() error {
		time.Sleep(2 * time.Second)
		return nil
	})
	tasks = append(tasks, func() error {
		time.Sleep(3 * time.Second)
		return nil
	})
	tasks = append(tasks, func() error {
		time.Sleep(4 * time.Second)
		return errors.New("Ошибка 5")
	})
	tasks = append(tasks, func() error {
		time.Sleep(5 * time.Second)
		return nil
	})
	Parallel(tasks, 3, 2)
}
