package main

import (
    "fmt"
    "math/big"
)

/*

 */
func main() {
    var num int
    fmt.Scan(&num)
    var a, b string
    i := 1
    for ; i <= num; i++ {
        fmt.Scan(&a, &b)
        fmt.Printf("Case %d:\n", i)
        ans := aAddBPlus(a, b)
        fmt.Printf("%s + %s = %s", a, b, ans)
        if i != num {
            fmt.Printf("\n\n")
        }
    }
}
func aAddBPlus(a, b string) string {
    i, _ := new(big.Int).SetString(a, 10)
    j, _ := new(big.Int).SetString(b, 10)
    k := i.Add(i, j)
    return fmt.Sprint(k)
}
