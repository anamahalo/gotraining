// Abstraction
//  Duplication is far cheaper than the wrong abstraction.
//  Purpose of abstracting is not to be vague,
//  but to create a new semantic level in which
//  one can be absolutely precise.
//  Key takeaway:
//      Abstractions leak so make them fluid.

// DOD
//  Focus on relationship interfaces have w/concrete data.
//  1. Problem space
//      Concrete data
//          physical state stored in memory, sent over a network,
//          write to a file and manipulate.
//      Mechanical sympathy
//          based on conrete data and how effectively
//          you allow the machine to perform these
//          transformations
//  2. Problem change => algorithm change
//      Designing for change => legibility, performance < cost
// .: Objective
//      Need a way to:
//          - allow algorithms to remain small and precise for
//            each data transform performed
//          - change w/o cascading throughout
// .: Interfaces!
//      Structs as types of concrete data
//      Vars as values of type
//      Interface - opposite of struct type
//          Can only declare method set of behavior
//          .: Nothing concrete about interface type
//  See notes example



// =================
// Value semantics
//  Interface can store its own copy of a value
// Pointer semantics
//  Value can be shared w/the interface by storing a COPY
//  of the value's address
