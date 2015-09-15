// Package trie provides a simple tree structure for word dictionary
package trie

import (
	"encoding/json"
	"fmt"
	"io"
)

type Leaf struct {
	Leaves []Leaf      `json:"l,omitempty"`
	Value  rune        `json:"v"`
	Data   interface{} `json:"d,omitempty"`
}

type Trie struct {
	Tree []Leaf `json:"t"`
	Size int64  `json:"c"`
}

type Word struct {
	W string      `json:"w"`
	D interface{} `json:"d"`
}

// Trie functions

func (t *Trie) AddEntry(w string, d interface{}) {
	var found bool = false
	var i int

	s := []rune(w)
	for i = 0; i < len(t.Tree); i++ {
		if t.Tree[i].Value == s[0] {
			found = true
			break
		}
	}
	if !found {
		t.Tree = append(t.Tree, Leaf{nil, s[0], nil})
		i = len(t.Tree) - 1
	}
	t.Size++
	t.Tree[i].addLeaf(s[1:], d)
}

func (t *Trie) FindEntry(w string) interface{} {
	var found bool = false
	var i int

	s := []rune(w)
	for i = 0; i < len(t.Tree); i++ {
		if t.Tree[i].Value == s[0] {
			found = true
			break
		}
	}
	if !found {
		return nil
	}

	return t.Tree[i].findLeaf(s[1:])
}

func (t *Trie) FindEntries(w string) []Word {
	var found bool = false
	var i int
	var dict []Word
	var l *Leaf

	s := []rune(w)
	for i = 0; i < len(t.Tree); i++ {
		if t.Tree[i].Value == s[0] {
			found = true
			break
		}
	}
	if found {
		if l = t.Tree[i].findBrunch(s[1:]); l == nil {
			return nil
		}
		l.findLeaves(w, 0, &dict)
	}

	return dict
}

func (t *Trie) DictSize() int64 {
	return t.Size
}

func (t *Trie) PrintDict() {
	for i := 0; i < len(t.Tree); i++ {
		t.Tree[i].printWords("")
	}
}

func (t *Trie) Marshall(w io.Writer) error {
	return json.NewEncoder(w).Encode(t)
}

func (t *Trie) Unmarshal(r io.Reader) error {
	return json.NewDecoder(r).Decode(t)
}

// Leaf functions

func (l *Leaf) addLeaf(s []rune, d interface{}) {
	var found bool = false
	var i int

	if len(s) == 0 {
		l.Data = d
		return
	}
	for i = 0; i < len(l.Leaves); i++ {
		if l.Leaves[i].Value == s[0] {
			found = true
			break
		}
	}
	if !found {
		l.Leaves = append(l.Leaves, Leaf{nil, s[0], nil})
		i = len(l.Leaves) - 1
	}
	l.Leaves[i].addLeaf(s[1:], d)
}

func (l *Leaf) findLeaf(s []rune) interface{} {
	var found bool = false
	var i int

	if len(s) == 0 {
		return l.Data
	}
	for i = 0; i < len(l.Leaves); i++ {
		if l.Leaves[i].Value == s[0] {
			found = true
			break
		}
	}
	if !found {
		return nil
	}

	return l.Leaves[i].findLeaf(s[1:])
}

func (l *Leaf) findBrunch(s []rune) *Leaf {
	var found bool = false
	var i int

	if len(s) == 0 {
		return l
	}
	for i = 0; i < len(l.Leaves); i++ {
		if l.Leaves[i].Value == s[0] {
			found = true
			break
		}
	}
	if !found {
		return nil
	}

	return l.Leaves[i].findBrunch(s[1:])
}

func (l *Leaf) findLeaves(s string, level int, dict *[]Word) {
	var stack string

	if level == 0 {
		stack = s
	} else {
		stack = s + string(l.Value)
	}
	if l.Data != nil {
		w := Word{stack, l.Data}
		*dict = append(*dict, w)
	}
	if len(l.Leaves) > 0 {
		for i := 0; i < len(l.Leaves); i++ {
			l.Leaves[i].findLeaves(stack, level+1, dict)
		}
	}
}

func (l *Leaf) printWords(s string) {
	stack := s + string(l.Value)

	if l.Data != nil {
		fmt.Println(stack, l.Data)
	}
	if len(l.Leaves) > 0 {
		for i := 0; i < len(l.Leaves); i++ {
			l.Leaves[i].printWords(stack)
		}
	}
}
