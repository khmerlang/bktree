### bktree

BK-tree is a data structure for search similar string.

#### Install

    go get github.com/khmerlang/bktree

#### Example

```go
package main

import (
  "fmt"
  "github.com/khmerlang/bktree"
)

func main() {
  var tree BKTree
  tree.Add("កកោស")
  tree.Add("កងកេង")
  tree.Add("កង្កែប")
  tree.Add("កង្ហែន")
  tree.Add("កញ្ចក់")

  results := tree.Search(កកោ, 2)
}

```
