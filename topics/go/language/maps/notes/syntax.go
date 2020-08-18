// Associative containers mapping keys to values.

// Construct:  m := map[key]value{}
// Insert:     m[k] = v
// Lookup:     v = m[k]
// Delete:     delete(m, k)
// Iterate:    for k, v := range m
// Size:       len(m)

// Key must have an == operation (no maps, slices, or funcs as keys)
// Operations run in constant expected time.
