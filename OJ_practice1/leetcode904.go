package main

/*
https://leetcode.cn/problems/fruit-into-baskets/
*/
//一个超大数组会超时 2s
//func totalFruit(fruits []int) int {
//
//    res := 0
//    for i := range fruits {
//        start, end := i, i
//        type1 := fruits[start]
//        for end <= len(fruits)-1 && fruits[end] == type1 {
//            end++
//        }
//        if end == len(fruits) {
//            res = maxOf2(res, end-start)
//            continue
//        }
//        type2 := fruits[end]
//        for end <= len(fruits)-1 && (fruits[end] == type2 || fruits[end] == type1) {
//            end++
//        }
//        res = maxOf2(res, end-start)
//    }
//    return res
//}
func maxOf2(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func totalFruit(fruits []int) int {
    cnt := map[int]int{}
    res := 0
    start := 0
    for end := range fruits {
        cnt[fruits[end]]++
        for len(cnt) > 2 {
            y := fruits[start]
            cnt[y]--
            if cnt[y] == 0 {
                delete(cnt, y)
            }
            start++
        }
        res = maxOf2(res, end-start+1)
    }
    return res
}
