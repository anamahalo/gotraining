// Frame boundaries

// Stacks
// It's during each function call, when the frame is taken,
// that the stack memory for that frame is wiped clean.
//  Done by initialization of any values placed in the frame.
//  Because all values are initialized to at least their
//  "zero value", stacks clean themselves properly on
//  every function call.

// =====================
// Pointers
// Serve one purpose: to share value w/a function so function
// can read and write to that value even though the value
// does not exist directly inside its own frame.
//  If "share" doesn't come out of your mouth
//  a pointer is not needed.
//  Think about clear vocab and not operators or syntax.
//  TL;DR pointers are for SHARING and replace the & operator
//  for the word "sharing" as you read code.

// Every declared type has a pointer
// All have same memory size and representation
//  4 or 8 bytes that rep an address
//  32bit architectures (playground) - 4 bytes of memory
//  64bit architecutres (machine)    - 8 bytes of memory

// Pointer types are considered type literals
//  Unnamed types composed from existing type

package main

func main() {
    // Declare var of type int w/value of 10
    count := 10

    // Display "value of" and "address of" count
    println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")

    // Pass the "address of" count.
    // "I am sharing the count var w/the increment func"
    increment(&count)

    println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
}

//go:noinline
// "Pass by value" - address instead of int
// Since value of address is being copied and pasted,
// need var inside frame of increment to receive and store
// this integer based address
func increment(inc *int) {
    // Increment "value of" count that the "pointer points to"
    // (dereferencing)

    // * = value pointer points to
    // pointer var allows INDIRECT memory access OUTSIDE
    // of function's frame that is using it
    //  i.e.: dereferencing the pointer
    // increment() must still have pointer variable
    // within its frame it can directly read to perform
    // the indirect access
    *inc++
    println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Point To[", *inc, "]")
}

// Pointer vars are like any other var in that they've
// memory allocation and hold a value.
// Difference is that pointer vars, REGARDLESS of the type
// of value they can point to, are ALWAYS the same size and rep
// Remember to distinguish the TYPE declaration from the POINTER operation


//===============
// THUS

// Functions execute w/in scope of frame boundaries
// that provide individual memory space for each respective
// function

// When function is called, transition takes palce b/w 2 frames

// Benefit of passing data "by value" is legibility

// Stack is important b/c it provides physical memory space
// for the frame boundaries that are given to each individual
// function

// All stack memory BELOW the active frame is INVALID but memory
// from the active frame and above is valid

// Making a function call means the goroutine needs to frame
// a new section of memory on the stack

// It's during each function call, when the frame is taken,
// that the stack memory is wiped clean.

// Pointers serve one purpose: SHARING :)
// share value w/function so function can read/write to that value
// EVEN THOUGH the value doesn't exist directly inside its own frame

// Every type gets a free compliment pointer type for sharing use

// Pointer var allows INDIRECT memory access OUTSIDE the function's frame that is using it

// Pointer vars are like other vars except they have memory allocation and hold a value.
