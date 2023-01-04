package main

/*
https://leetcode.cn/problems/diagonal-traverse/solutions/
*/
func findDiagonalOrder(mat [][]int) []int {
    m,n := len(mat),len(mat[0])
    //
    res := make([]int,0,m*n)
    count := m+n-1
    for nu:=0;nu<count;nu++{
        //向上
        if nu%2 ==0{
            curI := minOf2(nu,m-1)
            curJ:= maxOf2(0,nu-m+1)

            for curI >=0 && curJ<n{
                res = append(res,mat[curI][curJ])
                curI--
                curJ++
            }
        }else {
            //向下
            curI :=maxOf2(0,nu-n+1)
            curJ :=minOf2(nu,n-1)

            for curI<m  && curJ>=0{
                res = append(res,mat[curI][curJ])
                curI++
                curJ--
            }
        }
    }
    return res
}
func minOf2(a, b int) int {
    if a > b {
        return b
    }
    return a
}
func maxOf2(a, b int) int {
    if a > b {
        return a
    }
    return b
}
