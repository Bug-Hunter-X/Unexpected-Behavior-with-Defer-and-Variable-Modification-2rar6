# Go Defer Statement Surprise

This example demonstrates a common pitfall when using Go's `defer` keyword.  The `defer` statement schedules a function call to be executed after the surrounding function completes, but the arguments are evaluated *at the time of the `defer` statement*, not at the time of execution.  This seemingly simple code snippet showcases how modifications to a variable after a `defer` statement can lead to surprising results.

## How to reproduce

1. Save the code in `bug.go`.
2. Run it using `go run bug.go`.

## Solution

The `bugSolution.go` file contains a corrected version that achieves the desired behavior using different approaches.