package week319
/*
2471
 */


type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}
func minimumOperations(root *TreeNode) int {
    count := 0
    left := root.Left
    right := root.Right
    for left!=nil||right!=nil{
        leftVal:=0
        rightVal:=0
        if left !=nil{
            leftVal = root.Left.Val
        }
        if right !=nil{
            rightVal = root.Right.Val
        }

        if leftVal < rightVal {
            root.Left,root.Right = root.Right,root.Left
            count++
        }

    }
    return count
}