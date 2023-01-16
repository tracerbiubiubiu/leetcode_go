package main

/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
*/
//1
//func reverseList(head *ListNode) *ListNode {
//    if head == nil {
//        return nil
//    }
//    cur := head
//    var prev *ListNode
//    prev = nil
//    //prev := new(ListNode)
//    //prev = prev.Next
//    for ; cur != nil; {
//        a := cur.Next
//        cur.Next = prev
//        prev = cur
//        cur = a
//    }
//    //注意把尾部置空
//    //head.Next = nil
//    return prev
//}
//2递归
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    ret := reverseList(head.Next)
    head.Next.Next = head

    head.Next = nil
    return ret
}
