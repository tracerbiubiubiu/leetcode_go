package main

func countBits(n int) []int {
    res := make([]int, n+1)
    for i := 0; i <= n; i++ {
        res[i] = count(i)
    }
    return res
}
func count(n int) int {
    num := 0
    for n > 0 {
        n &= n - 1
        num++
    }
    return num
}
