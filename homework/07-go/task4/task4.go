package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Doubly linked list for Loubère's algorithm
// It helps to find a new element without tracking the borders of the square

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type IndList struct {
	Head *Node
}

// Doubly linked list creation

func CreateIndList() *IndList {
	head := new(Node)
	return &IndList{head}
}

// Adding slice Items to a List

func (list *IndList) Init(x []int) bool {
	if list.Head == nil {
		return false
	}
	p := list.Head
	pp := list.Head
	for i := 0; i < len(x)-1; i++ {
		p.Data = x[i]
		p.Next = new(Node)
		pp = p
		p = p.Next
		p.Prev = pp
	}
	p.Data = x[len(x)-1]
	p.Next = list.Head
	list.Head.Prev = p
	return true
}

// MAGIC SQUARE generation

func Magicsgen(n int) [][]*int {
	var (
		indexes []int
		magisq  [][]*int
	)
	// only for odd-order squares
	if n%2 != 1 {
		return magisq
	}

	row := CreateIndList()
	col := CreateIndList()
	// creation double slice before square generation
	for i := 0; i < n; i++ {
		indexes = append(indexes, i) // slice for indexes creation
		magisq = append(magisq, nil)
		for j := 0; j < n; j++ {
			magisq[i] = append(magisq[i], nil)
		}
	}
	row.Init(indexes)
	col.Init(indexes)

	// shift to the starting position of the algorithm
	for i := 0; i < n/2; i++ {
		col.Head = col.Head.Next
	}

	// algorithm implementation
	for i := 1; i <= n*n; i++ {
		magisq[row.Head.Data][col.Head.Data] = new(int)
		*magisq[row.Head.Data][col.Head.Data] = i
		if magisq[row.Head.Prev.Data][col.Head.Next.Data] == nil {
			row.Head = row.Head.Prev
			col.Head = col.Head.Next
		} else {
			row.Head = row.Head.Next
		}
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rand.Intn(4); i++ {
		Rotates(magisq)
	}
	for i := 0; i < rand.Intn(2); i++ {
		Reflects(magisq)
	}

	return magisq
}

// clockwise rotation 90'

func Rotates(sl [][]*int) {
	var (
		n      = len(sl)
		b      int
		pr, nt []int
	)

	for k := 0; k < n/2; k++ {
		b = n - k - 1

		j := 0
		pr = nil
		nt = nil
		for i := k; i < b; i++ {
			nt = append(nt, *sl[k][i])
		}
		for i := b; i > k; i-- {
			pr = append(pr, *sl[i][k])
		}
		for i := k; i < b; i++ {
			sl[k][i] = &pr[j]
			j++
		}

		j = 0
		pr = nt
		nt = nil
		for i := k; i < b; i++ {
			nt = append(nt, *sl[i][b])
			sl[i][b] = &pr[j]
			j++
		}

		j = 0
		pr = nt
		nt = nil
		for i := b; i > k; i-- {
			nt = append(nt, *sl[b][i])
		}
		for i := b; i > k; i-- {
			sl[b][i] = &pr[j]
			j++
		}

		j = 0
		for i := b; i > k; i-- {
			sl[i][k] = &nt[j]
			j++
		}
	}
}

// mirror slice

func Reflects(sl [][]*int) {
	var (
		p []*int
		f = 0
		l = len(sl) - 1
	)

	for f != l {
		p = sl[f]
		sl[f] = sl[l]
		sl[l] = p
		f++
		l--
	}
}

// Print double slice

func Prints(sl [][]*int) {
	for i := 0; i < len(sl); i++ {
		for j := 0; j < len(sl[i]); j++ {
			fmt.Printf("%3d ", *sl[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func main() {
	Prints(Magicsgen(3))
}
