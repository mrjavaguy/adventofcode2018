package main

import (
	"adventofcode2018/input"
	"fmt"
)

func main() {
	data, _ := input.FileToInts("day08/input08.txt")
	//data := []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}
	part1 := Day08Part1(data)
	part2 := Day08Part2(data)
	fmt.Printf("Day 8 Part 1 %v, Part 2 %v", part1, part2)
}

func Day08Part1(data []int) int {
	tree := buildTree(data)
	stack := NewStack()
	traverseTree(tree, stack)
	license := 0
	for {
		n := stack.Pop()
		if n == nil {
			break
		}
		for _, i := range n.metadata {
			license += i
		}
	}
	return license
}

func Day08Part2(data []int) int {
	tree := buildTree(data)
	return calcLicense(tree)
}

func calcLicense(node *node) int {
	license := 0
	for _, i := range node.metadata {
		if node.expectedChildCount != 0 {
			if i <= node.expectedChildCount {
				license += calcLicense(node.children[i-1])
			}
		} else {
			license += i
		}
	}

	return license
}

func traverseTree(node *node, stack *Stack) {
	stack.Push(node)

	for _, c := range node.children {
		traverseTree(c, stack)
	}
}

func buildTree(data []int) *node {
	root := NewNode()
	state := 0
	stack := NewStack()
	stack.Push(root)
	for _, d := range data {
		currentNode := stack.Pop()
		switch state {
		case 0:
			currentNode.expectedChildCount = d
			state = 1
		case 1:
			currentNode.expectedMetadataCount = d
			state = 2
		case 2:
			if len(currentNode.children) == currentNode.expectedChildCount {
				currentNode.metadata = append(currentNode.metadata, d)
				if len(currentNode.metadata) == currentNode.expectedMetadataCount {
					continue
				}
			} else {
				state = 1
				childNode := NewNode()
				childNode.expectedChildCount = d
				currentNode.children = append(currentNode.children, childNode)
				stack.Push(currentNode)
				currentNode = childNode
			}
		}
		stack.Push(currentNode)
	}

	return root
}

type node struct {
	expectedChildCount    int
	expectedMetadataCount int
	metadata              []int
	children              []*node
}

func NewNode() *node {
	return &node{metadata: make([]int, 0), children: make([]*node, 0)}
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*node
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}
