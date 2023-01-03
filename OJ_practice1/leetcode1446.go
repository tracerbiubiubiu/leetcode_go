package main

import "fmt"

/*
https://leetcode.cn/problems/consecutive-characters/
*/
//我的思路是便利一次，相同++，不同重置，每次更新res
func maxPower(s string) int {
    res := 0
    count := 1
    for i := 0; i < len(s)-1; i++ {

        if s[i] == s[i+1] {
            count++
        } else {
            res = maxOf2(res, count)
            count = 1
        }
    }
    return maxOf2(res, count)
}
func maxOf2(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func main() {
    a := "ss"
    fmt.Println(maxPower(a))
}
