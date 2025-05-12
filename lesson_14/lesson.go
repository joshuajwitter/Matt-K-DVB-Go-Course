package lesson_14

import (
	"fmt"
	"sync"
)

func DoLesson() {
	// Each append copies a reference to i tem with its last value, so we need to copy each
	// item into a (non-empty) slice
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	a := [][]byte{}
	for _, item := range items {

		// this is required otherwise you will only see the last value
		// [[7 8] [7 8] [7 8] [7 8]]
		i := make([]byte, len(item))
		copy(i, item[:]) // make unique
		a = append(a, i)
	}
	fmt.Println(items) // [[1 2] [3 4] [5 6] [7 8]]
	fmt.Println(a)     // [[1 2] [3 4] [5 6] [7 8]]
}

type Employee struct {
	mu   sync.Mutex
	Name string
}

// wait group, a concurrency structure
func do(emp *Employee) {
	emp.mu.Lock()
	defer emp.mu.Unlock()
	// ...
}
