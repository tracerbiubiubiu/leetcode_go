package main

import "fmt"

/*
9. Consecutive Integer
*/
func consecutiveInteger(num int) bool {
    //只要不是2的n次方 都符合要求
    return num&(num-1) != 0
}
func main() {
    var num int
    fmt.Scanf("%d", &num)
    isValid := consecutiveInteger(num)
    if isValid {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
