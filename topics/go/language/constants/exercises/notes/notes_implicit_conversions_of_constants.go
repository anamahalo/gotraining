package main

import (
    "fmt"
    "time"
)

const fiveSeconds = 5 * time.Second

func main() {
    now := time.Now()
    lessFiveNanoSeconds := now.Add(-5)
    lessFiveSeconds := now.Add(-fiveSeconds)

    fmt.Printf("Now     : %v\n", now)
    fmt.Printf("Nano    : %v\n", lessFiveNanoSeconds)
    fmt.Printf("Seconds : %v\n", lessFiveSeconds)
}

// Output
// Now     : 2020-08-17 20:58:37.698831 -0500 CDT m=+0.000188929
// Nano    : 2020-08-17 20:58:37.698830995 -0500 CDT m=+0.000188924
// Seconds : 2020-08-17 20:58:32.698831 -0500 CDT m=-4.999811071

// var difference int = -5
// var less FiveNano = now.Add(difference)
// Returns compiler error:
// ./const.go:#L: cannot use diff (type int) as type time.Duration in function argument
