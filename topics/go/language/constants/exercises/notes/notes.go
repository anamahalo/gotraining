// Untyped constant on the LHS of declaration is given same kind and value
//   as constant on the RHS:
const untypedInteger       = 12345     // kind: integer
const untypedFloatingPoint = 3.141592  // kind: floating-point

// Declare typed constant, must be same on both sides:
const typedInteger int           = 12345
const typedFloatingPoint float64 = 3.141592

// Numberic constants can be one of four kinds: integer, floating-point, complex and rune:
12345     // kind: integer
3.141592  // kind: floating-point
1E6       // kind: floating-point

// Integer constants must be 256 bits of precision
// Compile time only

// See examples:
// notes_can_be_compiled.go
// notes_cannot_be_compiled.go


// Must also fit into the range for declared type.
// Invalid:
const myUint8 uint8 = 1000
// uint8 can only rep nums from 0 to 255

