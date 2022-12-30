package main

import "fmt"
/*
冒泡排序
O(n^2)
*/

//func BubbleSort(origin []int) {
//    length := len(origin)
//    didSwap := false
//    for i := length - 1; i > 0; i-- {
//        // 每次从第一位开始比较，比较到第 i 位就不比较了，因为前一轮该位已经有序了
//        for j := 0; j < i; j++ {
//            if origin[j] < origin[j+1] {
//                origin[j], origin[j+1] = origin[j+1], origin[j]
//                didSwap = true
//            }
//        }
//        if !didSwap {
//            return
//        }
//    }
//}

func BubbleSort(origin []int){
    length := len(origin)
    //didswap := false
    for i := 1;i < length;i ++ {
        // 注意这里的 = ，在开头下标=0 时，还要判断一次
        for j := i - 1;j >= 0; j -- {
            if origin[j] > origin[j+1]{
                origin[j], origin[j+1] = origin[j+1], origin[j]
                //didswap = true
            }
        }
        //if !didswap{
        //    return
        //}
    }
}

func main() {
    list := []int{5, 9, 1, 6, 8, 14, 66, 49, 25, 4, 6, 83}
    BubbleSort(list)
    fmt.Println(list)
}