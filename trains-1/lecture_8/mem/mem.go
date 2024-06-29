package main

import "fmt"

type node struct {
	key   int
	index int
	left  int
	right int
}

type memory struct {
	nodes    []node
	lastFree int
}

func InitMem(size int) memory {
	nodes := make([]node, size)
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].left = i + 1
		nodes[i].index = i
	}
	nodes[len(nodes)-1].left = -1
	nodes[len(nodes)-1].index = len(nodes) - 1
	return memory{
		nodes:    nodes,
		lastFree: 0,
	}
}

func (m *memory) alloc() *node {
	if m.lastFree == -1 {
		panic("out of memory")
	}
	allocated := &m.nodes[m.lastFree]
	m.lastFree = m.nodes[m.lastFree].left
	fmt.Printf("alloc mem block, i=%d, lastFree=%d\n", allocated.index, m.lastFree)
	return allocated
}

func (m *memory) free(n *node) {
	fmt.Printf("free mem block, i=%d\n", n.index)
	m.nodes[n.index].left = m.lastFree
	m.lastFree = n.index
}

func main() {
	mem := InitMem(10)
	n1 := mem.alloc()
	n2 := mem.alloc()
	n1.key = 1234
	n2.key = 4123
	mem.free(n2)
	mem.free(n1)
	fmt.Println(mem)
	n1 = mem.alloc()
	n2 = mem.alloc()
	fmt.Println(n1.key)
	fmt.Println(n2.key)
}
