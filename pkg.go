package junk

import (
	"fmt"
	"time"
)

type Node struct {
	Data int
	Next *Node
}

func (n *Node) AddNode(data int) {
	newNode := Node{data, nil}
	iter := n
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = &newNode
}

func (n *Node) PrintNode() {
	iter := n
	for iter != nil {
		fmt.Println(iter.Data)
		iter = iter.Next
	}
}

func init() {

}

func Junk() string {
	currentTime := time.Now()
	return fmt.Sprintf("Junk: %s", currentTime)
}
