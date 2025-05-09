package lesson_06

import (
	"fmt"
	"os"
)

func Lesson06Main() {

	fmt.Println("Checking the value...")

	switch a := getA(); a {
	case 0, 1, 2:
		fmt.Println("underflow possible")

	case 3, 4, 5, 6, 7, 8:

	default:
		fmt.Println("warning: overload")
	}

	a := getA()

	switch {

	// see how we are able to use the short declaration operator?
	// normally we cannot redeclare, but it's out of scope as "a" is only
	// in the switch scope
	case a <= 2:
		fmt.Println("a less than or equal to 2")

	default:
		fmt.Println("a is greater than 2")
	}
}

func getA() int {
	return 2
}

// example of short declaration shadowing
func BadRead(f *os.File, buf []byte) error {
	var err error
	for {
		n, err := f.Read(buf) // shadows 'err' above
		if err != nil {
			break
		}
		fmt.Printf("doing foo %d\n", n)
	}

	return err // will always be nil
}

func testingSomeStructuralTyping() {

	// we can see this with arrays (note, slices are going to be more lenient)
	a := [...]int{1, 2, 3}
	b := [3]int{}

	a = b // this is ok

	c := [4]int{}

	fmt.Printf("a is %a, c is %d\n", a, c)

	//a = c // this is not ok
}
