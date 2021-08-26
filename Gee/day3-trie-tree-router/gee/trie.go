package gee

import (
	"strings"
)

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// /hello
//
// (root node #0)height 0 -> check child: no
//							 create child #01: part hello , isWild false
//							 (root node)#0 add (child)#01
//							 (child)#01 insert, height ++
// (child node #01)height 1 -> len(parts) == 1
// 							  (child)#01 pattern set
//							  complete
//
// /hello/world
//
// (root node #0)height 0 -> check child: yes (child)#01
//							 (child)#01 insert, height ++
// (child node #01)height 1 -> check child: no
//							  create child #11: part world, isWild false
//							  (child)#01 add (child)#11
//							  (child)#11 insert, height ++
// (child node #11)height 2 -> len(parts) == 2
//							   (child)#11 pattern set
//							   complete
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
