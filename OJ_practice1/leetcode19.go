package main

import "fmt"

/*
https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
*/
type ListNode struct {
    Val  int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {
        return nil
    }
    originalHead := head
    length := 0
    for ; head != nil; head = head.Next {
        length++
    }

    tmpHead := &ListNode{Next: originalHead}
    res := tmpHead
    for count := 1; count < length-n+1; count++ {
        //for count := 0; count < length-n; count++ {
        tmpHead = tmpHead.Next
    }
    //fmt.Println(tmpHead.Val)
    //fmt.Println(tmpHead.Next.Next.Val)
    tmpHead.Next = tmpHead.Next.Next

    return res.Next
}
func main() {
    a, b, c, d, e := 1, 2, 3, 4, 5
    nodeE := &ListNode{Val: e}
    fmt.Println(a, b, c, d)
    //nodeD := &ListNode{Val: d, Next: nodeE}
    //nodeC := &ListNode{Val: c, Next: nodeD}
    //nodeB := &ListNode{Val: b, Next: nodeC}
    //nodeA := &ListNode{Val: a, Next: nodeB}
    //removeNthFromEnd(nodeA, 2)
    //for i := nodeA; i != nil; i = i.Next {
    //    fmt.Println(i.Val)
    //}
    removeNthFromEnd(nodeE, 1)
    for i := nodeE; i != nil; i = i.Next {
        fmt.Println(i.Val)
    }
}
