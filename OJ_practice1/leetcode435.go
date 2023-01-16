package main

import (
    "fmt"
    "sort"
)

/*
https://leetcode.cn/problems/non-overlapping-intervals/description/
*/
func eraseOverlapIntervals(intervals [][]int) int {
    res := 0
    if len(intervals) == 1 {
        return res
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })

    for i, cur := 0, 1; cur < len(intervals); {
        for cur < len(intervals) && intervals[i][1] <= intervals[cur][0] {
            i = cur
            cur++
        }
        for cur < len(intervals) && intervals[i][1] > intervals[cur][0] {

            cur++
            res++
        }
    }
    return res
}
func main() {
    intervals := [][]int{
        {0, 2},
        {1, 3},
        {2, 4},
        {3, 5},
        {4, 6},
    }
    fmt.Println(eraseOverlapIntervals(intervals))
}
