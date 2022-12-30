package main

/*
go数据结构 栈
*/
type Stack struct {
    size int
    top  int
    data []interface{}
}

func initStack(size int) *Stack {
    newStack := Stack{}
    newStack.size = size
    newStack.data = make([]interface{}, size)
    return &newStack
}

func (s *Stack) IsEmpty() bool {
    return s.top == 0
}

func (s *Stack) IsFull() bool{
    return s.top == s.size
}
func (s *Stack) GetLength() int{
    return s.top
}
func (s *Stack)Push(i interface{}) bool{
    if s.IsFull(){
        return false
    }
    s.top += 1
    s.data[s.top] = i
    return true
}
func (s *Stack)Pop() (interface{}, bool){
    if s.IsEmpty(){
        return nil, false
    }
    element := s.data[s.top]
    s.top -= 1
    return element, true
}