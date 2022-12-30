package main
/*
树
*/
import (
    "fmt"
)
// BinaryTree 二叉树
type BinaryTree struct {
    root *BinTreeNode
    height int
    size int
}

// BinTreeNode 二叉树节点
type BinTreeNode struct {
    data   interface{}  //数据域
    parent *BinTreeNode //父节点
    lChild *BinTreeNode //左孩子
    rChild *BinTreeNode //右孩子
    height int          //以该结点为根的子树的高度
    size   int          //该结点子孙数(包括结点本身)
}

func NewBinTreeNode(data interface{}) *BinTreeNode{
    return &BinTreeNode{data: data,size: 1,height: 1}
}
// 获得节点的值
func (n *BinTreeNode)getVal() interface{} {
    if n == nil{
        return nil
    }
    return n.data
}
//设置节点数据
func (n *BinTreeNode)setVal(val interface{}) {
    n.data = val
}
// 判断是否有父亲
func (n *BinTreeNode) hasParent() bool {
    return n.parent != nil
}
// 获得父节点
func (n *BinTreeNode) getParent() *BinTreeNode {
    if !n.hasParent(){
        return nil
    }
    return n.parent
}
// 是左节点位置
func (n *BinTreeNode) isLchild() bool {
    return n.hasParent() && n == n.parent.lChild
}
func (n *BinTreeNode) hasLChild() bool {
    return n.lChild != nil
}
// 右节点位置
func (n *BinTreeNode) hasRChild() bool {
    return n.rChild != nil
}
func (n *BinTreeNode) isRchild() bool {
    return n.hasParent() && n == n.parent.rChild
}
// 设置父节点
func (n *BinTreeNode) setParent(node *BinTreeNode)  {
    n.parent = node
}
//获取左节点
func (n *BinTreeNode) getLChild() *BinTreeNode {
    if !n.hasLChild(){
        return nil
    }
    return n.lChild
}
//获取右节点
func (n *BinTreeNode) getRChild() *BinTreeNode {
    if !n.hasRChild(){
        return nil
    }
    return n.rChild
}
// 断开父节点 当个孤儿
func (n *BinTreeNode) beOrphan()  {
    if !n.hasParent(){
        return
    }
    if n.isLchild(){
        n.parent.lChild = nil
    }else {
        n.parent.rChild = nil
    }
    n.parent = nil
}
// 设置左孩子 返回原左孩子
func (n *BinTreeNode) setLChild(node *BinTreeNode) *BinTreeNode {
    oldLNode := n.lChild
    if n.hasLChild(){
        n.lChild.beOrphan()
    }
    if node != nil{
        //先后顺序？
        node.parent = n
        n.lChild = node
    }
    return oldLNode
}
// 设置右孩子 返回原右孩子
func (n *BinTreeNode) setRChild(node *BinTreeNode) *BinTreeNode {
    oldLNode := n.lChild
    if n.hasRChild(){
        n.rChild.beOrphan()
    }
    if node != nil{
        //先后顺序？
        node.parent = n
        n.rChild = node
    }
    return oldLNode
}
// 该节点为根节点的树的高度
func (n *BinTreeNode) getHeight() int{
    return n.height
}
// 更新当前结点及其祖先的高度  真是**啊 为什么不全部更新呢
//func (n *BinTreeNode) upgradeHeight() int {
//   newH := 1
//   if n.hasLChild(){
//       newH = int(math.Max(float64(newH), float64(n.getLChild().getHeight() + 1)))
//   }
//   if n.hasRChild(){
//       newH = int(math.Max(float64(newH), float64(n.getRChild().getHeight() + 1)))
//   }
//   n.height = newH
//   if n.hasParent(){
//       //n.parent.upgradeHeight()
//       n.getParent().upgradeHeight()
//   }
//   return newH
//}
// 有后代
func (n *BinTreeNode) hasChild() bool {
    return n.hasLChild() || n.hasRChild()
}
//由根节点，更新所有子节点的高度 这里的node为根节点
func (n *BinTreeNode) upgradeHeight() int {
   //n.height = 1
   if n.hasChild(){
       if n.hasLChild(){
           n.getLChild().height = n.height + 1
           n.getLChild().upgradeHeight()
       }
       if n.hasRChild(){
           n.getRChild().height = n.height + 1
           n.getRChild().upgradeHeight()
       }
   }
   return maxDepth(n)
}

func (n *BinTreeNode) getSize() int {
    return n.size
}
// 更新当前结点及其祖先的子孙数
//func (n *BinTreeNode) upgradeSize() int {
//    n.size = 1 //初始化自身为1
//    if n.hasLChild(){
//        n.size += n.getLChild().getSize()
//    }
//    if n.hasRChild(){
//        n.size += n.getRChild().getSize()
//    }
//    if n.hasParent(){
//        n.getParent().upgradeSize()
//    }
//    return n.size
//}
// 他的函数不知道在干嘛? 直接拿根节点更新不就完了
func (n *BinTreeNode) upgradeSize() int {
    n.size = 1
    if n.hasLChild(){
        lChildSize := n.getLChild().upgradeSize()
        n.size += lChildSize
    }
    if n.hasRChild(){
        rChildSize := n.getRChild().upgradeSize()
        n.size += rChildSize
    }
    //if n.hasParent(){
    //    n.getParent().upgradeSize()
    //}
    return n.size
}
// IsLeaf 判断是否为叶子结点
func (n *BinTreeNode) IsLeaf() bool {
    return !n.hasLChild() && !n.hasRChild()
}
func NewBinaryTree(root *BinTreeNode) *BinaryTree {
    return &BinaryTree{root: root,size: 1,height: 1}
}
//获得总节点
func (t *BinaryTree) getSize() int {
    return t.root.getSize()
}
//树判空
func (t *BinaryTree) isNull() bool {
    return t.root == nil
}
//获得高度
func (t *BinaryTree) getHeight() int {
    return t.root.getHeight()
}
//获得第一个与数据e相等的节点
func (t *BinaryTree) Find(data interface{}) *BinTreeNode {
    if t.isNull(){
        return nil
    }
    cur := t.root
    return isEqual(data, cur)
}
func isEqual(data interface{}, node *BinTreeNode) *BinTreeNode {
    if node.getVal() == data{
        return node
    }
    if node.hasLChild(){
        lp := isEqual(data, node.getLChild())
        if lp != nil{
            return lp
        }
    }
    if node.hasRChild(){
        rp := isEqual(data, node.getRChild())
        if rp != nil{
            return rp
        }
    }
    return nil
}

// 深度优先遍历需要优先使用栈
// PreOrder 前序遍历
//用链表存储
func (t *BinaryTree) PreOrder() *List {
    newList := initList()
    preOrderRecu(t.root, newList)
    return newList
}
//先序遍历 递归
func preOrderRecu(node *BinTreeNode, l *List)  {
    if node == nil{
        return
    }
    l.Append(node)
    preOrderRecu(node.getLChild(),l)
    preOrderRecu(node.getRChild(),l)
}
//中序
func (t *BinaryTree) inOrder() *List {
    newList := initList()
    inOrderRecu(t.root, newList)
    return newList
}
func inOrderRecu(node *BinTreeNode, l *List) {
    if node == nil{
        return
    }
    inOrderRecu(node.getLChild(), l)
    l.Append(node)
    inOrderRecu(node.getRChild(), l)
}
//后序
func (t *BinaryTree) postOrder() *List {
    newList := initList()
    postOrderRecu(t.root, newList)
    return newList
}
func postOrderRecu(node *BinTreeNode, l *List) {
    if node == nil{
        return
    }
    postOrderRecu(node.getLChild(), l)
    postOrderRecu(node.getRChild(), l)
    l.Append(node)
}

func max(v1, v2 int) int {
    if v1 >= v2 {
        return v1
    }
    return v2
}
//节点最大深度
func maxDepth(root *BinTreeNode) int {
    // 递归终止条件
    if root == nil {
        return 0
    }

    // 递归过程
    leftMaxDepth := maxDepth(root.getLChild())
    rightMaxDepth := maxDepth(root.getRChild())
    // 左右子树的最高深度+1
    return 1 + max(leftMaxDepth, rightMaxDepth)
}
//新建一个遍历的函数，之前的返回链表这里不方便 类型断言转换类型
func (l *List) bianli() []*BinTreeNode{
    if l.IsEmpty(){
        return nil
    }
    cur := l.headNode
    newList := make([]*BinTreeNode,0)
    for cur != nil{
        node,ok := cur.data.(*BinTreeNode);if !ok{ //前面遍历中append的是指针
            continue
        }
        newList = append(newList, node)
        cur = cur.next
    }
    return newList
}

//层序遍历
//不是有height属性么？
func (t *BinaryTree) LayerOrder() [][]interface{} {
    if t.isNull(){
        return nil
    }
    NodeList := t.postOrder().bianli()
    //给一个初始长度要不会越界panic
    result := make([][]interface{},maxDepth(t.root))
    for i :=1;i <= maxDepth(t.root); i++{
        currLevel := make([]interface{}, 0)
        for _,node := range NodeList{
            //fmt.Print(node.data)
            if node.getHeight() == i{
                //fmt.Print(node.data)
                currLevel = append(currLevel, node.data)
                //NodeList = append(NodeList[:1], NodeList[2:])
            }
        result[i-1] = currLevel
        }
    }
    return result
}
//测试
func main() {
    a := NewBinTreeNode(1)
    tree1 := NewBinaryTree(a)
    a.setLChild(NewBinTreeNode(2))
    a.setRChild(NewBinTreeNode(5))
    a.getLChild().setRChild(NewBinTreeNode(3))
    a.getLChild().getRChild().setLChild(NewBinTreeNode(4))
    a.getRChild().setLChild(NewBinTreeNode(6))
    a.getRChild().setRChild(NewBinTreeNode(7))
    a.getRChild().getLChild().setRChild(NewBinTreeNode(9))
    a.getRChild().getRChild().setRChild(NewBinTreeNode(8))

    node2 := a.getLChild()
    node9 := a.getRChild().getLChild().getRChild()

    node5 := a.getRChild()
    node6 := a.getRChild().getLChild()

    // 当孤儿了
    //node9.beOrphan()
    //node7 := a.getRChild().getRChild()

    fmt.Println("结点2是叶子结点吗? ", node2.IsLeaf())
    fmt.Println("结点9是叶子结点吗? ", node9.IsLeaf())
    /**/
    a.upgradeSize()
    //node4.upgradeSize()
    fmt.Println(node9.getSize())
    fmt.Println(node6.size)
    fmt.Println(node5.size)
    //a.upgradeSize()
    //node9.upgradeHeight()

    // 蠢得离谱 原函数不会更新树里面所有节点的高度
    //node8 := a.getRChild().getRChild().getRChild()
    //node4 := a.getLChild().getRChild().getLChild()
    //node4.upgradeHeight()
    //node9.upgradeHeight()
    //node8.upgradeHeight()
    a.upgradeHeight()
    fmt.Println(node9.height)
    //fmt.Println(maxDepth(a))
    fmt.Println("这棵树的结点总数：", tree1.getSize())
    fmt.Println("这棵树的高度：", a.upgradeHeight())

    l := tree1.inOrder()//中序遍历
    fmt.Println("中序遍历: ")
    l.ShowList()

    //for e := l.headNode; e != nil; e = e.next {
    //    obj, _ := e.data.(*BinTreeNode)
    //    fmt.Printf("data:%v\n", *obj)
    //}
    //result := tree1.Find(6)
    //fmt.Printf("结点6的父节点数据：%v\t结点6的右孩子节点数据: %v\n", result.getParent().getVal(), result.getRChild().getVal())
    //tree1.PreOrder().ShowList()
    //fmt.Println("\n")
    //list := l.bianli()
    //for _,i := range list{
    //    fmt.Println(i)
    //}

    fmt.Println("**********\n")
    ////fmt.Println(tree1.bianli())
    fmt.Println((tree1.LayerOrder()))
}

//https://blog.51cto.com/liuxp0827/1378672