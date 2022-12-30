package main

import "fmt"

/*
插入排序
O(n^2)
稳定
*/

// 默认左侧有序，右侧无序
//func InsertSort(list []int) {
//    n := len(list)
//    // = 号别忘
//    for i := 1; i <= n-1; i++ {
//        // 待排序
//        unhandle := list[i]
//        j := i - 1
//        if list[i] < list[j] {
//            //fmt.Println(list[i])
//            // 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
//            for ; j >= 0 && unhandle < list[j]; j-- {
//                //fmt.Println(list[i])
//                //fmt.Println(unhandle)
//                list[j+1] = list[j]
//            }
//            // 退出循环之前多减了1 这里j+1就是上面的j
//            list[j+1] = unhandle
//        }
//    }
//}
// 另一种 有点像冒泡？
func InsertSort(list []int) {
   n := len(list)
   // = 号别忘
   for i := 1; i <= n-1; i++ {
       // 待排序
       unhandle := list[i]
       j := i - 1
       if unhandle < list[j] {
           // 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
           for ; j >= 0 ; j-- {
               if unhandle < list[j] {
                   list[j+1], list[j] = list[j], list[j+1]
               }
           }
           // 退出循环之前多减了1 这里j+1就是上面的j
           //list[j+1] = unhandle
       }
   }
}
func main() {
    //list := []int{5}
    //InsertSort(list)
    //fmt.Println(list)
    //
    //list1 := []int{5, 9}
    //InsertSort(list1)
    //fmt.Println(list1)

    //list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
    list2 := []int{5, 9, 1, 6, 8}
    InsertSort(list2)
    fmt.Println(list2)
}

