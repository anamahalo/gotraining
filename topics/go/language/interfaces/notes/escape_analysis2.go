// Anytime value is shared OUTSIDE scope of a function's
// stack frame, it will be allocated on the heap
// Job of escape analysis algorithms to find these situations
// and maintain level of integrity in the program:
//      making sure that access to any value is always
//      accurate, consistent and efficient

// go:noinline <--- prevents compiler from inlining code
//                  for functions directly in main
//                  Inlining would erase function calls

// Reference: https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-escape-analysis.html
// for visuals

// Compiler Reporting
// go build -gcflags "-m -m" (2 levels of m here; up to 4 avail)
// see example above - compiler reporting will let you know
// whether something escapes to heap


// SUMMARY
// only how value is shared will determine what compiler will do
// with that value
// ANY time you share value up the call stack it's going to escape

// these are guidelines for choosing value or pointer semantics
// for any given type
// cost-benefit analysis:
//  value semantics
//      keep values on stack which reduces pressure on GC
//      => diff copies of any given value  must be stored,
//      tracked, maintained
//  pointer semantics
//      place values on heap which puts pressure on GC
//      => efficient b/c there's only one value that needs
//      to be stored, tracked, and maintained

// use each semantic correctly, consistently and in balance
