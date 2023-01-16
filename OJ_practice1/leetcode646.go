package main

import "sort"

/*
https://leetcode.cn/problems/maximum-length-of-pair-chain/
*/
func findLongestChain(pairs [][]int) int {
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i][1] < pairs[j][1]
    })
    res := 0
    cur := pairs[1][0]
    for i := range pairs {
        if pairs[i][0] > cur {
            cur = pairs[i][1]
            res++
        }
    }
    return res
}
