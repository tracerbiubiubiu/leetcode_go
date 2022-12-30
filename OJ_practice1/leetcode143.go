package main

/*
143. 重排链表
中等

给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



示例 1：
*/

//func reorderList(head *ListNode) {
//    if head == nil {
//        return
//    }
//    nodes := []*ListNode{}
//    for node := head; node != nil; node = node.Next {
//        nodes = append(nodes, node)
//    }
//    i, j := 0, len(nodes)-1
//    for i < j {
//        nodes[i].Next = nodes[j]
//        i++
//        nodes[j].Next = nodes[i]
//        j--
//    }
//    nodes[i].Next = nil
//}
//方法二：寻找链表中点 + 链表逆序 + 合并链表
func reorderList(head *ListNode) {
    mid := findMidNode(head)
    l2 := reverseTheList(mid.Next)
    //这步别忘啦
    mid.Next = nil
    mergeList(head, l2)
}
func findMidNode(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}
func reverseTheList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    prev, cur := head, head
    for cur != nil {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    head.Next = nil
    return prev
}
func mergeList(l1, l2 *ListNode) {
    for l1 != nil && l2 != nil {
        l1Temp := l1.Next
        l2Temp := l2.Next
        l1.Next = l2
        l1 = l1Temp
        l2.Next = l1
        l2 = l2Temp
    }
}
