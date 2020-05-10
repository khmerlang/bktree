package bktree

import (
	"github.com/khmerlang/levenshtein"
)

type node struct {
	text     string
	children []struct {
		distance int
		node     *node
	}
}

func (n *node) addChild(str string) {
	newnode := &node{text: str}

LOOP:
	d := levenshtein.Distance(n.text, str)
	for _, c := range n.children {
		if c.distance == d {
			n = c.node
			goto LOOP
		}
	}
	n.children = append(n.children, struct {
		distance int
		node     *node
	}{d, newnode})
}
