package junk

import (
	"fmt"
	"time"
)

type Node struct {
	data int
	next *Node
}

func (n *Node) AddNode(data int) {
	newNode := Node{data, nil}
	iter := n
	for iter.next != nil {
		iter = iter.next
	}
	iter.next = &newNode
}

func (n *Node) PrintNode() {
	iter := n
	for iter != nil {
		fmt.Println(iter.data)
		iter = iter.next
	}
}

func init() {

}

func Junk() string {
	currentTime := time.Now()
	return fmt.Sprintf("Junk: %s", currentTime)
}
