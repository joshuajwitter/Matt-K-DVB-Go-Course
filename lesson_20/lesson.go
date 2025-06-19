package lesson_20

import (
	"fmt"
	"math"
)

func DoLesson() {
	PartOne()
	PartTwo()
	PartThree()
}

// standard to prefix errors with "err"
type errFoo struct {
	err  error
	path string
}

func (e errFoo) Error() string {
	return fmt.Sprintf("%s: %s", e.path, e.err)
}

// returns a concrete pointer to an errFoo
func XYZ(a int) *errFoo {
	return nil
}

// the appropriate way to do this would really be:
//func XYZ(a int) error {
//	return nil
//}

func PartOne() {

	// BAD: interface gets a nil concrete ptr
	var err error = XYZ(1)

	// err is no not nil, even though it has a nil pointer inside
	if err != nil {
		// you would think this would not execute. but it will, because
		// err is not nil, because the function returns a concrete
		// pointer to an errFoo, which is a mistake
		fmt.Println("oops")
	} else {
		fmt.Println("OK")
	}
}

// This is a simple function that takes two integers, a and b, and returns their sum
// Example: Add(1, 2) returns 3
// Nothing special here—it’s just a normal function that needs both arguments at once
func Add(a, b int) int {
	return a + b
}

// This is where currying happens...
// AddToA takes one argument (a): Instead of taking both a and b at
// once, it only takes a.
// It returns a new function: The return type is func(int) int, which
// means it returns a function that takes an integer (b) and returns an
// integer (the sum)..
// The returned function "remembers" a... inside the returned function,
// a is still available because of Go’s closures (a fancy term for a
// function that can access variables from its outer scope). This inner
// function uses a and waits for b to call Add(a, b)
func AddToA(a int) func(int) int {

	// returns a function that closes over a
	return func(b int) int {
		return Add(a, b)
	}
}

// currying example
func PartTwo() {

	// Currying is a technique where a function that takes multiple
	// arguments (like a and b) is transformed into a series of
	// functions, each taking one argument at a time. Instead of
	// calling a function like Add(1, 2) all at once, currying lets
	// you "partially apply" the function by giving it one argument
	// first, which returns a new function that waits for the next
	// argument.

	// returns the function that closes over `a` (in this case 1)
	addTo1 := AddToA(1)

	// these are the same
	fmt.Println(Add(1, 2) == addTo1(2))

	// this will return 4
	fmt.Println(addTo1(3))
}

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// method values
func PartThree() {

	p := Point{1, 2}
	q := Point{4, 6}

	// this is a method value that can be called later
	// see how it binds the receiver p
	distanceFromP := p.Distance

	// these are the same, but the first one has one less parameter
	result1 := distanceFromP(q)
	result2 := p.Distance(q)

	fmt.Println(result1)
	fmt.Println(result2)

	fmt.Printf("%T\n", Point.Distance)

	// what happens if you change p?
	p = Point{2, 2}

	// you will get the same result as we closed over the *value*
	// normally closures close over references, but since Distance
	// takes in a *value* param this is what happens
	fmt.Println(distanceFromP(q))
}
