package main

import "fmt"

/*
https://leetcode.cn/problems/integer-replacement/solutions/
*/
//二进制末尾有多少个0
//改成2个返回值，总长度与末尾是连续0的长度
func getZeroNum(a int) (int, int) {
    tmp := a & -a
    zeroCount := 0
    for tmp > 1 {
        tmp >>= 1
        zeroCount++
    }
    total := 0
    for a > 0 {
        a >>= 1
        total++
    }
    return total, zeroCount
}
func integerReplacement(n int) int {
    res := 0
    if n == 1 {
        return 0
    }
    for n > 1 {
        for n%2 == 0 {
            n = n / 2
            res++
        }
        if n < 2 {
            break
        }
        total1, plusOne := getZeroNum(n + 1)
        total2, subOne := getZeroNum(n - 1)
        //用例为3时会有问题 还应该关注n+1是否进位的问题
        if total1-plusOne < total2-subOne {
            n += 1
            res++
            continue
        }
        n -= 1
        res++
    }

    return res
}
func main() {
    n := 3
    fmt.Println(integerReplacement(n))
}

//wode思路是偶数不用管，奇数的话二进制表现一定为1 看看+1后末尾的0与-1后末尾的0哪个多
//然后取多的那个 /2 剩下的递归直到num==1
//n&(1<<i)
//n& -n v& -v 的作用就是取出右起连续的 0 以及首次出现的 1
//一般N!的质因数2的个数为N - （N二进制表示中1的个数）

/*
func integerReplacement(n int) int {
    if n == 1 {
        return 0
    }
    if n%2 == 0 {
        return 1 + integerReplacement(n/2)
    }
    return 2 + min(integerReplacement(n/2), integerReplacement(n/2+1))
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/integer-replacement/solutions/1108099/zheng-shu-ti-huan-by-leetcode-solution-swef/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
