package bktree

import (
	"github.com/khmerlang/levenshtein"
	"sort"
	"strings"
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
		lowText := strings.ToLower(c.text)
		lowStr := strings.ToLower(str)
		d := levenshtein.Distance(lowText, lowStr)
		if d <= tolerance && strings.ContainsAny(lowText, lowStr) {
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

	sort.Slice(results, func(i, j int) bool {
		return results[i].Distance < results[j].Distance
	})

	return results
}

func (bk *BKTree) AddWithContext(str string, context string) {
	if bk.root == nil {
		bk.root = &node{
			text: str,
			context: context,
		}
		return
	}
	bk.root.addChildWithContext(str, context)
}

func (bk *BKTree) SearchGetContext(str string, tolerance int) []*Result {
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
				Text:     c.context,
			})
		}

		low, high := d-tolerance, d+tolerance
		for _, c := range c.children {
			if low <= c.distance && c.distance <= high {
				candidates = append(candidates, c.node)
			}
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Distance < results[j].Distance
	})

	return results
}
