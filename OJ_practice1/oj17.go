package main

import (
    "fmt"
    "math/big"
)

func main() {
    for {
        var n, m string
        if _, err := fmt.Scan(&n, &m); err != nil {
            break
        }
        a, _ := new(big.Int).SetString(n, 10)
        b, _ := new(big.Int).SetString(m, 10)
        fmt.Println(a.Mul(a, b))
    }
}

//试试大数包
