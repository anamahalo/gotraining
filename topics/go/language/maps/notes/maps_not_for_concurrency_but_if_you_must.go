// Maps aren't safe for concurrent use
// b/c it's not defined what happens when you read and write
// to them simultaneously
// But if you must, accesses must be mediated by some kind of
// synchronization mechanism.
// Common way to protect maps is with sync.RWMutex

// declares counter variable that is an anonymous struct
// containing a map and an embedded sync.RWMutex
var counter = struct {
    sync.RWMutex
    m map[string]int
}{m: make(map[string]int)}

// to read from the counter, take the read lock:
counter.RLock()
n := counter.m["some_key"]
counter.RUnlock()
fmt.Println("some_key:", n)

// to write to the counter, take the write lock:
counter.Lock()
counter.m["some_key"]++
counter.Unlock()

// note: notice the locks and unlocks
