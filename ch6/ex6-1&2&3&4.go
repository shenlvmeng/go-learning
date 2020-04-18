package main

import (
	"bytes"
	"fmt"
	"math"
)

// IntSet is set of small non-negative integers
type IntSet struct {
	words []uint64
}

// Has report whether x in set s
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds non-negative integers
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds all given elements
func (s *IntSet) AddAll(e ...int) {
	for _, x := range e {
		s.Add(x)
	}
}

// Remove remove x from set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

func (s *IntSet) bitWith(t IntSet, compute func(uint64, uint64) uint64, keepT bool, keepS bool) {
	for i, maxLen := 0, math.Max(float64(len(s.words)), float64(len(t.words))); i < int(maxLen); i++ {
		if i >= len(s.words) {
			if keepT {
				s.words = append(s.words, t.words[i])
			}
		} else if i >= len(t.words) {
			if !keepS {
				s.words[i] = 0
			}
		} else {
			s.words[i] = compute(s.words[i], t.words[i])
		}
	}
}

// UnionWith return union of s and t
func (s *IntSet) UnionWith(t IntSet) {
	s.bitWith(t, func(si uint64, ti uint64) uint64 {
		return si | ti
	}, true, true)
}

// IntersectWith return intersection of s and t
func (s *IntSet) IntersectWith(t IntSet) {
	s.bitWith(t, func(si uint64, ti uint64) uint64 {
		return si & ti
	}, false, false)
}

// DifferenceWith return difference of s and t
func (s *IntSet) DifferenceWith(t IntSet) {
	s.bitWith(t, func(si uint64, ti uint64) uint64 {
		var temp uint64
		for i := 0; i < 64; i++ {
			if si&(1<<uint(i)) != 0 && ti&(1<<uint(i)) == 0 {
				temp |= 1 << uint(i)
			}
		}
		return temp
	}, false, true)
}

// SymmetricDifference return symmetric difference of s and t
func (s *IntSet) SymmetricDifference(t IntSet) {
	s.bitWith(t, func(si uint64, ti uint64) uint64 {
		return si ^ ti
	}, true, true)
}

// Copy return the copy of input set
func (s *IntSet) Copy() *IntSet {
	t := IntSet{
		words: make([]uint64, len(s.words)),
	}
	for i, word := range s.words {
		t.words[i] = word
	}
	return &t
}

// Clear remove all elements in set
func (s *IntSet) Clear() {
	s.words = []uint64{}
}

// Len returns the number of elements
func (s IntSet) Len() int {
	var count int
	for _, w := range s.words {
		if w == 0 {
			continue
		}
		for i := 0; i < 64; i++ {
			if w&(1<<uint(i)) != 0 {
				count++
			}
		}
	}
	return count
}

func (s IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) == 0 {
				continue
			}
			if buf.Len() > 1 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(&buf, "%d", 64*i+j)
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Elem returns elem store in set
func (s IntSet) Elem() []int {
	elem := []int{}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) == 0 {
				continue
			}
			elem = append(elem, 64*i+j)
		}
	}
	return elem
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(49)
	x.Add(201)
	fmt.Println(x, x.Len())

	y.AddAll(9, 49)
	x.UnionWith(y)
	// x.IntersectWith(y)
	// x.DifferenceWith(y)
	// x.SymmetricDifference(y)
	fmt.Println(x, x.Len())
	fmt.Println(x.Has(49), x.Has(9))

	x.Remove(100)
	x.Remove(9)
	fmt.Println(x)
	fmt.Println(x.Has(9), x.Has(100))

	x.Clear()
	fmt.Println(x, x.Len())
	x = *y.Copy()
	y.Add(66)
	fmt.Println(x, x.Len(), y, y.Len())
	fmt.Println(x.Elem())
}
