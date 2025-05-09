package lesson_09

import "fmt"

func Lesson09() {
	x := doIt()
	fmt.Println("x = ", *x)

	// this is a function that has a state (the variables it
	// closes over)
	fibFunction1 := fib()
	fibFunction2 := fib()

	fibFunction2()

	// note that the two closures have their own states
	// Note: I found an overflow!
	for j := 0; j < 100; j++ {
		fmt.Println(fibFunction1(), fibFunction2())
	}

	slicedFuncs()
}

func doIt() *int {
	var b int
	b = 3

	// once the compiler sees this is does "escape analysis" and it
	// will have a lifetime as long as the function, and will put this
	// value onto the heap, which is inefficient
	return &b
}

// this is a function that returns a function that returns an int
// so calling this function gives you a new function that you can call
// to get Fibonacci numbers
func fib() func() int {
	// first we set the first two values of the sequence
	a, b := 0, 1
	// then we return an anonymous function (a function without a name) that does the actual work. this is called a closure because
	// it "closes over" the variables a and b, remembering their
	// values between calls
	return func() int {

		// every time we call the inner function, it updates a and b
		// to the next pair of Fibonacci numbers. this saves the state
		// using the variables the function closes over, sliding the
		// values up one Fibonacci number at a time
		a, b = b, a+b

		// this returns the next sequence number (a+b)
		return b
	}
}

// this used to reinstantiate on every iteration
// this was changed in more recent versions of Go :-)
func slicedFuncs() {
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
	}

	for i := 0; i < 4; i++ {
		s[i]()
	}
}
