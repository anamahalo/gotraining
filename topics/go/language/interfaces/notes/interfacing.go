// Interfaces are valueless

// type reader interface {
//    read(b []byte) (int,error)
// }

// var r reader <-- r is valueless
// .: Interfaces define and create valueless values
//      - nothing real about var r
//      - nothing concrete about var r
//      - var r is valueless

// ====================
// polymorphic function
//  behaves differently depending on data it operates on

// func retrieve(r reader) error {
//      data := make([]byte, 100)
//
//      len, err := r.read(data)
//      if err != nil {
//          return err
//      }
//
//      fmt.Println(string(data[:len]))
//      return nil
// }

// .: concrete data is driver allowing for abstraction
//    of diff behavior (polymorphism) to be designed
//    and implemented

// .: above function is saying (knowing reader is valueless)
//      "I will accept any piece of concrete data
//      (any value or pointer) that implements the
//      reader contract.  That implements the full
//      method set of behavior defined by the reader
//      interface."
//          - func retrieve() is not bound to a single
//            piece of conrete dat, but bound to ANY
//            concrete data that exhibits the read BEHAVIOR
//
// .: methods!
//      give data behavior
//      once piece of data has behavior, polymorphism is realized
//  When to choose method over function?
//      when piece of data NEEDS behavior to satisfy
//      method set of a given INTERFACE
// See interfacing2.go
