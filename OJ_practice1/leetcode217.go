package main

func containsDuplicate(nums []int) bool {
    tmp := make(map[int]bool)
    for i := range nums {
        if _, ok := tmp[nums[i]]; !ok {
            tmp[nums[i]] = true
            continue
        }
        return true
    }
    return false
}
