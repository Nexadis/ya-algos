package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {

	type test struct {
		k1 int
		m  int
		k2 int
		p2 int
		n2 int
	}
	type want struct {
		p1 int
		n1 int
	}
	tests := []struct {
		t test
		w want
	}{
		{
			t: test{22, 10, 11, 1, 1},
			w: want{1, 0},
		},
		{
			t: test{89, 20, 41, 1, 11},
			w: want{2, 3},
		},
		{
			t: test{11, 1, 1, 1, 1},
			w: want{0, 1},
		},
		{
			t: test{3, 2, 2, 2, 1},
			w: want{-1, -1},
		},
		{
			t: test{30, 10, 15, 1, 2},
			w: want{1, 0},
		},
		{
			t: test{11, 4, 7, 2, 3},
			w: want{3, 3},
		},
		{
			t: test{11, 4, 6, 2, 3},
			w: want{-1, -1},
		},
		{
			t: test{18, 3, 10, 2, 3},
			w: want{-1, -1},
		},
		{
			t: test{18, 3, 10, 2, 2},
			w: want{3, 3},
		},
		{
			t: test{2, 4, 1, 2, 2},
			w: want{-1, -1},
		},
		{
			t: test{2, 4, 1, 1, 1},
			w: want{1, 0},
		},
		{
			t: test{44, 4, 9, 1, 3},
			w: want{0, 3},
		},
		{
			t: test{44, 4, 12, 1, 3},
			w: want{3, 0},
		},
		{
			t: test{44, 4, 8, 1, 2},
			w: want{0, 0},
		},
		{
			t: test{44, 4, 8, 1, 3},
			w: want{4, 3},
		},
		{
			t: test{44, 4, 9, 1, 4},
			w: want{-1, -1},
		},
		{
			t: test{44, 4, 12, 1, 4},
			w: want{4, 3},
		},
		{
			t: test{45, 4, 9, 1, 3},
			w: want{0, 3},
		},
		{
			t: test{7, 2, 3, 1, 1},
			w: want{0, 0},
		},
	}
	for i, tst := range tests {
		tst := tst
		name := fmt.Sprintf("test %d", i)
		t.Run(name, func(t *testing.T) {
			t.Logf("%d %d %d %d %d\n", tst.t.k1, tst.t.m, tst.t.k2, tst.t.p2, tst.t.n2)
			p1, n1 := solution(tst.t.k1, tst.t.m, tst.t.k2, tst.t.p2, tst.t.n2)
			assert.Equal(t, tst.w, want{p1, n1})
		})

	}

}
