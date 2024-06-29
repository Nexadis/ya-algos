package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	t := NewTree()
	for i := 0; i < n-1; i++ {
		var child, parent string
		fmt.Fscan(r, &child, &parent)
		t.Add(child, parent)

	}
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	t.Walk(incChilds)

	t.Walk(func(n *Node) {
		fmt.Printf("%s %d\n", n.Name, n.Childs)
	})
}

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Name   string
	Childs int
}

type Tree struct {
	root *Node
	mu   sync.Mutex
	size int
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Add(key, parent string) *Node {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.size++

	if t.root == nil {
		t.root = &Node{Name: key}
		p := add(t.root, parent)
		t.root.Parent = p
		return t.root
	}
	p := add(t.root, parent)
	c := add(t.root, key)
	c.Parent = p
	return c
}

func add(root *Node, key string) *Node {
	switch {
	case key < root.Name:
		if root.Left == nil {
			root.Left = &Node{Name: key}
			return root.Left
		}
		return add(root.Left, key)
	case key > root.Name:
		if root.Right == nil {
			root.Right = &Node{Name: key}
			return root.Right
		}
		return add(root.Right, key)
	case key == root.Name:
		return root
	}
	return nil
}

func incChilds(p *Node) {
	for p.Parent != nil {
		p = p.Parent
		p.Childs++
	}
}

func (t *Tree) Get(key string) *Node {
	t.mu.Lock()
	defer t.mu.Unlock()
	return get(t.root, key)
}

func get(root *Node, key string) *Node {
	if root == nil {
		return nil
	}
	switch {
	case root.Name == key:
		return root
	case key < root.Name:
		return get(root.Left, key)
	case key > root.Name:
		return get(root.Right, key)
	}
	return nil
}

func (t *Tree) Walk(do func(*Node)) {
	t.mu.Lock()
	defer t.mu.Unlock()
	walk(t.root, do)
}

func walk(root *Node, do func(*Node)) {
	if root.Left != nil {
		walk(root.Left, do)
	}
	do(root)
	if root.Right != nil {
		walk(root.Right, do)
	}
}
