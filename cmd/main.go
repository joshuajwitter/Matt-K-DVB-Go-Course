package main

import (
	"hello"
)

func main() {
	hello.DoOtherThings()
	// fmt.Println(hello.Say(os.Args[1:])) // don't include the first item as it is the program name
}
