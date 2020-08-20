// Benefit to exception-based languages is that, compared to Go,
// even an UNHANDLED exception will still be raised via a stack trace
// at runtime if it occurs.
// In Go, it's possible to not handle a critical error AT ALL.
// .: Go offers you full control of error handling,
//      but also FULL RESPONSIBILITY.
//
//
// ========================================
// Why Go doesn't have Exception Handling
// ========================================
// Included in Zen of Go:
// 1. Simplicity matters
// 2. Plan for failure, NOT success
//
// `if err != nil` returning `(value, error)` makes you think of failures
// FIRST AND FOREMOST
// no `try catch` blocks
//
// Since exception-based code appropriately handles exceptions it encourages
// programmers to never check errors, knowing that at the very least,
// some exception will be handled automatically at runtime.
// 
// "violating referential transparency" - when raising exceptions becomes opaque
// From "Functional Programming in Scala"
//  In the case of exceptions, the POINT OF EVALUATION for an exception MATTERS.
//
//
//
// =========================================
// Writing idiomatic error handling in Go
// =========================================
// 1. Add stack traces when your errors are actionable to developers
// 2. Do something w/your returned errors: don't just bubble them up to main,
//    log them, and then forget them
// 3. Keep your error chains unambiguous
//
// EXAMPLE
// // In controllers/user.go
// if err := db.CreateUser(user); err != nil {
//      return fmt.Errorf("could not create user: %w", err)
// }
//
// // In database/user.go
// func (db *Database) CreateUser(user *User) error {
//      ok, err := db.DoesUserExist(user)
//      if err != nil {
//          return fmt.Errorf("could not check if user already exists in db: %w", err)
//      }
//      ...
// }
//
// func (db *Database) DoesUserExist(user *User) error {
//      if err := db.Connected(); err != nil {
//          return fmt.Errorf("could not establish db connection: %w", err)
//      }
//      ...
// }
//
// func (db *Database) Connected() error {
//      if !hasInternetConnection() {
//          return errors.New("no internet connection")
//      }
//      ...
// }
//
//
// Note above how ELEGANT and completely NAMESPACED each error is by their
// respective functions.  The fmt.Errorf() utilized in the way above
// is called error chaining.  Also note how DEFENSIVE it is, keeping in line
// with Golang Zen about preparing for FAILURES not successes.

