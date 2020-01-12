package parallel

import (
	"fmt"
)

// Parallel function runs maximum @parallelLimit tasks parallel and stops after @errorsLimit errors
func Parallel(tasks []func() error, parallelLimit int, errorsLimit int) {
	ch := make(chan error, errorsLimit)

	activeCount := 0
	for i := 0; i < len(tasks); {
		if activeCount <= parallelLimit {
			activeCount++
			go func(t func() error, activeCount *int) {
				res := t()
				(*activeCount)--
				if res != nil {
					ch <- res
				}
			}(tasks[i], &activeCount)
			i++
		}
	}

	errorsCount := 0
	for err := range ch {
		errorsCount++
		fmt.Println(err)
		if errorsCount >= errorsLimit {
			close(ch)
			break
		}
	}
}
