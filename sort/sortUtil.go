package main

import "fmt"

func swap(array []int, a, b int) []int {
    if a == b {
        return array
    }
    array[a], array[b] = array[b], array[a]
    return array
}
func main() {
    a := []int{1, 2, 3}
    b := swap(a, 0, 1)
    fmt.Println(b)
}
