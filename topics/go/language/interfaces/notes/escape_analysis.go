// What happens when you share a value UP the stack?

// Escape Analysis
//  Process that compiler uses to determine
//  placement of values created by your program.
//  => compiler performs static code analysis to determine
//     if value can be placed on the stack frame for the function
//     constructing it, or if the value must "escape" to the HEAP
//          ONLY through HOW CODE IS WRITTEN is this decision
//          dictated, i.e.: no keyword or function to direct
//          compiler

// Heaps
// 2nd area of memory in addition to stack for storing values
// Not self-cleaning like stacks
//  .: cost is greater
//      GC needs to get involved (25% avail CPU capacity)
//      potential for "stop the world" latency
//          benefit is no heap memory mgmt
// Values on heap constitute memory allocations in Go
//      Allocations put pressure on GC b/c EVERY value
//      on heap no longer referenced by a pointer
//      needs to be rm'ed
//      more values checked and rm'ed => more GC work per run
//      .: pacing algorithm is constantly working to balance
//         size of heap w/pace it runs at

// No goroutine is allowed to have a pointer that points
// to memory on another goroutine's stack.
//  b/c stack memory for goroutine can be replaced w/new
//  block of memory when stack has to grow or shrink
//      if runtime had to track pointers to other goroutine
//      stacks, it would be too much to manage and the
//      "stop the world" latency in updating pointers on those
//      stacks would be overwhelming

//======================
// How stacks grow/change
package main

// Number of elements to grow each stack frame
// Run w/10 and then w/1024
const size = 1024

// main is entry point for app
func main() {
    s := "HELLO"
    stackCopy(&s, 0, [size]int{})
}

// stackCopy recursively runs increasing the size of the stack
func stackCopy(s *string, c int, a [size]int) {
    println(c, s, *s)

    c++
    if c == 10 {
        return
    }

    stackCopy(s, c, a)
}

// NOTE STRING VALUE INSIDE STACK CHANGES TWICE
// Cont. to escape_analysis2.go
