package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

/*

 */
//func insertionSortList(head *ListNode) *ListNode {
//    listArry := make([]*ListNode, 0)
//    if head == nil {
//        return nil
//    }
//    for head != nil {
//        listArry = append(listArry, head)
//        head = head.Next
//    }
//    for i := range listArry {
//        cur := listArry[i]
//        j := i - 1
//        if i == 0 {
//            continue
//        }
//        for ; j >= 0; j-- {
//            if cur.Val < listArry[j].Val {
//                listArry[j], listArry[j+1] = listArry[j+1], listArry[j]
//                cur.Next = listArry[j]
//            }
//        }
//    }
//
//    return listArry[0]
//}
//func insertionSortList(h *ListNode) *ListNode {
//    dummy := &ListNode{0, h}
//    for cur := dummy.Next; cur != nil && cur.Next != nil; {
//        if cur.Val <= cur.Next.Val {
//            cur = cur.Next
//            continue
//        } // 找到了待颠倒的cur与cur.Next节点
//        tmp := cur.Next          // 记录，以备后用
//        cur.Next = cur.Next.Next // 删除上一行备用的cur.Next节点
//        prev := dummy
//        for prev.Next.Val < tmp.Val {
//            prev = prev.Next
//        } // 找到tmp的'目标前驱节点pre'了。期望组成prev=>tmp=>prev.Next的组合
//        tmp.Next = prev.Next
//        prev.Next = tmp
//    }
//    return dummy.Next
//}

func insertionSortList(head *ListNode) *ListNode {
    res := &ListNode{Next: head}
    if head == nil {
        return nil
    }
    cur := head.Next
    end := head
    for cur != nil {
        if end.Val <= cur.Val {
            //zheli别忘啦
            end = cur
            cur = cur.Next
            continue
        } else {
            tmp := cur
            prev := res
            end.Next = cur.Next //4->2
            for prev.Next.Val < tmp.Val {
                prev = prev.Next
            }
            //prev->4->3->2->1
            //3->4
            cur.Next = prev.Next
            //prev->3 prev->3->4->2->1
            prev.Next = cur
        }
        cur = end.Next
    }
    return res.Next
}
func main() {
    a, b, c, d, e := -1, 5, 3, 4, 0
    e1 := &ListNode{Val: e}
    d1 := &ListNode{Val: d, Next: e1}
    c1 := &ListNode{Val: c, Next: d1}
    b1 := &ListNode{Val: b, Next: c1}
    a1 := &ListNode{Val: a, Next: b1}
    res := insertionSortList(a1)
    fmt.Println(res.Next)
}
