// Before starting any project, ask yourself:
//      "What is the purpose of the logs for this application?"
//
// Logging over dependency on debuggers
//
// "I believe logging, when reasonable and practical, should be isolated to the
// single purpose of debugging errors. This is the same whether the application is
// running in production or development. Logging can’t always be isolated to the 
// debugging of errors, but when it can, it will simplify things tremendously. If 
// you move concerns like metrics and informational data out of the logs, the 
// guidelines for how, what and when to log become clear to every developer and 
// can be applied in a consistent manner. This also extends to code reviews and 
// specialized linters."

// Above quote also includes warnings as exclusions in logs, i.e.:
// logs should be errors ONLY.
// So where should the logs for metrics and informational data go if so?
//      "Use dashboarding to publish points of interest for an administrator
//       or developer."
//      "... you should question the usefulness of any information
//      you publish to validate you are establishing a clear relationship
//      between the problem and its cause." - Gabriel Janson
//
// Handling Errors means:
//  Error has been logged
//  Application is back to 100% integrity
//  Current error is not reported any longer
//
//
// NOTE:
// Pkgs that are reusable across many projects only return ROOT error values

// If the error is NOT going to be handled, wrap and return up the call stack
//  Good way to determine whether you are logging enough context or too much
//  is by inspecting the log and validating they're working for you during
//  development.
// 
// Once an error is handled, it is not allowed to be passed up the call stack
// any longer.
//  i.e.: once it's decided the function/method will handle the error,
//        the error ceases to be an error
//        if the function/method is still required to issue a return,
//        it can't return an error value; it should only return nil
