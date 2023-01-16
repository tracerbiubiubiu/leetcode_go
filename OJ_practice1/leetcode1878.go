package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "sort"
    "strconv"
    "strings"
)

/*
中位数计算
*/
/*
 * Copyright (c) Huawei Technologies Co., Ltd. 2019-2021. All rights reserved.
 * Description: 上机编程认证
 * Note: 缺省代码仅供参考，可自行决定使用、修改或删除
 * 只能import Go标准库
 */

// 待实现函数，在此函数中填入答题代码
func getMedian(numbers []int) int {

    sort.Ints(numbers)
    length := len(numbers) / 3
    if len(numbers)%3 != 0 {
        length = length + 1
    }

    numbers = numbers[:length]
    if len(numbers)&1 != 0 {
        return numbers[len(numbers)/2]
    } else {
        left := numbers[(len(numbers)-1)/2]
        right := numbers[len(numbers)/2]
        if (left+right)&1 == 0 {
            return (left + right) / 2
        }
        return (left + right + 1) / 2
    }

}

func main() {
    reader := bufio.NewReader(os.Stdin)
    numbers := readInputIntArray(reader)
    result := getMedian(numbers)
    fmt.Println(result)
}

func readInputIntArray(reader *bufio.Reader) []int {
    var num int
    if _, err := fmt.Fscanf(reader, "%d\n", &num); err != nil {
        fmt.Println(err.Error())
        return nil
    }
    if num == 0 {
        return []int{}
    }
    lineBuf, err := reader.ReadString('\n')
    if err != nil && err != io.EOF {
        fmt.Println(err.Error())
        return nil
    }
    lineBuf = strings.TrimRight(lineBuf, "\r\n")
    lineBuf = strings.TrimSpace(lineBuf)
    intNums := map2IntArray(lineBuf, " ")
    if len(intNums) != num {
        fmt.Println("int string len is error")
        return nil
    }
    return intNums
}

func map2IntArray(str string, dem string) []int {
    tempArray := strings.Split(str, dem)
    result := make([]int, len(tempArray))
    for index, value := range tempArray {
        value = strings.TrimSpace(value)
        intVal, err := strconv.Atoi(value)
        if err == nil {
            result[index] = intVal
        }
    }
    return result
}
