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
  - If you go over the capacity (cap) of a slice it will need to be reallocated
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
  - len, cap and range can all be called on nil values fro 
  - The idea that we can make nil safe and useful is really helpful

Further research:
- Hash tables, confirm how they work
- Rob Pike
- Matt Holiday
- Compare a lambda expression to a closure

Weird things about Go:
- You cannot insert into a nil map, but you can read from it (always returns the default value of the map's datatype)

Common helpful patterns
- Checking to see if a key in a map is just a default or not
  - if k, ok = j["and"]; ok {
      // we know that k is not the default value (zero)
    }