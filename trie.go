// Package trie provides a simple tree structure for word dictionary
package trie

import (
	"encoding/json"
	"fmt"
	"io"
)

type Leaf struct {
	Leafes []Leaf      `json:"l,omitempty"`
	Value  rune        `json:"v"`
	Data   interface{} `json:"d,omitempty"`
}

type Trie struct {
	Tree []Leaf `json:"t"`
}

type Word struct {
	W string      `json:"w"`
	D interface{} `json:"d"`
}

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
		l.findLeafes(w, 0, &dict)
	}

	return dict
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

func (l *Leaf) addLeaf(s []rune, d interface{}) {
	var found bool = false
	var i int

	if len(s) == 0 {
		l.Data = d
		return
	}
	for i = 0; i < len(l.Leafes); i++ {
		if l.Leafes[i].Value == s[0] {
			found = true
			break
		}
	}
	if !found {
		l.Leafes = append(l.Leafes, Leaf{nil, s[0], nil})
		i = len(l.Leafes) - 1
	}
	l.Leafes[i].addLeaf(s[1:], d)
}

func (l *Leaf) findLeaf(s []rune) interface{} {
	var found bool = false
	var i int

	if len(s) == 0 {
		return l.Data
	}
	for i = 0; i < len(l.Leafes); i++ {
		if l.Leafes[i].Value == s[0] {
			found = true
			break
		}
	}
	if !found {
		return nil
	}

	return l.Leafes[i].findLeaf(s[1:])
}

func (l *Leaf) findBrunch(s []rune) *Leaf {
	var found bool = false
	var i int

	if len(s) == 0 {
		return l
	}
	for i = 0; i < len(l.Leafes); i++ {
		if l.Leafes[i].Value == s[0] {
			found = true
			break
		}
	}
	if !found {
		return nil
	}

	return l.Leafes[i].findBrunch(s[1:])
}

func (l *Leaf) findLeafes(s string, level int, dict *[]Word) {
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
	if len(l.Leafes) > 0 {
		for i := 0; i < len(l.Leafes); i++ {
			l.Leafes[i].findLeafes(stack, level+1, dict)
		}
	}
}

func (l *Leaf) printWords(s string) {
	stack := s + string(l.Value)

	if l.Data != nil {
		fmt.Println(stack, l.Data)
	}
	if len(l.Leafes) > 0 {
		for i := 0; i < len(l.Leafes); i++ {
			l.Leafes[i].printWords(stack)
		}
	}
}
