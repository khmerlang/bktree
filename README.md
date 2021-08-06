### bktree

BK-tree is a data structure for search similar string.

#### Install

    go get github.com/khmerlang/bktree

#### Example

1. Normal tree
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

1. Normal tree with context
```go
package main

import (
  "fmt"
  "github.com/khmerlang/bktree"
)

func main() {
  var tree BKTree
  tree.AddWithContext("koka", "កកោស")
  tree.AddWithContext("koki", "កងកេង")
  tree.AddWithContext("sona", "កង្កែប")
  tree.AddWithContext("soni", "កង្ហែន")
  tree.AddWithContext("sony", "កញ្ចក់")

  results := tree.SearchGetContext("kok", 2)
}

```
