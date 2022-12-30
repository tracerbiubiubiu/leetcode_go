package main

/*
给定一个头结点为 head 的非空单链表，返回链表的中间结点。

如果有两个中间结点，则返回第二个中间结点。
*/
func middleNode(head *ListNode) *ListNode {
    slow, fast := head.Next, head.Next
    if slow == nil {
        return head
    }
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}
