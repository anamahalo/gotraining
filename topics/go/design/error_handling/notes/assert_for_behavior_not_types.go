// Caller should not presume anything about state of other values
// returned from that call without first checking the error.

// Returning yes or no on an error versus whether it should retry
// Networking usually calls for investigation into a retry

// Common practice of package authors is to return errors of a known
// public type, so the caller can assert and inspect them.  This leads
// to undesirable outcomes:
//      Public error types increase surface area of the package's API
//      New implementations must only return types specified in the
//      interface's declaration, even if they are a poor fit
//      Error type cannot be changed or deprecated after introduction
//      w/o breaking compatability, making for a brittle API
//
// We want to communicate intention of original author / consumer,
//  without having to overly couple their implementation to the caller.
//
// .: Assert errors for BEHAVIOR, not type
//  Assert that the value implements a particular behavior
//
// Golang                      - "has a" nature
// Inheritance-based languages - "is a [subtype of]" nature
//
//
// Example:
// func isTimeout(err error) bool {
//      type timeout interface {
//              Timeout() bool
//      }
//      te, ok := err.(timeout)
//      return ok && te.Timeout()
// }
//
// Caller can use isTimeout() to determine if error is RELATED to a timeout,
//  via its implementation of `timeout` interface, and then confirm if
//  the error was timeout related -- all WITHOUT knowing anything
//  about the type, or original source of the error value.
//
//
// ================
// GUIDELINES
// ================
// Assert for BEHAVIOR
// Packages - Authors
//  If pkg generates errors of a TEMPORARY nature, ensure you return
//  error types that IMPLEMENT the respective interface methods.
//  If you wrap error values on the way out, ensure your wrappers
//  respect the interface(s) that the underlying error value implemented.
// Packages - Users
//  If you need to inspect an error, use INTERFACES to assert the behavior
//  you expect, NOT the error's type.  Rather than asking authors for
//  public error types, ask they make their types CONFORM to common interfaces
//  by supplying Timeout() or Temporary() methods as appropriate.
//

