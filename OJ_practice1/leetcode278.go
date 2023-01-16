package main

/*

 */
func isBadVersion(version int) bool {}

func firstBadVersion(n int) int {
    left, right := 1, n
    for left <= right {
        mid := left + (right-left)/2
        if !isBadVersion(mid) {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return left
}
