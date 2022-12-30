package main

import "fmt"

/*
golang数据结构 单链表 这里定义成头结点有data
*/
type Node struct {
    data interface{}
    next *Node
}
type List struct {
    headNode *Node //定义头结点
}

//func (l *List) initList() *List {
//    newList := List{headNode: l.headNode}
//    return &newList
//}
func initList() *List{
    headnode := &Node{}
    newList := &List{headNode: headnode}
    return newList
}
func (l *List) IsEmpty() bool {
    if l.headNode.data == nil {
        return true
    }
    return false
}
func (l *List) GetLength() int {
    cur := l.headNode
    i := 0
    if l.IsEmpty() {
        return 0
    }
    for cur != nil {
        i += 1
        cur = cur.next
    }
    return i
}

// Add 头部添加
func (l *List) Add(data interface{}) {
    newHead := &Node{data: data}
    newHead.next = l.headNode
    l.headNode = newHead
    //return newHead
}

// Append 尾部添加
func (l *List) Append(data interface{}) {
    newTail := &Node{data: data}
    if l.IsEmpty() {
        l.headNode = newTail
    } else {
        cur := l.headNode
        for cur.next != nil {
            cur = cur.next
        }
        cur.next = newTail
        //newTail.next = nil
    }
}

// Insert 中间添加
func (l *List) Insert(key int, data interface{}) {
    if l.IsEmpty() || key <= 0 {
        l.Add(data)
    } else if key > l.GetLength() {
        l.Append(data)
    } else {
        cur := l.headNode
        for i := 0; i < key -1 ; i++ {
            cur = cur.next
        }
        Node := &Node{data: data}
        Node.next = cur.next
        cur.next = Node
    }
}

// DelFromHead 删除头部
func (l *List) DelFromHead() {
    cur := l.headNode
    if cur != nil {
        l.headNode = cur.next
    }
}

// DelFromTail 删除尾部
func (l *List) DelFromTail() {
    cur := l.headNode
    if cur.next.next != nil {
        cur = cur.next
    }
    cur.next = nil
}

// Del 删除指定值
func (l *List) Del(data interface{}) {
    cur := l.headNode
    //如果删除的是头结点
    if cur.data == data {
        l.DelFromHead()
    } else {
        for cur.next != nil {
            if cur.next.data == data {
                //cur.next.data = cur.next.next.data
                cur.next = cur.next.next
            } else {
                cur = cur.next
            }

        }
    }
}

// Remove 删除指定位置
func (l *List) Remove(index int) {
    //注意起点从哪里开始 从head还是从head之后的第一个节点
    //cur := l.headNode.next
    cur := l.headNode
    count := 0
    if !l.IsEmpty() {
        if index <= 0 {
            l.DelFromHead()
        } else if index > l.GetLength() {
            l.DelFromTail()
        } else {
            for cur.next.next != nil && count != index-1 {
                    count += 1
                    cur = cur.next
            }
            cur.next = cur.next.next
        }
    }
}

//查询
func (l *List) Contain(data interface{}) bool {
    cur := l.headNode
    for cur != nil {
        if cur.data == data {
            return true
        } else {
            cur = cur.next
        }
    }
    return false
}

// ShowList 遍历
func (l *List) ShowList() {
    cur := l.headNode
    if !l.IsEmpty() {
        for cur != nil {
            fmt.Printf("%+v", cur.data)
            cur = cur.next
        }
    }
}
//翻转 迭代/211
//初始化两个变量，pre和cur,  遍历node，将node的next指向前一个及诶单，也就是pre,  将pre赋值为本节点，cur为本节点的next节点
func reverse(head *Node) *Node{
    cur := head
    var pre *Node = nil

    for cur != nil{
        //pre, cur ,cur.next = cur, cur.next, pre
        next := cur.next
        cur.next = pre
        pre = cur
        cur = next
    }
    //翻转后的头结点
    return pre
}

//反转 递归
func reverseList(head *Node) *Node{
    if head == nil{
       return nil
    }
    if head.next == nil{
        return head
    }
    newHead := reverseList(head.next)
    head.next.next = head
    head.next = nil
    return newHead
}
//打印指定节点后面的链表
func showFromHead(head *Node) {

    cur := head
    for cur != nil{
        fmt.Printf("%v", cur.data)
        cur = cur.next
    }
}
//测试
//func main() {
//    //var list *List
//    //list1 = &List{}
//    //list := &List{}
//    //list1 := list.initList()
//    list1 := initList()
//
//    //a.Append(1)
//    list1.Append(2)
//
//    list1.ShowList()
//    list1.Append(3)
//    //a.ShowList()
//
//    list1.ShowList()
//    list1.Add(1)
//    list1.ShowList()
//    list1.DelFromTail()
//    list1.ShowList()
//    fmt.Println(list1.Contain(1))
//    fmt.Println(list1.GetLength())
//    list1.ShowList()
//    list1.Append(3)
//    //list1.Insert(5,3)
//    fmt.Println("***")
//    list1.ShowList()
//    fmt.Println("***")
//    list1.Remove(3)
//    //list1.Del(3)
//    //list1.DelFromTail()
//    list1.ShowList()
//    list1.Insert(3,3)
//    list1.ShowList()
//    fmt.Println("---")
//    //fmt.Println(list1.headNode.next.data)
//    //a := reverse(list1.headNode)
//    a := reverseList(list1.headNode)//这里list1 已经反转过了 变成新的list
//    showFromHead(a)
//    fmt.Println("***")
//
//
//
//
//    //fmt.Println(a.next.data)
//    //list2 := initList()
//    //list2.Append(a.data)
//    //list2.ShowList()
//}
