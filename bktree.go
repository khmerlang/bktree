package bktree

import (
	"github.com/khmerlang/levenshtein"
)

type BKTree struct {
	root *node
}

func (bk *BKTree) Add(str string) {
	if bk.root == nil {
		bk.root = &node{
			text: str,
		}
		return
	}
	bk.root.addChild(str)
}

func (bk *BKTree) Search(str string, tolerance int) []*Result {
	results := make([]*Result, 0)
	if bk.root == nil {
		return results
	}
	candidates := []*node{bk.root}
	for len(candidates) != 0 {
		c := candidates[len(candidates)-1]
		candidates = candidates[:len(candidates)-1]
		d := levenshtein.Distance(c.text, str)
		if d <= tolerance {
			results = append(results, &Result{
				Distance: d,
				Text:     c.text,
			})
		}

		low, high := d-tolerance, d+tolerance
		for _, c := range c.children {
			if low <= c.distance && c.distance <= high {
				candidates = append(candidates, c.node)
			}
		}
	}
	return results
}
