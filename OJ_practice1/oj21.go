package main

import (
    "fmt"
    "io"
    "math/big"
)

/*

 */
func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}
func lcm(a, b *big.Int) *big.Int {
    tmp1 := big.NewInt(1)
    tmp2 := big.NewInt(1)
    tmp3 := big.NewInt(1)
    big1 := tmp1.Mul(a, b)
    big2 := tmp2.GCD(nil, nil, a, b)
    big3 := tmp3.Div(big1, big2)
    return big3
}
func main() {
    for {
        var n int64
        _, err := fmt.Scan(&n)
        if err == io.EOF {
            return
        }
        big1 := big.NewInt(n)

        fmt.Println(getLcm(big1))
    }
    //big1 := big.NewInt(10)
    //big2 := big.NewInt(100)
    //fmt.Println(getLcm(big1))
    //fmt.Println(lcm(big1, big2))

}

//大数不能用int
func getLcm(n *big.Int) *big.Int {
    tmp := big.NewInt(1)
    if n.Int64() == 1 {
        return tmp
    }
    tmp3 := big.NewInt(1)
    tmp3 = tmp3.Set(n)
    tmp2 := tmp3.Sub(n, tmp)

    return lcm(tmp3, getLcm(tmp2))
}
