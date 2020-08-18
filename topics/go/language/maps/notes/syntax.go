// Associative containers mapping keys to values.

// Construct:  m := map[key]value{}
// Insert:     m[k] = v
// Lookup:     v = m[k]
// Delete:     delete(m, k)
// Iterate:    for k, v := range m
// Size:       len(m)

// Key must have an == operation (i.e.: no maps, slices, or funcs as keys)
// (Value type can be anything)
// Operations run in constant expected time.
