package main

import (
    "fmt"
)

/*
https://leetcode.cn/problems/repeated-dna-sequences/
*/
func findRepeatedDnaSequences(s string) []string {
    if len(s) < 10 {
        return nil
    }
    res := []string{}
    mp := make(map[string]int)
    //i := 0
    for i := 0; i <= len(s)-10; i++ {
        j := i + 10
        tmpS := s[i:j]
        if _, ok := mp[tmpS]; ok {
            mp[tmpS]++
        } else {
            mp[tmpS] = 1
        }
    }
    //也可以再循环中只统计出现次数为2的
    for i, v := range mp {
        if v > 1 {
            res = append(res, i)
        }
    }

    return res
}
func main() {
    s := "AAAAAAAAAAA"
    //s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
    fmt.Println(findRepeatedDnaSequences(s))
}
