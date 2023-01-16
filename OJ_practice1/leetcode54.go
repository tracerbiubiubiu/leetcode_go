package main

import "fmt"

//func spiralOrder(matrix [][]int) []int {
//    m, n := len(matrix), len(matrix[0])
//    start := 0
//    res := []int{}
//    end := (minOf2(m, n)+1)/2 - 1 //最后一波环对应起始点下标
//    for ; start <= end; start++ {
//        tmp := []int{}
//        //i  m-i行   i   n-i列
//        //起点--》右边
//        for i, j := start, start; j <= n-start-1; j++ {
//            tmp = append(tmp, matrix[i][j])
//        }
//        //if start <= end
//        //右上顶点--》右下
//        for i, j := start+1, n-1-start; i <= m-start-1; i++ {
//            tmp = append(tmp, matrix[i][j])
//        }
//        //右下至左下 从右往左，如果这一层只有1行，那么第一个循环已经将该行打印了，这里就不需要打印了，即 （m-1-i ）!= i
//        for i, j := m-start-1, n-start-1-1; m-start-1 != start && j >= start; j-- {
//            tmp = append(tmp, matrix[i][j])
//        }
//        //左下至左上 这里要少遍历一个 否则真就便利了一整个环 从下往上，如果这一层只有1列，那么第2个循环已经将该列打印了，这里不需要打印，即(n-1-i) != i
//        for i, j := m-start-1-1, start; n-start-1 != start && i > start; i-- {
//            tmp = append(tmp, matrix[i][j])
//        }
//        res = append(res, tmp...)
//    }
//    return res
//}
func minOf2(a, b int) int {
    if a > b {
        return b
    }
    return a
}
func main() {
    matrix := [][]int{
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12},
        //{1},
        //{2},
        //{3},
        //{4},
        //{1, 2},
        //{3, 4},
    }
    fmt.Println(spiralOrder(matrix))
}

/*
func spiralOrder(matrix [][]int) (res []int) {
	// 每一圈
	for i := 0; i < len(matrix[0])/2+1; i++ {
		// 上
		for j := i; j < len(matrix[0])-i-1 && len(matrix)-1-i >= i; j++ {
			res = append(res, matrix[i][j])
		}
        // 右
		for j := i; j < len(matrix)-i && i <= len(matrix[0])-1-i; j++ {
			res = append(res, matrix[j][len(matrix[0])-i-1])
		}
        // 下
		for j := len(matrix[0]) - i - 2; j > i && len(matrix)-1-i > i; j-- {
			res = append(res, matrix[len(matrix)-1-i][j])
		}
        // 左
		for j := len(matrix) - 1 - i; j > i && i < len(matrix[0])-1-i; j-- {
			res = append(res, matrix[j][i])
		}
	}
	return
}

*/
//遍历m*n次
func spiralOrder(matrix [][]int) []int {
    res := []int{}
    m, n := len(matrix), len(matrix[0])

    top, down, left, right := 0, m-1, 0, n-1
    for count := 0; count < m*n; {
        for i, j := top, left; j <= right; j++ {
            res = append(res, matrix[i][j])
            count++
        }
        top++
        for i, j := top, right; i <= down; i++ {
            res = append(res, matrix[i][j])
            count++
        }
        right--
        for i, j := down, right; top <= down && j >= left; j-- {
            res = append(res, matrix[i][j])
            count++
        }
        down--
        for i, j := down, left; left <= right && i >= top; i-- {
            res = append(res, matrix[i][j])
            count++
        }
        left++
    }
    return res
}
