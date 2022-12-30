package main

import (
    "fmt"
    "math"
)

/*

 */

func reverse(x int) int {
    //位数
    res := 0
    for x != 0 {
        if res > math.MaxInt32/10 || res < math.MinInt32/10 {
            return 0
        }
        carry := x % 10

        x = x / 10

        res = res*10 + carry
    }
    return res
}
func main() {
    a := 1534236469
    fmt.Println(a > math.MaxInt32)
    fmt.Println(reverse(a))
}
