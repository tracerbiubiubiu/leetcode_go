package main

/*
https://leetcode.cn/problems/spiral-matrix-ii/
*/
func generateMatrix(n int) [][]int {
    res := make([][]int, n)
    for i := 0; i < n; i++ {
        res[i] = make([]int, n)
    }
    top, down, left, right := 0, n-1, 0, n-1
    for count := 1; count <= n*n; {
        for i, j := top, left; j <= right; j++ {
            res[i][j] = count
            count++
        }
        top++
        for i, j := top, right; i <= down; i++ {
            res[i][j] = count
            count++
        }
        right--
        for i, j := down, right; top <= down && j >= left; j-- {
            res[i][j] = count
            count++
        }
        down--
        for i, j := down, left; left <= right && i >= top; i-- {
            res[i][j] = count
            count++
        }
        left++
    }
    return res
}
