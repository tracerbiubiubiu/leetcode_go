package main

import "math"

func abs(a, b int) int {
    return int(math.Abs(float64(a - b)))
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
    for i := 0; i < len(nums)-1-indexDiff; i++ {
        j := i + indexDiff
        for j > i && j <= len(nums)-1 {
            if abs(nums[i], nums[j]) <= valueDiff {
                return true
            }
            j += indexDiff
        }

    }
    return false
}