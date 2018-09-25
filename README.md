# AstParseStr

Simple parsing of the ast tree from strings

## Example

```go
package main

import "github.com/saromanov/astparsestr"

func main() {
	src := `
4`
	astparsestr.AstParseExpr(src)
}

```
