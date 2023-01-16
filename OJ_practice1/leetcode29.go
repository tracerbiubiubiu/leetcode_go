package main

import "math"

/*
https://leetcode.cn/problems/divide-two-integers/
*/
func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
        if divisor == 1 {
            return math.MinInt32
        }
        if divisor == -1 {
            return math.MaxInt32
        }
    }
    if divisor == math.MinInt32 { // 考虑除数为最小值的情况
        if dividend == math.MinInt32 {
            return 1
        }
        return 0
    }
    if dividend == 0 { // 考虑被除数为 0 的情况
        return 0
    }

    res := 0
    diffSign := true
    if dividend < 0 && divisor > 0 || divisor < 0 && dividend > 0 {
        diffSign = false
    }
    a, b := abs(dividend), abs(divisor)
    for a < b {
        for i := 31; i > 0; i++ {
            if a-(b<<i) >= 0 {
                a -= b << i
                res += 1 << i
            }
        }

    }
    if !diffSign {
        return -res
    }
    return res
}
func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}

//快速乘

/*
func myPow(x float64, n int) float64 {
    if n >= 0 {
        return quickMul(x, n)
    }
    return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, N int) float64 {
    ans := 1.0
    // 贡献的初始值为 x
    x_contribute := x
    // 在对 N 进行二进制拆分的同时计算答案
    for N > 0 {
        if N % 2 == 1 {
            // 如果 N 二进制表示的最低位为 1，那么需要计入贡献
            ans *= x_contribute
        }
        // 将贡献不断地平方
        x_contribute *= x_contribute
        // 舍弃 N 二进制表示的最低位，这样我们每次只要判断最低位即可
        N /= 2
    }
    return ans
}



func quickMul(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    y := quickMul(x, n/2)
    if n%2 == 0 {
        return y * y
    }
    return y * y * x
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/powx-n/solutions/238559/powx-n-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
