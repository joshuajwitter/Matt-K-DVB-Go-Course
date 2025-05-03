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
- 