package main

import (
	"sync"
)

// TODO: создать дерево на основе Пула памяти
// TODO: обход дерева для возврата сортированного списка
// TODO: дерево Хаффмана, сериализация и десериализация деревьев

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

func (t *Tree) Add(key int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.size++

	if t.root == nil {
		t.root = &Node{Key: key}
		return
	}

	add(t.root, key)
}

func add(root *Node, key int) {
	switch {
	case key < root.Key:
		if root.Left == nil {
			root.Left = &Node{Parent: root, Key: key}
			return
		}
		add(root.Left, key)
	case key > root.Key:
		if root.Right == nil {
			root.Right = &Node{Parent: root, Key: key}
			return
		}
		add(root.Right, key)
	}
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
	del(n)
}

func del(root *Node) {
	if root == nil {
		return
	}
	switch {
	case root.Left == nil && root.Right == nil:
		delno(root)
		return
	case root.Left == nil || root.Right == nil:
		del1(root)
		return

	}
	del2(root)
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

func (t *Tree) Walk() []int {
	t.mu.Lock()
	defer t.mu.Unlock()
	ordered := make([]int, 0, t.size)
	return walk(t.root, ordered)
}

func walk(root *Node, ordered []int) []int {
	if root.Left != nil {
		ordered = walk(root.Left, ordered)
	}
	ordered = append(ordered, root.Key)
	if root.Right != nil {
		ordered = walk(root.Right, ordered)
	}
	return ordered
}
