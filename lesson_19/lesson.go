package lesson_19

import (
	"fmt"
	"path/filepath"
	"sort"
)

func DoLesson() {
	Example1()
	Example2()
	Example3()
	Example4()
	Example5()
}

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

func (p PairWithLength) String() string {
	return fmt.Sprintf("Hash of %s is %s; length: %d", p.Path, p.Hash, p.Length)
}

// The fields of an embedded struct are promoted to the level
// of the embedding structure
type PairWithLength struct {
	Pair   // this is promoted
	Length int
}

// example of promoting a field
func Example1() {
	p1 := PairWithLength{Pair{"/usr", "0xfdfe"}, 121}
	fmt.Println(p1.Path, p1.Length) // not pl.x.Path

	p := Pair{"/usr", "0xfdfe"}
	pl := PairWithLength{p, 3}

	fmt.Println(p)

	// these work
	fmt.Println(pl.Length)
	fmt.Println(pl.Path)

	// when it goes to resolve which method it uses it looks first for
	// a method for the actual struct, then it looks for one that is
	// promoted
	fmt.Println(pl)

	// keep in mind that PairWithLength is not a subclass of Pair, it
	// "has a" Pair (composition)
}

func Filename(p Pair) string {
	return filepath.Base(p.Path)
}

func Example2() {
	p := Pair{"/usr", "0xfdfe"}
	pl := PairWithLength{p, 133}

	fmt.Println(p)
	fmt.Println(pl)

	fmt.Println(Filename(p))
	// fmt.Println(Filename(pl)) this won't work as PairWithLength is
	// not a Pair (no inheritance)
	fmt.Println(Filename(pl.Pair))
}

// if we had our types implement this interface then we could use it to
// specify common functionality
type FileNamer interface {
	FileName() string
}

type Fizgig struct {
	// the methods of this type will be promoted even though it is
	// a pointer
	*PairWithLength
	Broken bool
}

// embedding a pointer in the struct
func Example3() {
	fg := Fizgig{
		&PairWithLength{Pair{"/usr", "0xfdfe"}, 121}, false,
	}
	fmt.Println(fg)
}

type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ

func (s Organs) Len() int {
	return len(s)
}

func (s Organs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// create our types that will have their own sorting methods
type ByName struct {
	Organs
}
type ByWeight struct {
	Organs
}

// add some methods
func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}
func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

// sorting
func Example4() {
	s := []Organ{{"brain", 1340}, {"liver", 1494}, {"spleen", 162}, {"pancreas", 131}, {"heart", 290}}

	// print unsorted
	fmt.Println(s)

	// sort by weight
	sort.Sort(ByWeight{s})
	fmt.Println(s)

	// sort by name
	sort.Sort(ByName{s})
	fmt.Println(s)

	// sort by reverse name, and example of embedding an interface into
	// a struct
	sort.Sort(sort.Reverse(ByName{s}))
	fmt.Println(s)
}

type StringStack struct {

	// NOTE: see how the `data` field is lowercased. This means that
	// when this is imported from a package, the importing code will
	// not see the field, hiding the internal implementation details

	data []string // "zero" value ready to use
}

func (s *StringStack) Push(x string) {
	s.data = append(s.data, x)
}

func (s *StringStack) Pop() string {
	if l := len(s.data); l > 0 {
		t := s.data[l-1]
		return t
	}
	panic("pop from empty stack")
}

// making nil useful
func Example5() {
	var s StringStack

	// this will cause the panic to happen
	// popped := s.Pop()
	// fmt.Println(popped)

	s.Push("something to pop later")
	popped2 := s.Pop()
	fmt.Println(popped2)
}

type IntList struct {
	Value int
	Tail  *IntList
}

// Important: Nothing in Go prevents calling a method with a nil
// receiver. This would be the equivalent of null.toString() in Java
func (list *IntList) Sum() int {
	// this lets us handle a nil receiver, we know that the list is
	// done when the tail is nil
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
