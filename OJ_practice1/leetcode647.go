package main

/*

 */
func countSubstrings(s string) int {
    if len(s) < 2 {
        return len(s)
    }
    res := 0
    for i := 0; i <= len(s)-1; i++ {
        res += doExpand(s, i, i)
        res += doExpand(s, i, i+1)
    }
    return res
}
func doExpand(s string, left, right int) int {
    count := 0
    for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
        left--
        right++
        count++
    }
    return count
}