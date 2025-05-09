package lesson_08

import "fmt"

func Lesson08() {
	m := map[int]int{4: 1, 7: 2, 8: 3}
	do(m)
	fmt.Println("m", m)
	doWithPointer(&m)
	fmt.Println("m", m)
	doSomethingWithDefer()
	doAnotherThingWithDefer()
	result := doAThirdThingWithDefer()
	fmt.Println(result)
}

// this is going to pass by value
func do(m1 map[int]int) {
	m1[3] = 0
	m1 = make(map[int]int)
	m1[4] = 4
	fmt.Println("ml", m1)
}

// this is going to "pass by reference", but in reality this just means that the value of the descriptor is passed
func doWithPointer(m1 *map[int]int) {
	(*m1)[3] = 0

	// in this case the literal reference in the Lesson08 function
	// changes because the descriptor changes
	*m1 = make(map[int]int)
	(*m1)[4] = 4
	fmt.Println("ml", *m1)
}

// nothing happens until the function exits
func doSomethingWithDefer() {
	fmt.Println("This is the first statement")
	defer fmt.Println("This is the first deferred statement")
	fmt.Println("This is the second statement")
	defer fmt.Println("This is the second deferred statement")
	fmt.Println("This is the third statement")
}

// unlike a closure, defer copies arguments to the deferred call
func doAnotherThingWithDefer() {
	a := 10

	// note that the parameter gets copied at the time that the
	// at the defer statement creation
	defer fmt.Println(a)
	a = 11
	fmt.Println(a)
}

// adding a named return value with a naked return
func doAThirdThingWithDefer() (a int) {
	defer func() {
		a = 2
	}()

	a = 1
	return
}
