// Basic Guidelines

//  At time you declare a type you must decide what
//  semantic is being used

//  Functions and methods must respect semantic choice for
//  given type

//  Avoid having method receivers that use diff semantics than
//  those corresponding to the given type

//  Avoid changing semantic for given type

//  Exception: Unmarshaling always requires use of pointer
//             semantics

// ======================
// Built-In Types
//  Numeric, textual, boolean data
//  Value semantics over pointers w/exception of a very good reason

// References Types
//  Slice, map, interface, function, channel types
//  Value semantics b/c designed to stay on stack and minimize heap pressure

// User Defined Types
//  Discernment needed

// Reference: https://www.ardanlabs.com/blog/2017/06/design-philosophy-on-data-and-semantics.html
