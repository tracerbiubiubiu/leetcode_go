package main

import "fmt"
/*
希尔排序
O(n^2)
不稳定
直接插入排序的改进版本
越有序的数列，直接插入排序的效率越高
*/

// 每次增量折半
func ShellSort(list []int) {
    n := len(list)
    // 注意 =
    // step = 1 有点像冒泡？
    for step := n / 2; step >= 1; step /= 2 {
        // 开始插入排序 步长为step
        for i := step; i < n; i += step {
            for j := i - step; j >= 0; j -= step {
                // 满足插入那么交换元素
                if list[j+step] < list[j] {
                    list[j], list[j+step] = list[j+step], list[j]
                    continue
                }
                break
            }
        }
    }
}

func main() {
    list := []int{5}
    ShellSort(list)
    fmt.Println(list)

    list1 := []int{5, 9}
    ShellSort(list1)
    fmt.Println(list1)

    list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
    ShellSort(list2)
    fmt.Println(list2)

    list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 2, 4, 23, 467, 85, 23, 567, 335, 677, 33, 56, 2, 5, 33, 6, 8, 3}
    ShellSort(list3)
    fmt.Println(list3)
}