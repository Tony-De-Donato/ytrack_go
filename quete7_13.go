package main

import (
	"fmt"
)

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	valeur := l
	for valeur != nil {
		noeud_suivant := valeur.Next
		for noeud_suivant != nil {
			if valeur.Data > noeud_suivant.Data {
				var1 := valeur.Data
				valeur.Data = noeud_suivant.Data
				noeud_suivant.Data = var1
			}
			noeud_suivant = noeud_suivant.Next
		}
		valeur = valeur.Next
	}
	return l
}

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *NodeI

	link = listPushBack(link, 5)
	link = listPushBack(link, 4)
	link = listPushBack(link, 3)
	link = listPushBack(link, 2)
	link = listPushBack(link, 1)

	PrintList(ListSort(link))
}
