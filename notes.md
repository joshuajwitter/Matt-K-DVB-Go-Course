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
    }```
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

Lesson 12: Structs, struct tags and JSON
- Code shows everything
- There is a limitation on maps. Maps store their underlying elements in a random order that can change, especially when the map is updated with new elements. So trying to hold a pointer to an element of a map is not a good idea. Therefore, if you want to store a map of structs you should actually store a pointer to the structs. A slice of structs would be different, as they are in a fixed order. So we would be able to hold a pointer to an actual struct instance that is part of a slice of structs
- Most structs are going to be given type names, but you can create anonymous structs like in `secondPart`
- We can assign one to the other even through they are anonymous types because they are structurally the same, this does not work for named types though
- Two `struct` types are compatible if
  - The fields have the same types and names
  - The field names are in the same order
  - They have the same tags (see `thirdPart` for tags)
- A struct may be copied or passed as a parameter in its entirety
- A struct is comparable if all its fields are comparable
- A zero value for a struct is the "zero value" for each field
- Structs are passed by value unless a pointer is used
  - If you don't pass by pointer you will not be able to modify the struct in a function
  - The way that the pointer syntax works is that you don't need to use the star notation when referencing fields, see the `employee.Boss` on line 56 in the lesson code file for an example
- The `Buffer` struct is written in a way that "makes the zero value useful" like in the video we watched earlier. It has sensible initial values that are immediately ready to be used
- You can create a struct with no fields, it a singleton that Go always have, an empty struct points to this reference
- Struct tags and JSON
  - Encoding and decoding JSON is shown in `fourthPart` which includes hiding missing JSON fields in some instance, this is a nice real world example
  - You can also change the field names by specifying them in the tag
- You can also use struct tags to convert struct field names to database (), it's a good example of another use of tag but it's unsafe because of injection attacks, use a library :-) 

Lesson 14: Reference & Value Semantics
- Pointers vs Values
  - Pointers are shared, not copied
  - Values are copied, not shared
  - Value semantics lead to higher integrity, particularly with concurrent (don't share, especially with concurrency)
  - Pointer semantics may be more efficient
- Advantages of pointers
  - Some objects can't be copied safely (mutex for instance)
  - Some objects are too large to be copied efficiently (consider pointers when size > 64 bytes)
  - Some methods need to mutate the receiver 
  - When decoding protocol data into an object (JSON, etc; often in a variable argument list) see the example in lesson 12 unmarshalling JSON
  - When using a pointer to signal a null object
- A struct with a mutex MUST be passed by reference
- It's ok to have a function that takes in a value parameter and return an updated value, you can just assign it back to the variable that was passed in by value, think functional programming. Mutating inputs in a function always makes me worry 
- Stack allocation is more efficient
  - Accessing a variable directly is more efficient than following a pointer
  - Accessing a dense sequence of data is more efficient than sparse data (an array is faster than a linked list, etc)
- Heap allocation
  - Go would prefer to allocation on the stack, but sometimes can't, some examples:
    - A function returns a pointer to a local object
    - A local object is captured in a function closure
    - A pointer to a local object is sent via a channel
    - Any object is assigned into an interface
    - Any object whose size is variable at runtime (e.g. slices)
  - Go uses "escape analysis" to determine what needs heap allocation
    - You can run with `-gcflags -m=2` to see the analysis
- The `new` operator in Go does not work like it does in JVM based languages. We will be digging into it further later. The part we care about here is that there is not an easy way to tell if something using `new` will end up on the heap or stack.
- The value returned by `range` is always a copy, use the index if you want to mutate the element
  ```
  for i := range things {
    things[i].which = whatever
  }
  ``` 
- Any time a function mutates a slice we must return the updated value from the function
- It is risky to keep a pointer to a slice element since they will re-allocate if over capacity

Lesson 15: HTTP Networking
- The Go standard library has many packages for making web servers:
  That includes:
  - client & server sockets
  - route multiplexing
  - TTP and HTML, including HTML templates
  - JSON and other data formats
  - cryptographic security
  - SQL database access
  - compression utilities
  - image generation 
  - There are also lots of 3rd-party packages with improvements
- An HTTP handler function is an instance of an interface, you can see in the code example file how we implemented this interface
- This is also the lesson that introduces us to methods, see the `ServeHTTP` method on the `HandlerFunc` type

Lesson 16: Homework #3
- Essentially we are doing an exercise where we fetch data using JSON from XKCD
  - Exercise 4.12 from GOPL: fetching from the web
- I did the exercise as he suggested without looking ahead, so it is written in a different way than Matt did, but it still works and uses the JSON marshaling, file IO, networking etc from previous exercises

Lesson 17: Go does OOP
- Go is an Object Oriented language
- Inheritance has conflicting meanings
  - In theory, inheritance should always imply subtyping, the subclass should be a "kind of" the superclass
  - See the Liskov substitution principle
  - As usual "Composition over inheritance" is a good rule for OOP, also remember the "Inheritance tax" that we've seen a bunch of times
- Classes in Go
  - Not having classes can be liberating
  - Go allows defining methods on any user-defined type, rather than only a "class"
  - Go allows any object to implement the method(s) of an interface, not just a "subclass"
- We don't call it "class oriented programming" or "interface oriented programming", just because Go doesn't provide those doesn't mean it is not object oriented

Lesson 18: Methods and Interfaces
- An interface specifies abstract behavior in terms of methods
- "Concrete types" offer methods that satisfy the interface
- A method is a special type of function, it has an invisible parameter which is called "self" or "this" that references the entity itself 
- Again, we can put methods on any user-declared type, and only on user defined types
- Any object that provides the correct methods to satisfy the interface is considered a concrete implementation of an interface that defines those methods
- A method may take a pointer or value receiver, but not both
- Example he showed was how `io.Copy` takes in two interfaces as its parameters, and we called it with concrete types
- All methods must be present to satisfy an interface, this is why is pays to keep interfaces as small as possible
- We can do composition with interfaces, see the `io.ReadWriter` interface for example, small interfaces with composition where needed are more flexible
- All methods for a given type must be defined in the same package where the type is declared, this allows a package importing the type to know all of the methods at compile time
  - We can extend the type in a new package through embedding though (a struct that embeds another struct)

Lesson 19: Composition
- The fields of an embedded struct are promoted to the level of the embedding structure
- You can embed a type into another type even if it is not a struct, something like []int etc
- Even when you embed a pointer to a type inside of the struct its methods will be promoted
- Instead of thinking that a struct that is composed of two or more types "inherits" from those types, think instead, "the methods of those types are promoted into the composed struct"
- Important: Nothing in Go prevents calling a method with a nil receiver (see Sum() in the example code). This would be the equivalent of null.toString() in Java

Further study:
- Hash tables, confirm how they work
- Rob Pike
  - Go Proverbs
- Matt Holiday
- Francesc Campoy (understanding nil)
  - This was a pretty good video
- Compare a lambda expression to a closure
- Read the Go language specification

Weird things about Go:
- You cannot insert into a nil map, but you can read from it (always returns the default value of the map's datatype)

Common helpful patterns
- Checking to see if a key in a map is just a default or not
  - if k, ok = j["and"]; ok {
      // we know that k is not the default value (zero)
    }