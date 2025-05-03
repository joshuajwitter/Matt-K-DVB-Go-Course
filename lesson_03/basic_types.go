package lesson_03

import (
	"fmt"
	"os"
)

func DoSomeBasicTypes() {
	a := 2
	b := 2.1

	fmt.Printf("a: %8T %[1]v\n", a)
	fmt.Printf("b: %8T %[1]v\n", b)
}

// you can do this: cat nums.txt | go run ./cmd
// or go run ./cmd < nums.txt
func DoOtherThings() {
	var sum float64
	var n int

	for {
		var val float64

		// &val is a pointer to the value that was read in
		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			break
		}

		sum += val
		n++
	}

	if n == 0 {
		// same as println, but i will tell you which output stream, in this
		// case os,Stderr
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	// ctrl-d will tell the computer that it is the end of input
	fmt.Println("The average is", sum/float64(n))
}
