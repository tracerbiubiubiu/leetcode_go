package main

import (
    "fmt"
    "math/big"
)

//func main() {
//    var n int64
//    fmt.Scan(&n)
//    if n == 0 {
//        fmt.Println(1)
//        return
//    }
//    ans := calN(n)
//    fmt.Println(ans)
//}
func calN(n int64) int64 {
    if n == 1 {
        return 1
    }
    tmp := big.NewInt(int64(n))
    next := big.NewInt(int64(calN(n - 1)))
    a := tmp.Mul(tmp, next)
    if a.Int64() < 0 {
        fmt.Println(a)
    }

    return a.Int64()
}

//
func main() {
    var n string
    fmt.Scan(&n)
    big1 := new(big.Int)
    //big.SetInt64(n)
    big1, ok := big1.SetString(n, 10)
    if !ok {
        return
    }
    big2 := big.NewInt(1)
    result := big2.MulRange(1, big1.Int64())
    fmt.Print(result.String())
}
