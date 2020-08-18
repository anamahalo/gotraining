// As we continue to add or rm k/v pairs from the map
// the efficiency of the map lookups begin to deteriorate.
// Load threshold values that determine
// when to grow the hash table are based on the
// following 4 factors:

// % overflow  : Percentage of buckets which have an overflow bucket
// bytes/entry : # of overhead bytes used per k/v pair
// hitprobe    : # of entries needing to be checked upon key lookup
// missprobe   : # of entries needing to be checked upon absent key lookup

// Growing the hash table starts with assigning a pointer
// called the "old bucket" pointer to the current bucket array; then
// new bucket array is allocated to hold 2n of existing buckets.
// Memory is not initialized so allocation is fast, though
// allocations could be large.

// Once memory for new bucket array is avail, k/v pairs from old
// bucket array can be moved or "evacuated" to new bucket array.
// Evacs happens as k/v pairs are added/removed from map.
// k/v pairs together in old bucket could be moved to diff buckets
// inside new bucket array.
// Evac algorithm attempts to dist k/v pairs evenly across new
// bucket array.

// Iterators need to run thru old buckets until every old bucket
// has been evac'd: this affects how k/v pairs are returned
// during iteration operations.

// Under the hood (recall written in C), a lot of memory and 
// pointer manipulation is had to keep things fast, efficient, and safe

// THUS:
// If n keys are known ahead of time it's best to allocate
// that space *during* initialization.
// (Also why maps are unsorted collections and why iterators
// *seem* random when mapping)

// Instead of N runtime.hashmap implementations in the final binary,
// only N runtime.maptype *values* - program space saving and compile
// time.
// (For contrast, C++ has N runtime.hashmap implementations in final)
