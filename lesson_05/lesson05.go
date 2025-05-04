package lesson_05

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func CompositeTypes() {
	t := []byte("string")

	fmt.Println(len(t), t)

	// indexed at zero
	fmt.Println(t[2])

	// see how this does not include the last value
	fmt.Println(t[:2])

	// last 4 values
	fmt.Println(t[2:])

	// see how this does not include the last value
	fmt.Println(t[3:5], len(t[3:5]))

	w := [...]int{1, 2, 3} // this cannot be mutated inside a function
	x := []int{0, 0, 0}    // this can as it is a slice passed by reference

	y := do(w, x)
	fmt.Printf("%8T, %8T, %8T\n", w, x, y)
	fmt.Println(w, x, y)

	// you can create a slice from an array
	z := x[0:1]

	fmt.Println(z)

	mapTime()

	countWordsMain()
}

// warning: this function has side effects! :-P
func do(a [3]int, b []int) []int {
	// a = b cannot do this conversion
	a[0] = 4 // this will not change the actual array (not passed by ref)
	b[0] = 3 // this will change the slice (passed by ref)
	// b is now 3,0,0, which changes x

	c := make([]int, 5) // makes a new array of size 5
	// c is now 0,0,0,0,0

	c[4] = 42 // set the value of the 4th slice element
	// c is now 0,0,0,0,42

	// copies b into c starting at the beginning
	copy(c, b) // mutate the slice that was passed by reference
	// c is now 3,0,0,0,42

	return c
}

func mapTime() {
	var m map[string]int // maps a string to an int, nil, no storage

	// use make because we cannot insert into a nil map
	p := make(map[string]int) // non-nil but empty
	a := p["the"]             // missing key returns zero
	b := m["the"]             // missing key returns zero
	//m["and"] = 1              // panic as we have a nil map
	m = p         // this makes m the descriptor of p, we can insert now
	m["and"]++    // OK, is the same as p now
	c := p["and"] // returns 1

	fmt.Println(a, b, c)

	var j = map[string]int{
		"and": 0,
		"the": 1,
		"or":  2,
	}

	// checking for the index of a key (does it exist?)
	k, ok := j["and"]

	// ok will be true if the key exists in the map
	fmt.Println(k, ok)

	// checking for the index of a key (does it exist?)
	k, ok = j["non-existent-key"]

	// ok will be false if the key does not exist in the map
	fmt.Println(k, ok)

	// common helpful pattern
	if k, ok = j["and"]; ok {
		// we know that k is not the default value (zero)
	}
}

// can test this with: go run ./cmd < lesson_04/tests.txt
func countWordsMain() {

	// create the scanner for std in
	scan := bufio.NewScanner(os.Stdin)

	// create a map that will hold the words mapped to their number of
	// instances, use `make` because inserting into a nil map panics
	words := make(map[string]int)

	// set up the scanner to split on words
	scan.Split(bufio.ScanWords)

	// scan in each word and increment the value in the map for that word
	// each time it is found
	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "unique words")

	// create a key value pair struct that we can use to hold values
	type kv struct {
		key string
		val int
	}

	// create a slice of key value pairs, we can add them to this
	var ss []kv // since we do not define the size, this is a slice
	// but won't we get an error when we try to append here (no make)?

	// range over the map and i get the k and v, which i make the slice of
	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	// pass in a function that will be used to do the comparison
	// this is like a java comparator, this is a lambda expression and a
	// closure
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	// _ means ignore the value, in this case the index
	for _, s := range ss {
		fmt.Println(s.key, "appears", s.val, "times")
	}

	fmt.Println("just the top three!")

	// just do the top three
	for _, s := range ss[0:3] {
		fmt.Println(s.key, "appears", s.val, "times")
	}
}
