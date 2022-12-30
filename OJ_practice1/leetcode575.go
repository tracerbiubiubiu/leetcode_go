package main

func distributeCandies(candyType []int) int {
    length := len(candyType)
    if length == 0 {
        return 0
    }
    tMap := make(map[int]bool, 0)
    count := 0
    for _, v := range candyType {
        if _, ok := tMap[v]; ok {
            count++
        }
        tMap[v] = true
    }
    return min(count, length/2)
}

//func min(a, b int) int {
//    if a > b {
//        return b
//    }
//    return a
//}
