package main

import (
    "fmt"
    "math/rand"
)

/*
https://leetcode.cn/problems/random-flip-matrix/
*/
//可以降维
//map没事干别初始化长度，会默认对int string类型赋值
type Solution struct {
    m, n     int
    remain   int
    isFliped map[int]int
}

func Constructor(m int, n int) Solution {

    return Solution{
        m:        m,
        n:        n,
        remain:   m * n,
        isFliped: make(map[int]int),
    }
}

/*
https://leetcode.cn/problems/random-flip-matrix/solutions/1124464/pythonjavajavascriptgo-jiang-wei-ha-xi-b-8ipu/
*/

func (this *Solution) Flip() []int {
    ans := []int{}
    if this.remain == 0 {
        return nil
    }
    randNu := rand.Intn(this.remain) //rand的值也是下标
    this.remain--
    if v, ok := this.isFliped[randNu]; !ok {
        i := randNu / this.n
        j := randNu % this.n
        ans = []int{i, j}
    } else {
        i := v / this.n
        j := v % this.n
        ans = []int{i, j}
    }
    //如果rand对应的下标指向队尾，保持不变
    //否则把队尾元素与rand作交换，即rand下标对应remain下标的队尾元素
    if v, ok := this.isFliped[this.remain]; ok {
        this.isFliped[randNu] = v
    } else {
        //map的值是对应下标，把当前下标的值指向队尾
        this.isFliped[randNu] = this.remain
    }

    return ans
}

func (this *Solution) Reset() {
    this.remain = this.m * this.n
    this.isFliped = make(map[int]int)
}
func main() {
    solution := Constructor(2, 2)
    fmt.Println(solution.Flip())
    fmt.Println(solution.Flip())
    fmt.Println(solution.Flip())
    fmt.Println(solution.Flip())
}
