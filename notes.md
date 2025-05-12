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

Lesson 5:
- Composite types
- Arrays
  - These are more rare
  - Arrays have the fixed size defined at compile time
  - When you do "d = b" that means you are copying the elements
  - Passed by value
  - Comparable
  - Can be used as a map key
  - Useful as "pseudo" constants
  - Example: starting array for a DES encryption starting state
- Slices
  - Passed by reference
  - Not comparable
  - Cannot be used as a map key
  - Slice is a variable length array (like a vector)
  - They are backed by an array
  - Works like a string, you have a descriptor that points to a data component and a current length and capacity
  - Doing `a = b` overwrites a, and changes the descriptor of a to be the same as b, pointing to the same data, length and capacity
  - Slices are indexed like `[8:11]`
    - Read as the starting element and one past the ending element, so we have 11-8=3 elements in our slice
  - Beware the "off by one / fencepost" bug  
  - Useful as function parameters
  - Use `make` to make a slice of 
  - Strings are immutable, slices are mutable
  - If you go over the capacity (cap) of a slice it will need to be reallocated, but don't worry, it gets reallocated to twice the size
  - When you append you always reassign the result of append, like this: `c=append(c,d)`
- Maps (keys to values)
  - Not consecutive in memory
  - Maps are dictionaries, indexed by key, returning a value
  - You can read from a nil map, but inserting will panic
  - Passed by reference, no copying, updating ok
  - They type used for the key must have == and != defined (not slices, maps or funcs)
  - Map literals allow us to defined a map at initialization time
  - There are two ways to index a map, one that provides the value and one that lets you know if the key exists
  - You can delete keys with `delete(m,k)` which succeeds even if the key does not exist
- Make nil useful
  - len, cap and range can all be called on nil values 
  - The idea that we can make nil safe and useful is really helpful
  - Essentially nil maps are read-only, empty maps

Lesson 5a: "Understanding nil" by Francesc Campoy
- https://youtu.be/ynoY2xz-F8s?si=6_IyzvfKN-YnI9lm
- nil: nothing
- null: not any
- nil is a kind of zero
- nil is not a keyword and you can define it
- nil has no type
  - You cannot defined a variable with type nil
- It is the zero value for many types in Go
- pointers in Go point to a position in memory
  - Similar to C++ but you cannot do pointer arithmetic (memory safety)
    - unless you use `unsafe`
  - A nil pointer points to nothing
- Slice has a pty, len, and cap
- Interfaces
  - An interface is not a pointer
  - It has a type and a value
  - So the value is nil
  - When is nil not nil?
    - ```go
      func do() error {
        var err *doErr
        return err // this is an interface that points to nil
      }
    
      func main() {
        err := do() // err will not be ni, it will just point to nil
        fmt.Println(err == nil)
    }
    ```
  - The lessons here: 
    - Do not declare concrete error vars
    - Do not return concrete error types
  - nil is different depending on the type that is referencing nil
    - nil pointers point to nothing
    - nil slices have no backing array
    - nil maps, functions and channels ar enot initialized
    - nil interfaces have no value assigned, not even a nil pointer
  - How are nil values used?
    - pointers
      - you can make a pointer to an int, for example, and set it to nil
      - dereferencing it leads to a panic `*p`
  - See the code example in lesson05a.go
    - Try to make your default values useful. nil can be more useful than 0, the default value returned for most types
    - I learned about "method receiver" functions, which are like extension in Kotlin or Swift, see the Sum() function

Lesson 6:
- All if-then statements require braces
- You can use short assignment operators in if statements
- for loops work just like C
- Go only has one loop, the for loop
- There is a range operator that allows you to range over a slice or a map, remember it gives you one or two values, if you want the index or not
- Maps are unordered since they are based off of a hashtable, so if you get the keys they will be in a random order. This is why Go is so fast, it doesn't order things like this and makes you do it if you need to
- break and continue work as you would expect
- You can use loop labels, such as "outer:" to conditionally break out of inner loops, I have seen this in other languages
- switch statements are available as syntactic sugar to replace more complicated for loops. 
- There is no break statement and there is no fallthrough, this makes everything really easy to understand
- Go does not make you cover every enumerated case (there are no actual enums in Go)
- Short declaration operator can only be used in functions
- You can declare anything at package scope
- Every name that is capitalized is exported to every other package, otherwise it is private to the package
- Every file in a package must include the things it needs, unused imports are an error
- Cyclic dependencies are prohibited (move common dependencies to a third package or eliminate them think tree structure, not graph (a tree is a graph with no cycles), this speeds up the compile time and how the program is initialized
  - Variables are often setup before `main` is executed, and we would not know how to order them
- Packages as a concept is really about information hiding
  - A good package encompasses deep functionality behind a simple API
  - The Unix file API is perhaps the best example of this model
  - There are only 5 functions but there is so much complexity that they hide
- Declarations
  - You can declare `const` variables
  - You can declare `types`
  - If you do not specifically use a type, the compiler must be able ot infer the type at compile time from the assigned value
- "Structural Typing"
  - See code example in lesson06.go
  - "Duck typing" (if it looks like a duck, and quacks like a duck, it is a duck)
  - It's the same type if it has the same structure or behavior:
    - arrays of the same size and base type
    - slices with the same base type
    - maps of the same key and value types
    - structs with the same sequence of field names/types
    - functions with the same parameter & return types
- Go keeps "arbitrary" precision for literal values (256 bits or more)
  - Integer literals are untyped
  - assign a literal to any size integer without conversion
  - assign an integer literal to float, complex also
  - Ditto float and complex; picked by syntax of the literal 2.0 or 2e9 or 2.0i or 213
  - Mathematical constants can be very precise
  Pi = 3.14159265358979323846264338327950288419716939937510582097494459
  - Constant arithmetic done at compile time doesn't lose precision
- Only one true overloaded operator in Go, which is the `+`, which adds numbers or concats strings
- Again, no truthy or falsy (thank you God)

Lesson 7:
- Uses the standard three streams and the basic format codes that we are already familiar with (%s, %d, etc)
- %T is the type, %v is the value, we used this earlier
- This lesson focused on printing things to various outputs, the code is mainly in lesson07.go, but it's pretty basic
- io/ioutil is really useful for reading an entire file for example
- strconv has other utilities to parse strings to numbers for example, like parsing CSV

Lesson 8:
- Functions
  - Functions are “first class” objects; you can:
    - Define them — even inside another function
    - Create anonymous function literals
    - Pass them as function parameters / return values
    - Store them in variables
    - Store them in slices and maps (but not as keys)
    - Store them as fields of a structure type
    - Send and receive them in channels
    - Write methods against a function type
    - Compare a function var against nil
  - Parameters
    - In theory, it seems that some datatypes are pass by reference, and some passed by value
      - But that's not actually true. All parameters are passed by copying something (i.e. by value)
      - If the thing being copied is a pointer or descriptor, then the shared backing store (array, hashtable, etc) can be changed through it, thus we think of it as "by reference"
      - Remember, slice and map are descriptors. You are copying the DESCRIPTOR and not the actual reference. A slice or map CONTAINS a reference to the underlying data
  - Recursion
    - We have an example of this previously with the tree walking
  - Call Stack is the same as usual
  - Deferred Execution
    - The `defer` statement captures a function call to run later
      - This means we can write less code and cleaner code
      - So if we add `defer f.Close()` to any point in the function, at the end of the function's execution we will run the deferred closing
      - If we have more than one, they will run in LIFO order
      - Always operates on a function scope, not a block scope
  - Named return values can create a defer gotcha, see the code example
    - You can do something like `doIt() (a int)` which automatically creates a variable in the scope of the function that you can manipulate. You can then have a "naked return" which is a simple "return" without anything after it, this returns the named variable. This is not recommended by the teacher because it is not easy to understand

Lesson 9:
- Closures
  - Scope vs. Lifetime
    - Scope
      - Static, it's a compile time thing
        - It's about who can reference the thing 
    - Lifetime
      - Depends on program execution (runtime)
      - See the code example, once the compiler sees that a reference is used outside of a scope it does an "escape analysis" and that variable will have a lifetime as long as the function, and will put this value onto the heap
      - So it is possible for a variable to have a lifetime that exceed the lifetime of its scope
  - A closure is like a descriptor, see the example code
    - It has a function and an environment (the variables it closes over)
    - This is pretty much the same as I previously understood closures
  - You can easily make a slice of functions for example, see the code example

Lesson 10: Slices in more details
- Nil slice versus empty slice
  - See the code example, this can be useful when translating slices to json
  - Don't use `if a == nil` instead use `if len(a) == 0`
    - You can always ask the length of a nil slice
- Length vs capacity and the slice operator
  - When you make a slice that uses the two argument slice operator to slice an existing slice, the underlying capacity of the new slices is the same as the original slice, see from the code example, there is a special three argument operator to make a new slice with a set capacity
- See the code for the rest of the lesson

Lesson 11: Homework #2
- Exercise 5.5 from GOPL: implement `countWordsAndImages`
  - Given some HTML as raw text, parse it into a document and then call a counting routine to detect and count words and images
  - See the code for my solution

Further study:
- Hash tables, confirm how they work
- Rob Pike
  - Go Proverbs
- Matt Holiday
- Francesc Campoy (understanding nil)
- Compare a lambda expression to a closure
- Read the Go language specification

Weird things about Go:
- You cannot insert into a nil map, but you can read from it (always returns the default value of the map's datatype)

Common helpful patterns
- Checking to see if a key in a map is just a default or not
  - if k, ok = j["and"]; ok {
      // we know that k is not the default value (zero)
    }