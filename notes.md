Lesson 1:
- Go variables are machine native like C

Lesson 2:
- Default number type is int, and it is represented by the number of bits for the environment
- Don't represent money with floating point, use Go "money" library

Lesson 3:
- "Short declaration operator" (:=) is only allowed inside a function, like c := 2
- %s, %T (which shows the type) are called the "verbs"
- Boolean values are not convertible to integers, no truthiness (thank you god)
- A pointer that points to nil literally points to nothing
- There is a package called "unsafe" where you can do C-like pointer manipulation
- All variables in Go are initialized when they are created
  - All numbers are zero
  - bool gets false
  - string gets "" (length zero)
  - Everything else gets nil
- Constants
  - Can only be a number, string or boolean
  - Are immutable (of course) given concurrency of Go
- You must cast explicitly (thank god)

Lesson 4:
- Strings are immutable (like Java)
- "Rune" is the equivalent of a character
  - It is a 32 bit int
- A string is a UTF-8 sequence of Unicode characters
- The length of a string is the number of bytes that are needed to store the string, not necessarily the number of runes in the string
- When you define a string in go, you give it a pointer to the data and a length in bytes. This is in contrast to other languages that have a null byte where you would need to traverse the whole string to get the null byte
- You can use `utf8.RuneCountInString` to determine the number of runes in a string
- Strings are passed by reference, thus they are not copied