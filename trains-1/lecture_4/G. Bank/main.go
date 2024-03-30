package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	solution()
}

func solution() {
	b := newBank()
	defer b.Close()
	for {
		cmd, ok := b.ReadCmd()
		if !ok {
			return
		}
		b.HandleCmd(cmd)
	}
}

type bank struct {
	clients map[string]int
	s       *bufio.Scanner
	w       *bufio.Writer
}

func newBank() bank {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	m := map[string]int{}
	return bank{
		clients: m,
		s:       s,
		w:       w,
	}
}

func (b bank) ReadCmd() (string, bool) {
	if !b.s.Scan() {
		return "", false
	}
	return b.s.Text(), true
}

func (b bank) HandleCmd(cmd string) {
	switch cmd {
	case "DEPOSIT":
		b.deposit()
	case "WITHDRAW":
		b.withdraw()
	case "BALANCE":
		b.balance()
	case "TRANSFER":
		b.transfer()
	case "INCOME":
		b.income()
	default:
	}
}

func (b bank) deposit() {
	b.s.Scan()
	name := b.s.Text()
	b.s.Scan()
	val, _ := strconv.Atoi(b.s.Text())
	b.clients[name] += val
}

func (b bank) withdraw() {
	b.s.Scan()
	name := b.s.Text()
	b.s.Scan()
	val, _ := strconv.Atoi(b.s.Text())
	b.clients[name] -= val
}

func (b bank) balance() {
	b.s.Scan()
	name := b.s.Text()
	val, ok := b.clients[name]
	if !ok {
		fmt.Fprintln(b.w, "ERROR")
		return
	}
	fmt.Fprintln(b.w, val)
}

func (b bank) transfer() {
	b.s.Scan()
	from := b.s.Text()
	b.s.Scan()
	to := b.s.Text()
	b.s.Scan()
	sum, _ := strconv.Atoi(b.s.Text())
	b.clients[from] -= sum
	b.clients[to] += sum
}

func (b bank) income() {
	b.s.Scan()
	percent, _ := strconv.Atoi(b.s.Text())
	for client, sum := range b.clients {
		if sum <= 0 {
			continue
		}
		sum = int(float64(sum) * float64(percent) / 100)
		b.clients[client] += sum
	}
}

func (b bank) Close() error {
	return b.w.Flush()
}
