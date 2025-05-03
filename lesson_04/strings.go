package lesson_04

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FooString() {
	s := "élite"

	fmt.Printf("s: %8T %[1]v\n", s)

	// slices introduced, this shows how runes are really int 32s
	// []int32 [233 108 105 116 101]
	fmt.Printf("s: %8T %[1]v\n", []rune(s))

	// this shows how runes are really bytes under the covers, the first char
	// is actually represented by two bytes
	// []uint8 [195 169 108 105 116 101]
	fmt.Printf("s: %8T %[1]v\n", []byte(s))

	// this shows us the length of the srting, which surprisingly shows the
	// physical length of the string in bytes, not the number of runes
	b := []byte(s)
	fmt.Printf("%8T %[1]v %d\n", b, len(b))

	// this shows us the length of the string, which surprisingly shows the
	// physical length of the string in bytes, not the number of runes
	fmt.Printf("%8T %[1]v %d\n", s, len(s))

	// you can see how, once hw has been instantiated, you can reference
	// its data in memory without reallocating
	// hw := "hello, world"
	// hello := hw[:5]
	// world := hw[7:]

	// if you do this, you just get a copy of the descriptor (pointer plus
	// length combination) and no duplicate allocation
	t := s

	// this will reallocate
	t += "es"

	// this will also reallocate
	t = strings.ToUpper(t)
}

// Sample output
// joshuawitter@spiedie ~/Intrepid/Code/Go/Learning Go [lesson_4_strings] # go run ./cmd josh gina < lesson_04/tests.txt
// gina went to greece
// where is gina go
// gina went to israel
// gina didn't go there
func BarString() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "not enough args")
		os.Exit(-1)
	}

	// initialize two variables at the same time
	old, new := os.Args[1], os.Args[2]

	// create a scanner on Sdin, so we can scan a file
	scan := bufio.NewScanner(os.Stdin)

	// while scan.scan returns true (this is a for loop essentially)
	for scan.Scan() {

		// look for the target work (old) and split on it
		s := strings.Split(scan.Text(), old)

		// join on the new word
		t := strings.Join(s, new)

		// output the replaced strings
		fmt.Println(t)
	}
}
