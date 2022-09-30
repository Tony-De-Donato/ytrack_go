package main

import (
	"fmt"
)

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

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	element := l.Head
	var precedent *NodeL
	for element != nil {
		if element.Data == data_ref {
			if precedent == nil {
				l.Head = element.Next
			} else {
				precedent.Next = element.Next
			}
		} else {
			precedent = element
		}
		element = element.Next
	}

}

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func main() {
	link := &List{}
	link2 := &List{}

	fmt.Println("----normal state----")
	ListPushBack(link2, 1)
	PrintList(link2)
	ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	ListPushBack(link, 1)
	ListPushBack(link, "Hello")
	ListPushBack(link, 1)
	ListPushBack(link, "There")
	ListPushBack(link, 1)
	ListPushBack(link, 1)
	ListPushBack(link, "How")
	ListPushBack(link, 1)
	ListPushBack(link, "are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)
	PrintList(link)

	ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
}
