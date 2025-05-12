package lesson_10

import "fmt"

func Lesson10() {
	someSliceStuff()
	moreSliceStuff()
}

func someSliceStuff() {

	// each descriptor has a length, capacity and pointer
	var s []int            // descriptor points to nil
	t := []int{}           // descriptor points to an empty array
	u := make([]int, 5)    // normal slice
	v := make([]int, 0, 5) // 0 length, 0 capacity (empty)

	// common mistake when creating u...
	// if you append to u, it will append to the sixth index

	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(s), cap(s), s, s == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(t), cap(t), t, t == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(u), cap(u), u, u == nil)
	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(v), cap(v), v, v == nil)

	a := [3]int{1, 2, 3}
	b := a[0:1]

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := b[0:2] // WTF

	// this is crazy, it gives [1 2]
	fmt.Println("c = ", c)

	// the reason why is because when you make a slice that uses the two index slice operator to slice an existing slice, the underlying capacity of the new slices is the same as the original slice
	fmt.Println(len(b))
	fmt.Println(cap(b))

	fmt.Println(len(c))
	fmt.Println(cap(c))

	// this sets the initial offset, length and capacity of the new slice
	d := c[0:1:1]

	fmt.Println("d = ", d)
	fmt.Println(len(d))
	fmt.Println(cap(d))

	// this will panic because d only has capacity 1
	//e := d[0:2]
	//
	//fmt.Println("e = ", e)
	//fmt.Println(len(e))
	//fmt.Println(cap(e))
}

func moreSliceStuff() {
	a := [...]int{1, 2, 3}
	b := a[0:1]
	c := b[0:2] // WTF

	// the pointer is to the same memory location for all three as they refer to the same
	// backing array
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)

	c = append(c, 5)

	// the capacity of c is 3 so there is room to add the value to the backing array,
	// so the pointer to the array does not change
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)

	c = append(c, 3)

	// the capacity of c is 3 so there is no room to add the value to the backing array,
	// so the pointer changes to the newly reacllocated backing array
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)
}
