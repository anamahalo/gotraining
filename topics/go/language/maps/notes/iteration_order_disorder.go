// When iterating over a map w/a range loop the iteration order
// is not specified; and
// is not guaranteed to be the same from iteration to the next.
// Require a stable iteration order?; then
// a separate data structure that specifies that order
// is required.

// Example below uses a separate sorted slice of keys to print
// a map[int]string in key order:

import "sort"

var m map[int]string
var keys []int
for k := range m {
    keys = append(keys, k)
}
sort.Ints(keys)
for _, k := range keys {
    fmt.Println("Key:", k, "Value:", m[k])
}
