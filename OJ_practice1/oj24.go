package main

import "fmt"

func main() {
    var a, k, n int
    _, err := fmt.Scan(&a, &k, &n)
    if err != nil {
        return
    }
    fmt.Println(((n * k) + 2*a) * (n + 1) / 2)
}
