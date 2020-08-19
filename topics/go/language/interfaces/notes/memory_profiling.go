// Reference: https://www.ardanlabs.com/blog/2017/06/language-mechanics-on-memory-profiling.html

// -memprofile
// go test -run none -bench AlgorithmOne -benchtime 3s -benchmen -memprofile mem.out

// pprof tool to study profile data
// go tool pprof -alloc_space memcpu.test mem.out
//  When profiling memory and looking for "low hanging fruit"
//  use -alloc_space instead of default -inuse_space
//      Shows where every allocation is happening
//      regardless if it's still in memory or not at the time
//      you take the profile
// (pprof) list algOne <-- list command takes regex as argument to find function(s) you want to look at

// ====================================
// Validating use of interfaces in code
// Use interface when:
//  - users of API need to provide implementation detail
//  - APIs have multiple implementations they need to maintain internally
//  - parts of the API that can change have been ID'd and require decoupling
// Do NOT use interface:
//  - for sake of using one
//  - to generalize algorithm
//  - when users can declare their OWN interfaces

// ======================================
// SUMMARY
// Correctness as 1st priority
//      Integrity
//      Readability
//      Simplicity
