```go
func main() {
    var i int
    fmt.Println(i) // Output: 0
    i = 5
    defer fmt.Println(i) // Output: 5
    i++
    fmt.Println(i) //Output: 6
}
```