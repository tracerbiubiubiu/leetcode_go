package main

import "fmt"

/*
https://leetcode.cn/problems/number-of-islands/
*/
func numIslands(grid [][]byte) int {
    res := 0
    for r := 0; r < len(grid); r++ {
        for c := 0; c < len(grid[0]); c++ {
            if grid[r][c] == '1' {
                res++
                dfs(r, c, grid)
            }
        }
    }
    return res
}

//dfs与bfs
/*
岛屿通用思路
https://leetcode.cn/problems/number-of-islands/solutions/211211/dao-yu-lei-wen-ti-de-tong-yong-jie-fa-dfs-bian-li-/
*/
func dfs(r, c int, grid [][]byte) {
    //判断坐标是否依旧在网格中
    if !isValid(r, c, grid) {
        return
    }
    //这个格子不是陆地，直接返回
    if grid[r][c] != '1' {
        return
    }
    //还需要标记已经访问过的陆地
    grid[r][c] = '2'
    dfs(r-1, c, grid)
    dfs(r+1, c, grid)
    dfs(r, c-1, grid)
    dfs(r, c+1, grid)
}
func isValid(r, c int, grid [][]byte) bool {
    m, n := len(grid), len(grid[0])
    if r >= 0 && r < m && c >= 0 && c < n {
        return true
    }
    return false
}
func main() {
    grid := [][]byte{
        {'1', '1', '0', '0', '0'},
        {'1', '1', '0', '0', '0'},
        {'0', '0', '1', '0', '0'},
        {'0', '0', '0', '1', '1'},
    }
    fmt.Println(numIslands(grid))
}

//bfs
