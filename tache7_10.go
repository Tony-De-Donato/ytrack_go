package main

import (
	"fmt"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListPushBack(l *List, data interface{}) {
	value := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = value
		l.Tail = value
	} else {
		l.Tail.Next = value
		l.Tail = value
	}

}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	var rep *interface{}
	test := false
	for l.Head != nil {
		if comp(l.Head.Data, ref) {
			rep = &l.Head.Data
			test = true
		}
		if test {
			break
		}
		l.Head = l.Head.Next
	}
	return rep
}

func main() {
	link := &List{}

	ListPushBack(link, "hello")
	ListPushBack(link, "hello1")
	ListPushBack(link, "hello2")
	ListPushBack(link, "hello3")

	found := ListFind(link, interface{}("hello2"), CompStr)

	fmt.Println(found)
	fmt.Println(*found)
}
