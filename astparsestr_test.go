package astparsestr

import (
	"testing"
)

func TestAstParseExpr(t *testing.T) {
	src := `
4
`
	node := AstParseExpr(src)
	if !node.Pos().IsValid() {
		t.Errorf("unable to parse expression")
	}

}

func TestAstParseFile(t *testing.T) {
	src := `
const item = "AAA"`
	_, err := AstParseFile(src)
	if err != nil {
		t.Errorf("unable to parse file: %v", err)
	}
}
