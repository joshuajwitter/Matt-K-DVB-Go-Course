package lesson_01

import (
	"strings"
)

func Say(names []string) string {

	if len(names) == 0 {
		names = []string{"world"}
	}

	return "Hello, " + strings.Join(names, ", ") + "!"
}
