package bktree

import (
	"fmt"
	"strconv"
	"testing"
)

func TestEmptySearch(t *testing.T) {
	var tree BKTree
	results := tree.Search("", 0)
	if len(results) != 0 {
		t.Fatalf("empty tree should return empty results, bot got %d results", len(results))
	}
}

func TestExactMatch(t *testing.T) {
	var tree BKTree
	text := "hello"
	tree.Add(text)
	for i := 0; i < 100; i++ {
		tree.Add(text + strconv.Itoa(i))
	}

	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("searching %d", i), func(st *testing.T) {
			tmp := text + strconv.Itoa(i)
			results := tree.Search(tmp, 0)

			if len(results) != 1 {
				st.Fatalf("exact match should return only one result, but got %d results (%#v)", len(results), results)
			}
			if results[0].Distance != 0 {
				st.Fatalf("exact match result should have 0 as Distance field, but got %d", results[0].Distance)
			}
			if results[0].Text != tmp {
				st.Fatalf("expected result entry value is %v, but got %v", tmp, results[0].Text)
			}
		})
	}
}

func TestFuzzyMatch(t *testing.T) {
	var tree BKTree
	text := "hello"
	tree.Add(text)
	for i := 0; i < 100; i++ {
		tree.Add(text + strconv.Itoa(i))
	}

	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("searching %d", i), func(st *testing.T) {
			tmp := text + strconv.Itoa(i)
			results := tree.Search(tmp, 2)

			for _, result := range results {
				if result.Distance > 2 {
					st.Fatalf("Distance fields of results should be less than or equal to 2, but got %d", result.Distance)
				}
			}
		})
	}
}
