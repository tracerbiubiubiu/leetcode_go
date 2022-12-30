package main

/*
选择排序
O(n^2)
*/
import "fmt"

func SelectSort(origin []int) {
    length := len(origin)
    for i := 0; i < length-1; i++ {
        min := origin[i]
        minIndex := i
        for j := i + 1; j < length; j++ {
            if origin[j] < min {
                min = origin[j]
                minIndex = j
            }
        }
        if i != minIndex {
            origin[i], origin[minIndex] = origin[minIndex], origin[i]
        }
    }
}

//改进版 同时找最大最小 提升不大
func SelectGoodSort(origin []int) {
    length := len(origin)
    for i := 0; i < length/2; i++ {
        minIndex := i
        maxIndex := i
        for j := i + 1; j < length-i; j++ {
            if origin[j] > origin[maxIndex] {
                maxIndex = j
            }
            if origin[j] < origin[minIndex] {
                minIndex = j
            }
        }
        // 如果最大值是开头的元素，而最小值不是最尾的元素
        // 先将最大值和最尾的元素交换
        if maxIndex == i && minIndex != length-i-1 {
            origin[length-i-1], origin[maxIndex] = origin[maxIndex], origin[length-i-1]
            origin[i], origin[minIndex] = origin[minIndex], origin[i] // 然后最小的元素放在最开头
        } else if maxIndex == i && minIndex == length-i-1 {
            origin[maxIndex], origin[minIndex] = origin[minIndex], origin[maxIndex]
        } else {
            origin[i], origin[minIndex] = origin[minIndex], origin[i]
            origin[length-i-1], origin[maxIndex] = origin[maxIndex], origin[length-i-1]
        }
    }
}
func main() {
    list := []int{5}
    SelectGoodSort(list)
    fmt.Println(list)

    list1 := []int{5, 9}
    SelectGoodSort(list1)
    fmt.Println(list1)

    list2 := []int{5, 9, 1}
    SelectGoodSort(list2)
    fmt.Println(list2)

    list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
    SelectGoodSort(list3)
    fmt.Println(list3)

    list4 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6}
    SelectGoodSort(list4)
    fmt.Println(list4)
}
