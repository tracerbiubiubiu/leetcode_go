package main

import "fmt"

func main() {
    metrix := make([][]int, 101)
    for k, _ := range metrix {
        metrix[k] = make([]int, 101)
    }
    var n, m, t int
    fmt.Scan(&n, &m)
    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            fmt.Scan(&t)
            metrix[i][j] = t + max(metrix[i][j-1], metrix[i-1][j])
        }
    }
    fmt.Println(metrix[n][m])
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
