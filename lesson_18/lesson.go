package lesson_18

import (
	"fmt"
	"image/color"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func DoLesson() {
	firstExample()
	secondExample()
	thirdExample()
	fourthExample()
	fifthExample()
	sixthExample()
	seventhExample()
}

// defining a type and showing how we can assign it as a concrete
// implementation to an abstract interface
func firstExample() {
	var v IntSlice = []int{1, 2, 3}

	// this is an interface variable, I can assign this to a concrete type
	var s fmt.Stringer = v

	for i, x := range v {
		fmt.Printf("%d: %d\n", i, x)
	}

	fmt.Printf("%T %[1]v\n", v)

	// this prints the concrete type, not the interface
	fmt.Printf("%T %[1]v\n", s)
}

// creating a user named type from a basic type
type IntSlice []int

// see how this doesn't use any type of `implements` keyword like
// in Java, instead we use duck typing, where, because we provide
// the needed method signature we consider it a concrete implementation
func (is IntSlice) String() string {
	var strs []string
	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}
	return "[" + strings.Join(strs, ";") + "]"
}

func secondExample() {
	// error checking is ignored to make this easier for now
	f1, _ := os.Open("lesson_18/a.txt")
	f2, _ := os.Create("lesson_18/out.txt")

	// takes in two interface concrete types, a reader and writer
	n, _ := io.Copy(f2, f1)

	fmt.Println("copied", n, "bytes")
}

// create our own type that will later "implement" the io.Writer
// interface
type ByteCounter int

// we can define a new type that satisfies the Writer interface,
// again using "duck typing" where we add the needed method
func (b *ByteCounter) Write(p []byte) (int, error) {

	// we need to cast to a ByteCounter because we use +=
	// even though ByteCounter is a int
	*b += ByteCounter(len(p))

	// the writing is supposed to happen here, we do nothing,
	// except print the number of bytes. We also return nil
	// as the error
	return len(p), nil
}

func thirdExample() {
	var c ByteCounter

	f1, _ := os.Open("lesson_18/a.txt")
	f2 := &c

	// notice how we can pass in a reference to our new type
	n, _ := io.Copy(f2, f1)

	fmt.Println("copied", n, "bytes")
	fmt.Println(c)
}

// create a sample type that represents a set of integers
type IntSet struct {
}

// add a method to the type that is essentially a "toString()"
func (*IntSet) String() string {
	return ""
}

func fourthExample() {

	// this will fail, because we did not define the String() for
	// an actual value, only for a reference. Since we didn't
	// define this as a variable we can't get a pointer to it
	// var _ = IntSet{}.String()

	// create an actual variable
	var s IntSet

	// this works because s is a variable, and Go will
	//automatically use &s
	var _ = s.String()

	// this is fine because we pass in a pointer
	var _ fmt.Stringer = &s

	// this is not ok because we cannot overload the String method
	// to take in an actual value
	// var _ fmt.Stringer = s
}

type Point struct {
	X, Y float64
}

// type composition
type Line struct {
	Begin, End Point
}

// define our method on the line type
func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

func fifthExample() {

	// literal that creates the line
	side := Line{Point{1, 2}, Point{4, 6}}
	fmt.Println(side.Distance())
}

type Path []Point

// add a method to the Path type that will compute the distance
func (p Path) Distance() (sum float64) {
	// iterate through the points and sum the distances
	for i := 1; i < len(p); i++ {
		// distance takes a value receiver, so we can create a
		// value on the fly and calculate its distance, then throw
		// the value away. not efficient of course
		sum += Line{p[i-1], p[i]}.Distance()
	}
	return sum
}

// we can create an interface
type Distancer interface {
	Distance() float64
}

func PrintDistance(d Distancer) {
	fmt.Println(d.Distance())
}

// this function takes in a value, so it will not actually modify it
//func (l Line) ScaleBy(f float64) {
//	l.End.X += (f-1)*(l.End.X - l.Begin.X)
//	l.End.Y += (f-1)*(l.End.X - l.Begin.X)
//}

func (l Line) ScaleBy(f float64) Line {
	l.End.X += (f - 1) * (l.End.X - l.Begin.X)
	l.End.Y += (f - 1) * (l.End.X - l.Begin.X)
	return Line{l.Begin, Point{l.End.X, l.End.Y}}
}

func sixthExample() {
	perimeter := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}

	// these both work
	fmt.Println(perimeter.Distance())
	PrintDistance(perimeter)

	// testing ScaleBy
	fmt.Println(Line{Point{1, 2}, Point{4, 6}}.ScaleBy(2).Distance())
}

// distance from another point
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func seventhExample() {
	// create a point and colored point
	p, q := Point{1, 1}, ColoredPoint{Point{5, 4}, color.RGBA{255, 0, 0, 255}}

	l1 := q.Distance(p)
	l2 := p.Distance(q.Point)

	fmt.Println(l1, l2)

	// the following is not allowed, because q is a
	//ColoredPoint, not a Point
	// p.Distance(q)
}
