package dns

import (
	"fmt"
)

type Node struct {
	list  []string
	index map[string]int
}

func NewNode() *Node {
	return &Node{index: make(map[string]int)}
}

func (node *Node) Add(ip string) int {
	id := len(node.list)
	if id > 255 {
		node.index[ip] = -1
		return -1
	}
	node.list = append(node.list, ip)
	node.index[ip] = id
	return id
}

func (node *Node) GetID(ip string) string {
	id, ok := node.index[ip]
	if !ok {
		id = node.Add(ip)
	}
	if id < 0 {
		return ""
	}
	return fmt.Sprintf("%02X", id)
}
