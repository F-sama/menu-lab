package main

import (
    "fmt"
    "errors"
)

//定义数据类型
type data interface{}

//定义节点
type linkNode struct {
    Data data   //数据
    Next *linkNode //指向下一个节点
}

//链表头
type HeadLinkNode struct {
    Length int
    Node   *linkNode
}

//定义节点接口
type LinkNoder interface {
    Add(node data)               //插入元素至表尾
    Remove(index int) error         //删除指定位置的节点
    Insert(index int, node data) //指定位置插入节点
    Len() int                       //获取链表长度
    Search(node data) int        //获取指定元素的位置
    Get(index int) *linkNode        //获取指定位置的元素
}

func New() LinkNoder {
    return &HeadLinkNode{Length: 0, Node: &linkNode{Data: "", Next: nil}}
}

//链表尾插入节点
func (h *HeadLinkNode) Add(node data) {
    l := h.Node
    for {
        if l.Next == nil {
            newNode := &linkNode{Data: node, Next: nil}
            l.Next = newNode
            break
        } else {
            l = l.Next
        }
    }
    h.Length++
}

//删除指定位置的节点，节点从1开始计算
func (h *HeadLinkNode) Remove(index int) error {
    l := h.Node
    lLen := h.Len()
    if index < 1 || index > lLen {
        return errors.New("index error")
    }
    for i := 0; i < index-1; i++ {
        l = l.Next
    }
    if index == lLen {
        l.Next = nil
    } else {
        l.Next = l.Next.Next
    }
    h.Length--
    return nil
}

//在链表的指定位置增加节点
func (h *HeadLinkNode) Insert(index int, node data) {
    l := h.Node
    lLen := h.Len()
    if index < 1 || index > lLen {
        return
    }
    for i := 0; i < index-1; i++ {
        l = l.Next
    }
    newNode := &linkNode{Data: node, Next: l.Next}
    l.Next = newNode
    h.Length++
}

//获取链表的长度
func (h *HeadLinkNode) Len() int {
    l := h.Node
    i := 0
    for l.Next != nil {
        i++
        l = l.Next
    }
    if i != h.Length {
        fmt.Println("------")
        fmt.Println(i)
        fmt.Println(h.Length)
        fmt.Println("------")
    }
    return i
}

//获取指定元素的位置
func (h *HeadLinkNode) Search(node data) int {
    l := h.Node
    i := 0
    for {
        i++
        if l.Next != nil && l.Next.Data == node {
            return i
        } else {
            return 0
        }
    }
    return 0
}

//获取指定位置的元素
func (h *HeadLinkNode) Get(index int) *linkNode {
    lLen := h.Len()
    if index < 1 || index > lLen {
        return nil
    }
    l := h.Node
    for i := 0; i < index; i++ {
        l = l.Next
    }
    return l
}

func main() {
    l := New()
    l.Add(data("golang"))
    fmt.Println(l.Search(data("golang")))
    l.Remove(1)
    fmt.Println(l.Len())
    l.Add(data("java"))
    fmt.Println(l.Len())
    l.Insert(1, data("c++"))
    fmt.Println(l.Len())
    fmt.Println(l.Remove(l.Len()))
    fmt.Println(l.Len())
    fmt.Println(l.Search(data("java")))
    fmt.Println(l.Search(data("c++")))
    fmt.Println(l.Get(l.Len()))
}

