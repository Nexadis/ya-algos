package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	t := NewTree()
	for {
		var x int
		fmt.Fscan(r, &x)
		if x == 0 {
			break
		}
		t.Add(x)
	}
	t.Del(t.Max())
	secMax := t.Max()
	fmt.Fprintln(w, secMax)
}

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Key    int
}

type Tree struct {
	root *Node
	mu   sync.Mutex
	size int
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Add(key int) int {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.size++

	if t.root == nil {
		t.root = &Node{Key: key}
		return 1
	}

	return add(t.root, key, 1)
}

func add(root *Node, key int, level int) int {
	switch {
	case key < root.Key:
		if root.Left == nil {
			root.Left = &Node{Parent: root, Key: key}
			return level + 1
		}
		return add(root.Left, key, level+1)
	case key > root.Key:
		if root.Right == nil {
			root.Right = &Node{Parent: root, Key: key}
			return level + 1
		}
		return add(root.Right, key, level+1)
	}
	return -1
}

func (t *Tree) Get(key int) *Node {
	t.mu.Lock()
	defer t.mu.Unlock()
	return get(t.root, key)
}

func get(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	switch {
	case root.Key == key:
		return root
	case key < root.Key:
		return get(root.Left, key)
	case key > root.Key:
		return get(root.Right, key)
	}
	return nil
}

func (t *Tree) Del(key int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	n := get(t.root, key)
	if n != nil {
		t.size--
	}
	newroot := del(n)
	if n == t.root {
		t.root = newroot
	}
}

func del(root *Node) *Node {
	if root == nil {
		return nil
	}

	switch {
	case root.Left == nil && root.Right == nil:
		if root.Parent == nil {
			return nil
		}
		delno(root)
		return root
	case root.Left == nil || root.Right == nil:
		child := root.Left
		if root.Left == nil {
			child = root.Right
		}
		if root.Parent == nil {
			return child
		}
		del1(root)
		return root

	}
	del2(root)
	return root
}

func delno(root *Node) {
	if root.Parent.Left == root {
		root.Parent.Left = nil
		return
	}
	root.Parent.Right = nil
}

func del1(root *Node) {
	var rewrite *Node
	if root.Left == nil {
		rewrite = root.Right
	}
	if root.Right == nil {
		rewrite = root.Left
	}

	if root.Parent.Left == root {
		root.Parent.Left = rewrite
		return
	}
	root.Parent.Right = rewrite
}

func del2(root *Node) {
	right := root.Right
	left := right
	for left.Left != nil {
		left = left.Left
	}

	root.Key = left.Key
	del(left)
}

func (t *Tree) Max() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	m := t.root
	for m.Right != nil {
		m = m.Right
	}
	return m.Key
}
