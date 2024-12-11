```go
package main

import "fmt"

func main() {
    var i int
    fmt.Println(i) // Output: 0
    i = 5
    defer func() {
        fmt.Println(i)
    }()
    i++
    fmt.Println(i) // Output: 6
}
```