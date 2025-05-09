package lesson_07

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Lesson07() {
	printingSomeFormats()
	doSomeFileIO()
	doSomeFileIOBetter()
	doSomeWordCounting()
}

func printingSomeFormats() {
	s := "a string"
	b := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)
	fmt.Printf("%v\n", string(b))
}

// you can run this like this: go run ./cmd lesson_07/sample1.txt lesson_07/sample2.txt
func doSomeFileIO() {
	for _, fname := range os.Args[1:] { // ignore first arg, name of program
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if _, err := io.Copy(os.Stdout, file); err != nil {
			if _, writeErr := fmt.Fprintln(os.Stderr, err); writeErr != nil {
				// maybe log to a file, or just ignore
				log.Printf("failed to write error to stderr: %v", writeErr)
			}
			continue
		}

		file.Close()
	}
}

func doSomeFileIOBetter() {
	for _, fname := range os.Args[1:] { // ignore first arg, name of program
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		data, err := io.ReadAll(file)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		fmt.Println("The file has", len(data), "bytes")

		file.Close()
	}
}

func doSomeWordCounting() {
	for _, fname := range os.Args[1:] {
		var lc, wc, cc int

		file, err := os.Open(fname)

		if err != nil {
			fmt.Println(err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()
			wc += len(strings.Fields(s)) // this will split a string using a space or tab
			cc += len(s)                 // this is only true for one char runes
			lc++
		}
		fmt.Printf("%7d lines %7d words %7d chars for file %s\n", lc, wc, cc, fname)
	}
}
